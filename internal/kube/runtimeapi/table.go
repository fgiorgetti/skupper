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

func adaptorTableColumns() []metav1.TableColumnDefinition {
	return []metav1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: "Name of the runtime adaptor"},
		{Name: "Type", Type: "string", Description: "Adaptor type (tcpListener or tcpConnector)"},
		{Name: "Oper-Status", Type: "string", Description: "Live router operational status"},
		{Name: "Address", Type: "string", Description: "Skupper routing address"},
		{Name: "Host", Type: "string", Description: "Adaptor host"},
		{Name: "Port", Type: "string", Description: "Adaptor port"},
		{Name: "Age", Type: "string", Description: "Time since adaptor was recorded"},
		{Name: "Conn-Opened", Type: "integer", Priority: 1, Description: "Total connections opened"},
		{Name: "Conn-Closed", Type: "integer", Priority: 1, Description: "Total connections closed"},
		{Name: "Bytes-In", Type: "integer", Priority: 1, Description: "Total ingress bytes"},
		{Name: "Bytes-Out", Type: "integer", Priority: 1, Description: "Total egress bytes"},
		{Name: "Site", Type: "string", Priority: 1, Description: "Skupper site name"},
		{Name: "Site-ID", Type: "string", Priority: 1, Description: "Skupper site ID"},
		{Name: "SSL-Profile", Type: "string", Priority: 1, Description: "Router SSL profile name"},
		{Name: "Observer", Type: "string", Priority: 1, Description: "Router protocol observer"},
	}
}

func adaptorsToTable(adaptors []Adaptor, tableVersion string, includeObject string, now time.Time) (*metav1.Table, error) {
	if tableVersion == "" {
		tableVersion = "v1"
	}
	table := &metav1.Table{
		TypeMeta: metav1.TypeMeta{
			APIVersion: metav1.GroupName + "/" + tableVersion,
			Kind:       "Table",
		},
		ColumnDefinitions: adaptorTableColumns(),
		Rows:              make([]metav1.TableRow, 0, len(adaptors)),
	}
	for _, adp := range adaptors {
		if adp.APIVersion == "" {
			adp.APIVersion = APIVersion
		}
		if adp.Kind == "" {
			adp.Kind = AdaptorKind
		}
		row := metav1.TableRow{
			Cells: []interface{}{
				adp.Name,
				string(adp.Status.Type),
				adp.Status.OperStatus,
				adp.Status.Address,
				adp.Status.Host,
				adp.Status.Port,
				formatAge(adp.CreationTimestamp, now),
				adp.Status.ConnectionsOpened,
				adp.Status.ConnectionsClosed,
				adp.Status.BytesIn,
				adp.Status.BytesOut,
				adp.Status.SiteName,
				adp.Status.SiteID,
				adp.Status.SslProfile,
				adp.Status.Observer,
			},
		}
		raw, err := marshalRowObject(adp, includeObject)
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

func marshalRowObject(adp Adaptor, includeObject string) ([]byte, error) {
	switch includeObject {
	case string(metav1.IncludeNone):
		return nil, nil
	case string(metav1.IncludeMetadata):
		metaObj := metav1.PartialObjectMetadata{
			TypeMeta: metav1.TypeMeta{
				APIVersion: metav1.SchemeGroupVersion.String(),
				Kind:       "PartialObjectMetadata",
			},
			ObjectMeta: adp.ObjectMeta,
		}
		return json.Marshal(metaObj)
	default:
		return json.Marshal(adp)
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
