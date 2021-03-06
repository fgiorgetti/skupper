package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/skupperproject/skupper/api/types"
	"github.com/skupperproject/skupper/client"
	"github.com/skupperproject/skupper/pkg/data"
	"github.com/skupperproject/skupper/pkg/event"
	"github.com/skupperproject/skupper/pkg/kube"
	"github.com/skupperproject/skupper/pkg/qdr"
)

const (
	SiteQueryError      string = "SiteQueryError"
	SiteQueryRequest    string = "SiteQueryRequest"
	ServiceCheckError   string = "ServiceCheckError"
	ServiceCheckRequest string = "ServiceCheckRequest"
)

type SiteQueryServer struct {
	client    *client.VanClient
	tlsConfig *tls.Config
	agentPool *qdr.AgentPool
	server    *qdr.RequestServer
	iplookup  *IpLookup
	siteInfo  data.Site
}

func newSiteQueryServer(cli *client.VanClient, config *tls.Config) *SiteQueryServer {
	sqs := SiteQueryServer{
		client:    cli,
		tlsConfig: config,
		agentPool: qdr.NewAgentPool("amqps://"+types.LocalTransportServiceName+":5671", config),
		iplookup:  NewIpLookup(cli),
	}
	sqs.getLocalSiteInfo()
	sqs.server = qdr.NewRequestServer(getSiteQueryAddress(sqs.siteInfo.SiteId), &sqs, sqs.agentPool)
	return &sqs
}

func siteQueryError(err error) {
}

func (s *SiteQueryServer) getLocalSiteInfo() {
	s.siteInfo.SiteId = os.Getenv("SKUPPER_SITE_ID")
	s.siteInfo.SiteName = os.Getenv("SKUPPER_SITE_NAME")
	s.siteInfo.Namespace = os.Getenv("SKUPPER_NAMESPACE")
	s.siteInfo.Version = client.Version
	url, err := getSiteUrl(s.client)
	if err != nil {
		event.Recordf(SiteQueryError, "Failed to get site url: %s", err)
	} else {
		s.siteInfo.Url = url
	}
}

func (s *SiteQueryServer) getLocalSiteQueryData() (data.SiteQueryData, error) {
	data := data.SiteQueryData{
		Site: s.siteInfo,
	}
	agent, err := s.agentPool.Get()
	if err != nil {
		return data, fmt.Errorf("Could not get management agent: %s", err)
	}
	defer s.agentPool.Put(agent)

	routers, err := agent.GetAllRouters()
	if err != nil {
		return data, fmt.Errorf("Error retrieving routers: %s", err)
	}
	err = getServiceInfo(agent, routers, &data, s.iplookup)
	if err != nil {
		return data, fmt.Errorf("Error getting local service info: %s", err)
	}
	return data, nil
}

func getSiteUrl(vanClient *client.VanClient) (string, error) {
	if vanClient.RouteClient == nil {
		service, err := vanClient.KubeClient.CoreV1().Services(vanClient.Namespace).Get(types.TransportServiceName, metav1.GetOptions{})
		if err != nil {
			return "", err
		} else {
			if service.Spec.Type == corev1.ServiceTypeLoadBalancer {
				host := kube.GetLoadBalancerHostOrIp(service)
				return host, nil
			} else {
				return "", nil
			}
		}
	} else {
		route, err := vanClient.RouteClient.Routes(vanClient.Namespace).Get("skupper-inter-router", metav1.GetOptions{})
		if err != nil {
			return "", err
		} else {
			return route.Spec.Host, nil
		}
	}
}

func getSiteQueryAddress(siteId string) string {
	return siteId + "/skupper-site-query"
}

const (
	ServiceCheck string = "service-check"
)

func (s *SiteQueryServer) Request(request *qdr.Request) (*qdr.Response, error) {
	if request.Type == ServiceCheck {
		return s.HandleServiceCheck(request)
	} else {
		return s.HandleSiteQuery(request)
	}
}

func (s *SiteQueryServer) HandleSiteQuery(request *qdr.Request) (*qdr.Response, error) {
	//if request has explicit version, send SiteQueryData, else send LegacySiteData
	if request.Version == "" {
		event.Record(SiteQueryRequest, "legacy site data request")
		data := s.siteInfo.AsLegacySiteInfo()
		bytes, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("Could not encode response: %s", err)
		}
		return &qdr.Response{
			Version: client.Version,
			Body:    string(bytes),
		}, nil
	} else {
		event.Record(SiteQueryRequest, "site data request")
		data, err := s.getLocalSiteQueryData()
		if err != nil {
			return nil, fmt.Errorf("Could not get response: %s", err)
		}
		bytes, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("Could not encode response: %s", err)
		}
		return &qdr.Response{
			Version: client.Version,
			Body:    string(bytes),
		}, nil
	}
}

func (s *SiteQueryServer) HandleServiceCheck(request *qdr.Request) (*qdr.Response, error) {
	event.Recordf(ServiceCheckRequest, "checking service %s", request.Body)
	data, err := s.getServiceDetail(context.Background(), request.Body)
	if err != nil {
		return &qdr.Response{
			Version: client.Version,
			Type:    ServiceCheckError,
			Body:    err.Error(),
		}, nil
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Could not encode service check response: %s", err)
	}
	return &qdr.Response{
		Version: client.Version,
		Type:    request.Type,
		Body:    string(bytes),
	}, nil
}

func (s *SiteQueryServer) run() {
	for {
		ctxt := context.Background()
		err := s.server.Run(ctxt)
		if err != nil {
			event.Recordf(SiteQueryError, "Error handling requests: %s", err)
		}
	}
}

func (s *SiteQueryServer) start(stopCh <-chan struct{}) error {
	err := s.iplookup.start(stopCh)
	go s.run()
	return err
}

func (s *SiteQueryServer) getServiceDetail(context context.Context, address string) (data.ServiceDetail, error) {
	detail := data.ServiceDetail{
		SiteId: s.siteInfo.SiteId,
	}
	definition, err := s.client.ServiceInterfaceInspect(context, address)
	if err != nil {
		return detail, err
	}
	if definition == nil {
		return detail, fmt.Errorf("No such service %q", address)
	}
	detail.Definition = *definition

	service, err := kube.GetService(address, s.client.Namespace, s.client.KubeClient)
	if err != nil {
		return detail, err
	}
	if len(service.Spec.Ports) == 1 {
		detail.IngressBinding.ServicePort = int(service.Spec.Ports[0].Port)
		detail.IngressBinding.ServiceTargetPort = service.Spec.Ports[0].TargetPort.IntValue()
	} else if len(service.Spec.Ports) > 1 {
		var name string
		for _, ports := range service.Spec.Ports {
			if int(ports.Port) == detail.Definition.Port {
				name = ports.Name
				detail.IngressBinding.ServicePort = int(ports.Port)
				detail.IngressBinding.ServiceTargetPort = ports.TargetPort.IntValue()
				break
			}
		}
		if name == "" {
			detail.AddObservation("Service Spec has multiple ports defined, none of which match port in definition")
		} else {
			detail.AddObservation("Service Spec has multiple ports defined; using " + name)
		}
	} else {
		detail.AddObservation("Service Spec has no ports defined")
	}
	detail.IngressBinding.ServiceSelector = service.Spec.Selector

	agent, err := s.agentPool.Get()
	if err != nil {
		return detail, fmt.Errorf("Could not get management agent: %s", err)
	}
	defer s.agentPool.Put(agent)

	if detail.Definition.Protocol == "tcp" {
		listener, err := agent.GetLocalTcpListener(detail.Definition.Address, detail.IngressBinding.ServiceTargetPort)
		if err != nil {
			return detail, fmt.Errorf("Error retrieving tcp listener for %s: %s", detail.Definition.Address, err)
		}
		if listener == nil {
			detail.AddObservation(fmt.Sprintf("No tcp listener defined for %s on %d", detail.Definition.Address, detail.IngressBinding.ServiceTargetPort))
		} else {
			if detail.Definition.Address != listener.Address {
				detail.AddObservation(fmt.Sprintf("Wrong address for tcp listener on %d", detail.IngressBinding.ServiceTargetPort))
			} else {
				port, err := strconv.Atoi(listener.Port)
				if err != nil {
					detail.AddObservation(fmt.Sprintf("Bad port for listener %s: %s %s", listener.Name, listener.Port, err))
				}
				detail.IngressBinding.ListenerPort = port
				if detail.IngressBinding.ListenerPort != detail.IngressBinding.ServiceTargetPort {
					detail.AddObservation(fmt.Sprintf("listener port does not match service target port (%d != %d)",
						detail.IngressBinding.ListenerPort, detail.IngressBinding.ServiceTargetPort))
				}
			}
		}

		connectors, err := agent.GetLocalTcpConnectors(detail.Definition.Address)
		if err != nil {
			return detail, fmt.Errorf("Error retrieving tcp connectors for %s: %s", detail.Definition.Address, err)
		}
		for _, connector := range connectors {
			port, err := strconv.Atoi(connector.Port)
			if err != nil {
				detail.AddObservation(fmt.Sprintf("Bad port for connector %s: %s %s", connector.Name, connector.Port, err))
			}
			detail.EgressBindings = append(detail.EgressBindings, data.EgressBinding{
				Port: port,
				Host: connector.Host,
			})
		}
	} else if detail.Definition.Protocol == "http" || detail.Definition.Protocol == "http2" {
		listener, err := agent.GetLocalHttpListener(detail.Definition.Address, detail.IngressBinding.ServiceTargetPort)
		if err != nil {
			return detail, fmt.Errorf("Error retrieving http listener for %s: %s", detail.Definition.Address, err)
		}
		if listener == nil {
			detail.AddObservation(fmt.Sprintf("No http listener defined for %s on %d", detail.Definition.Address, detail.IngressBinding.ServiceTargetPort))
		} else {
			if detail.Definition.Address != listener.Address {
				detail.AddObservation(fmt.Sprintf("Wrong address for http listener on %d", detail.IngressBinding.ServiceTargetPort))
			} else {
				port, err := strconv.Atoi(listener.Port)
				if err != nil {
					detail.AddObservation(fmt.Sprintf("Bad port for listener %s: %s %s", listener.Name, listener.Port, err))
				}
				detail.IngressBinding.ListenerPort = port
				if detail.IngressBinding.ListenerPort != detail.IngressBinding.ServiceTargetPort {
					detail.AddObservation(fmt.Sprintf("listener port does not match service target port (%d != %d)",
						detail.IngressBinding.ListenerPort, detail.IngressBinding.ServiceTargetPort))
				}
			}
		}

		connectors, err := agent.GetLocalHttpConnectors(detail.Definition.Address)
		if err != nil {
			return detail, fmt.Errorf("Error retrieving http connectors for %s: %s", detail.Definition.Address, err)
		}
		for _, connector := range connectors {
			port, err := strconv.Atoi(connector.Port)
			if err != nil {
				detail.AddObservation(fmt.Sprintf("Bad port for connector %s: %s %s", connector.Name, connector.Port, err))
			}
			detail.EgressBindings = append(detail.EgressBindings, data.EgressBinding{
				Port: port,
				Host: connector.Host,
			})
		}
	} else {
		return detail, fmt.Errorf("Unrecognised protocol: %s", detail.Definition.Protocol)
	}

	if len(detail.Definition.Targets) > 0 && len(detail.EgressBindings) == 0 {
		detail.AddObservation(fmt.Sprintf("No connectors on %s for %s ", detail.SiteId, detail.Definition.Address))
	}
	return detail, nil
}

func querySites(agent qdr.RequestResponse, sites []data.SiteQueryData) {
	for i, s := range sites {
		request := qdr.Request{
			Address: getSiteQueryAddress(s.SiteId),
			Version: client.Version,
		}
		response, err := agent.Request(&request)
		if err != nil {
			event.Recordf(SiteQueryError, "Request to %s failed: %s", s.SiteId, err)
		} else if response.Version == "" {
			//assume legacy version of site-query protocol
			info := data.LegacySiteInfo{}
			err := json.Unmarshal([]byte(response.Body), &info)
			if err != nil {
				event.Recordf(SiteQueryError, "Error parsing legacy json %q from %s: %s", response.Body, s.SiteId, err)
			} else {
				sites[i].SiteName = info.SiteName
				sites[i].Namespace = info.Namespace
				sites[i].Url = info.Url
				sites[i].Version = info.Version
			}
		} else {
			site := data.SiteQueryData{}
			err := json.Unmarshal([]byte(response.Body), &site)
			if err != nil {
				event.Recordf(SiteQueryError, "Error parsing json for site query %q from %s: %s", response.Body, s.SiteId, err)
			} else {
				sites[i].SiteName = site.SiteName
				sites[i].Namespace = site.Namespace
				sites[i].Url = site.Url
				sites[i].Version = site.Version
				sites[i].TcpServices = site.TcpServices
				sites[i].HttpServices = site.HttpServices
			}
		}
	}
}

func getServiceInfo(agent *qdr.Agent, network []qdr.Router, site *data.SiteQueryData, lookup data.NameMapping) error {
	routers := qdr.GetRoutersForSite(network, site.SiteId)
	bridges, err := agent.GetBridges(routers)
	if err != nil {
		return fmt.Errorf("Error retrieving bridge configuration: %s", err)
	}
	httpRequestInfo, err := agent.GetHttpRequestInfo(routers)
	if err != nil {
		return fmt.Errorf("Error retrieving http request info: %s", err)
	}
	tcpConnections, err := agent.GetTcpConnections(routers)
	if err != nil {
		return fmt.Errorf("Error retrieving tcp connection info: %s", err)
	}

	site.HttpServices = data.GetHttpServices(site.SiteId, httpRequestInfo, qdr.GetHttpConnectors(bridges), qdr.GetHttpListeners(bridges), lookup)
	site.TcpServices = data.GetTcpServices(site.SiteId, tcpConnections, qdr.GetTcpConnectors(bridges), lookup)
	return nil
}

func checkServiceForSites(agent qdr.RequestResponse, address string, sites *data.ServiceCheck) error {
	details := []data.ServiceDetail{}
	for _, s := range sites.Details {
		request := qdr.Request{
			Address: getSiteQueryAddress(s.SiteId),
			Version: client.Version,
			Type:    ServiceCheck,
			Body:    address,
		}
		response, err := agent.Request(&request)
		if err != nil {
			event.Recordf(ServiceCheckError, "Request to %s failed: %s", s.SiteId, err)
			return err
		}
		if response.Type == ServiceCheckError {
			sites.AddObservation(fmt.Sprintf("%s on %s", response.Body, s.SiteId))
		} else {
			detail := data.ServiceDetail{}
			err = json.Unmarshal([]byte(response.Body), &detail)
			if err != nil {
				event.Recordf(ServiceCheckError, "Error parsing json for service check %q from %s: %s", response.Body, s.SiteId, err)
				return err
			}
			details = append(details, detail)
		}
	}
	sites.Details = details
	data.CheckService(sites)
	return nil
}
