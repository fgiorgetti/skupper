// Code generated by go-swagger; DO NOT EDIT.

package volumes_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new volumes compat API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for volumes compat API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	VolumeCreate(params *VolumeCreateParams, opts ...ClientOption) (*VolumeCreateCreated, error)

	VolumeDelete(params *VolumeDeleteParams, opts ...ClientOption) (*VolumeDeleteNoContent, error)

	VolumeInspect(params *VolumeInspectParams, opts ...ClientOption) (*VolumeInspectOK, error)

	VolumeList(params *VolumeListParams, opts ...ClientOption) (*VolumeListOK, error)

	VolumePrune(params *VolumePruneParams, opts ...ClientOption) (*VolumePruneOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  VolumeCreate creates a volume
*/
func (a *Client) VolumeCreate(params *VolumeCreateParams, opts ...ClientOption) (*VolumeCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumeCreate",
		Method:             "POST",
		PathPattern:        "/volumes/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumeCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VolumeDelete removes volume
*/
func (a *Client) VolumeDelete(params *VolumeDeleteParams, opts ...ClientOption) (*VolumeDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumeDelete",
		Method:             "DELETE",
		PathPattern:        "/volumes/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumeDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VolumeInspect inspects volume
*/
func (a *Client) VolumeInspect(params *VolumeInspectParams, opts ...ClientOption) (*VolumeInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumeInspect",
		Method:             "GET",
		PathPattern:        "/volumes/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeInspectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumeInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VolumeList lists volumes

  Returns a list of volume
*/
func (a *Client) VolumeList(params *VolumeListParams, opts ...ClientOption) (*VolumeListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumeListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumeList",
		Method:             "GET",
		PathPattern:        "/volumes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumeListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumeListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumeList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  VolumePrune prunes volumes
*/
func (a *Client) VolumePrune(params *VolumePruneParams, opts ...ClientOption) (*VolumePruneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVolumePruneParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "VolumePrune",
		Method:             "POST",
		PathPattern:        "/volumes/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VolumePruneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VolumePruneOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for VolumePrune: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
