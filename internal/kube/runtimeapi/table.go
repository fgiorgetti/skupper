package runtimeapi

import (
	"encoding/json"
	"mime"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/duration"
)

func wantsTable(acceptHeader string) (bool, string) {
	for _, part := range strings.Split(acceptHeader, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		mediaType, params, err := mime.ParseMediaType(part)
		if err != nil {
			continue
		}
		if mediaType != "application/json" && mediaType != "*/*" {
			continue
		}
		if params["as"] != "Table" {
			continue
		}
		if g := params["g"]; g != "" && g != metav1.GroupName {
			continue
		}
		v := params["v"]
		if v == "" {
			v = "v1"
		}
		if v == "v1" || v == "v1beta1" {
			return true, v
		}
	}
	return false, ""
}

func endpointTableColumns() []metav1.TableColumnDefinition {
	return []metav1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: "Name of the runtime endpoint"},
		{Name: "Type", Type: "string", Description: "Endpoint type (tcpListener or tcpConnector)"},
		{Name: "Address", Type: "string", Description: "Skupper routing address"},
		{Name: "Host", Type: "string", Description: "Endpoint host"},
		{Name: "Port", Type: "string", Description: "Endpoint port"},
		{Name: "Age", Type: "string", Description: "Time since endpoint was recorded"},
		{Name: "Site", Type: "string", Priority: 1, Description: "Skupper site name"},
		{Name: "Site-ID", Type: "string", Priority: 1, Description: "Skupper site ID"},
		{Name: "SSL-Profile", Type: "string", Priority: 1, Description: "Router SSL profile name"},
		{Name: "Observer", Type: "string", Priority: 1, Description: "Router protocol observer"},
	}
}

func endpointsToTable(endpoints []Endpoint, tableVersion string, includeObject string, now time.Time) (*metav1.Table, error) {
	if tableVersion == "" {
		tableVersion = "v1"
	}
	table := &metav1.Table{
		TypeMeta: metav1.TypeMeta{
			APIVersion: metav1.GroupName + "/" + tableVersion,
			Kind:       "Table",
		},
		ColumnDefinitions: endpointTableColumns(),
		Rows:              make([]metav1.TableRow, 0, len(endpoints)),
	}
	for _, ep := range endpoints {
		if ep.APIVersion == "" {
			ep.APIVersion = APIVersion
		}
		if ep.Kind == "" {
			ep.Kind = EndpointKind
		}
		row := metav1.TableRow{
			Cells: []interface{}{
				ep.Name,
				string(ep.Status.Type),
				ep.Status.Address,
				ep.Status.Host,
				ep.Status.Port,
				formatAge(ep.CreationTimestamp, now),
				ep.Status.SiteName,
				ep.Status.SiteID,
				ep.Status.SslProfile,
				ep.Status.Observer,
			},
		}
		raw, err := marshalRowObject(ep, includeObject)
		if err != nil {
			return nil, err
		}
		if raw != nil {
			row.Object = runtime.RawExtension{Raw: raw}
		}
		table.Rows = append(table.Rows, row)
	}
	return table, nil
}

func marshalRowObject(ep Endpoint, includeObject string) ([]byte, error) {
	switch includeObject {
	case string(metav1.IncludeNone):
		return nil, nil
	case string(metav1.IncludeMetadata):
		metaObj := metav1.PartialObjectMetadata{
			TypeMeta: metav1.TypeMeta{
				APIVersion: metav1.SchemeGroupVersion.String(),
				Kind:       "PartialObjectMetadata",
			},
			ObjectMeta: ep.ObjectMeta,
		}
		return json.Marshal(metaObj)
	default:
		return json.Marshal(ep)
	}
}

func formatAge(ts metav1.Time, now time.Time) string {
	if ts.IsZero() {
		return "<unknown>"
	}
	d := now.Sub(ts.Time)
	if d < 0 {
		d = 0
	}
	return duration.HumanDuration(d)
}
