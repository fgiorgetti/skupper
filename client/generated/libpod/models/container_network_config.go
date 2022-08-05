// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerNetworkConfig ContainerNetworkConfig contains information on a container's network
// configuration.
//
// swagger:model ContainerNetworkConfig
type ContainerNetworkConfig struct {

	// Aliases are a list of network-scoped aliases for container
	// Optional
	Aliases map[string][]string `json:"aliases,omitempty"`

	// CNINetworks is a list of CNI networks to join the container to.
	// If this list is empty, the default CNI network will be joined
	// instead. If at least one entry is present, we will not join the
	// default network (unless it is part of this list).
	// Only available if NetNS is set to bridge.
	// Optional.
	CNINetworks []string `json:"cni_networks"`

	// DNSOptions is a set of DNS options that will be used in the
	// container's resolv.conf, replacing the host's DNS options which are
	// used by default.
	// Conflicts with UseImageResolvConf.
	// Optional.
	DNSOptions []string `json:"dns_option"`

	// DNSSearch is a set of DNS search domains that will be used in the
	// container's resolv.conf, replacing the host's DNS search domains
	// which are used by default.
	// Conflicts with UseImageResolvConf.
	// Optional.
	DNSSearch []string `json:"dns_search"`

	// DNSServers is a set of DNS servers that will be used in the
	// container's resolv.conf, replacing the host's DNS Servers which are
	// used by default.
	// Conflicts with UseImageResolvConf.
	// Optional.
	DNSServers []IP `json:"dns_server"`

	// Expose is a number of ports that will be forwarded to the container
	// if PublishExposedPorts is set.
	// Expose is a map of uint16 (port number) to a string representing
	// protocol. Allowed protocols are "tcp", "udp", and "sctp", or some
	// combination of the three separated by commas.
	// If protocol is set to "" we will assume TCP.
	// Only available if NetNS is set to Bridge or Slirp, and
	// PublishExposedPorts is set.
	// Optional.
	Expose interface{} `json:"expose,omitempty"`

	// HostAdd is a set of hosts which will be added to the container's
	// etc/hosts file.
	// Conflicts with UseImageHosts.
	// Optional.
	HostAdd []string `json:"hostadd"`

	// NetworkOptions are additional options for each network
	// Optional.
	NetworkOptions map[string][]string `json:"network_options,omitempty"`

	// PortBindings is a set of ports to map into the container.
	// Only available if NetNS is set to bridge or slirp.
	// Optional.
	PortMappings []*PortMapping `json:"portmappings"`

	// PublishExposedPorts will publish ports specified in the image to
	// random unused ports (guaranteed to be above 1024) on the host.
	// This is based on ports set in Expose below, and any ports specified
	// by the Image (if one is given).
	// Only available if NetNS is set to Bridge or Slirp.
	PublishExposedPorts bool `json:"publish_image_ports,omitempty"`

	// UseImageHosts indicates that /etc/hosts should not be managed by
	// Podman, and instead sourced from the image.
	// Conflicts with HostAdd.
	UseImageHosts bool `json:"use_image_hosts,omitempty"`

	// UseImageResolvConf indicates that resolv.conf should not be managed
	// by Podman, but instead sourced from the image.
	// Conflicts with DNSServer, DNSSearch, DNSOption.
	UseImageResolvConf bool `json:"use_image_resolve_conf,omitempty"`

	// netns
	Netns *Namespace `json:"netns,omitempty"`

	// static ip
	StaticIP IP `json:"static_ip,omitempty"`

	// static ipv6
	StaticIPV6 IP `json:"static_ipv6,omitempty"`

	// static mac
	StaticMac HardwareAddr `json:"static_mac,omitempty"`
}

// Validate validates this container network config
func (m *ContainerNetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticIPV6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticMac(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerNetworkConfig) validateDNSServers(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSServers) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSServers); i++ {

		if err := m.DNSServers[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_server" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns_server" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ContainerNetworkConfig) validatePortMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.PortMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.PortMappings); i++ {
		if swag.IsZero(m.PortMappings[i]) { // not required
			continue
		}

		if m.PortMappings[i] != nil {
			if err := m.PortMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("portmappings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("portmappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContainerNetworkConfig) validateNetns(formats strfmt.Registry) error {
	if swag.IsZero(m.Netns) { // not required
		return nil
	}

	if m.Netns != nil {
		if err := m.Netns.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netns")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerNetworkConfig) validateStaticIP(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticIP) { // not required
		return nil
	}

	if err := m.StaticIP.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_ip")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("static_ip")
		}
		return err
	}

	return nil
}

func (m *ContainerNetworkConfig) validateStaticIPV6(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticIPV6) { // not required
		return nil
	}

	if err := m.StaticIPV6.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_ipv6")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("static_ipv6")
		}
		return err
	}

	return nil
}

func (m *ContainerNetworkConfig) validateStaticMac(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticMac) { // not required
		return nil
	}

	if err := m.StaticMac.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_mac")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("static_mac")
		}
		return err
	}

	return nil
}

// ContextValidate validate this container network config based on the context it is used
func (m *ContainerNetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticMac(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerNetworkConfig) contextValidateDNSServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSServers); i++ {

		if err := m.DNSServers[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dns_server" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dns_server" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ContainerNetworkConfig) contextValidatePortMappings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PortMappings); i++ {

		if m.PortMappings[i] != nil {
			if err := m.PortMappings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("portmappings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("portmappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ContainerNetworkConfig) contextValidateNetns(ctx context.Context, formats strfmt.Registry) error {

	if m.Netns != nil {
		if err := m.Netns.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netns")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netns")
			}
			return err
		}
	}

	return nil
}

func (m *ContainerNetworkConfig) contextValidateStaticIP(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticIP.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_ip")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("static_ip")
		}
		return err
	}

	return nil
}

func (m *ContainerNetworkConfig) contextValidateStaticIPV6(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticIPV6.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_ipv6")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("static_ipv6")
		}
		return err
	}

	return nil
}

func (m *ContainerNetworkConfig) contextValidateStaticMac(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticMac.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_mac")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("static_mac")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerNetworkConfig) UnmarshalBinary(b []byte) error {
	var res ContainerNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
