// Code generated by go-swagger; DO NOT EDIT.

package pods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pods API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pods API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PodCreateLibpod(params *PodCreateLibpodParams, opts ...ClientOption) (*PodCreateLibpodCreated, error)

	PodDeleteLibpod(params *PodDeleteLibpodParams, opts ...ClientOption) (*PodDeleteLibpodOK, error)

	PodExistsLibpod(params *PodExistsLibpodParams, opts ...ClientOption) (*PodExistsLibpodNoContent, error)

	PodInspectLibpod(params *PodInspectLibpodParams, opts ...ClientOption) (*PodInspectLibpodOK, error)

	PodKillLibpod(params *PodKillLibpodParams, opts ...ClientOption) (*PodKillLibpodOK, error)

	PodListLibpod(params *PodListLibpodParams, opts ...ClientOption) (*PodListLibpodOK, error)

	PodPauseLibpod(params *PodPauseLibpodParams, opts ...ClientOption) (*PodPauseLibpodOK, error)

	PodPruneLibpod(params *PodPruneLibpodParams, opts ...ClientOption) (*PodPruneLibpodOK, error)

	PodRestartLibpod(params *PodRestartLibpodParams, opts ...ClientOption) (*PodRestartLibpodOK, error)

	PodStartLibpod(params *PodStartLibpodParams, opts ...ClientOption) (*PodStartLibpodOK, error)

	PodStatsAllLibpod(params *PodStatsAllLibpodParams, opts ...ClientOption) (*PodStatsAllLibpodOK, error)

	PodStopLibpod(params *PodStopLibpodParams, opts ...ClientOption) (*PodStopLibpodOK, error)

	PodTopLibpod(params *PodTopLibpodParams, opts ...ClientOption) (*PodTopLibpodOK, error)

	PodUnpauseLibpod(params *PodUnpauseLibpodParams, opts ...ClientOption) (*PodUnpauseLibpodOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PodCreateLibpod creates a pod
*/
func (a *Client) PodCreateLibpod(params *PodCreateLibpodParams, opts ...ClientOption) (*PodCreateLibpodCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodCreateLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodCreateLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/pods/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodCreateLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodCreateLibpodCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodCreateLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodDeleteLibpod removes pod
*/
func (a *Client) PodDeleteLibpod(params *PodDeleteLibpodParams, opts ...ClientOption) (*PodDeleteLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodDeleteLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodDeleteLibpod",
		Method:             "DELETE",
		PathPattern:        "/libpod/pods/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodDeleteLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodDeleteLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodDeleteLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodExistsLibpod pods exists

Check if a pod exists by name or ID
*/
func (a *Client) PodExistsLibpod(params *PodExistsLibpodParams, opts ...ClientOption) (*PodExistsLibpodNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodExistsLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodExistsLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/pods/{name}/exists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodExistsLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodExistsLibpodNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodExistsLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodInspectLibpod inspects pod
*/
func (a *Client) PodInspectLibpod(params *PodInspectLibpodParams, opts ...ClientOption) (*PodInspectLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodInspectLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodInspectLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/pods/{name}/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodInspectLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodInspectLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodInspectLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodKillLibpod kills a pod
*/
func (a *Client) PodKillLibpod(params *PodKillLibpodParams, opts ...ClientOption) (*PodKillLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodKillLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodKillLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/pods/{name}/kill",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodKillLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodKillLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodKillLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodListLibpod lists pods
*/
func (a *Client) PodListLibpod(params *PodListLibpodParams, opts ...ClientOption) (*PodListLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodListLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodListLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/pods/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodListLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodListLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodListLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodPauseLibpod pauses a pod

Pause a pod
*/
func (a *Client) PodPauseLibpod(params *PodPauseLibpodParams, opts ...ClientOption) (*PodPauseLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodPauseLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodPauseLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/pods/{name}/pause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodPauseLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodPauseLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodPauseLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodPruneLibpod prunes unused pods
*/
func (a *Client) PodPruneLibpod(params *PodPruneLibpodParams, opts ...ClientOption) (*PodPruneLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodPruneLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodPruneLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/pods/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodPruneLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodPruneLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodPruneLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodRestartLibpod restarts a pod
*/
func (a *Client) PodRestartLibpod(params *PodRestartLibpodParams, opts ...ClientOption) (*PodRestartLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodRestartLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodRestartLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/pods/{name}/restart",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodRestartLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodRestartLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodRestartLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodStartLibpod starts a pod
*/
func (a *Client) PodStartLibpod(params *PodStartLibpodParams, opts ...ClientOption) (*PodStartLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodStartLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodStartLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/pods/{name}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodStartLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodStartLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodStartLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodStatsAllLibpod gets stats for one or more pods

Display a live stream of resource usage statistics for the containers in one or more pods
*/
func (a *Client) PodStatsAllLibpod(params *PodStatsAllLibpodParams, opts ...ClientOption) (*PodStatsAllLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodStatsAllLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodStatsAllLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/pods/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodStatsAllLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodStatsAllLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodStatsAllLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodStopLibpod stops a pod
*/
func (a *Client) PodStopLibpod(params *PodStopLibpodParams, opts ...ClientOption) (*PodStopLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodStopLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodStopLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/pods/{name}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodStopLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodStopLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodStopLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodTopLibpod lists processes

List processes running inside a pod
*/
func (a *Client) PodTopLibpod(params *PodTopLibpodParams, opts ...ClientOption) (*PodTopLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodTopLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodTopLibpod",
		Method:             "GET",
		PathPattern:        "/libpod/pods/{name}/top",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodTopLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodTopLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodTopLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PodUnpauseLibpod unpauses a pod
*/
func (a *Client) PodUnpauseLibpod(params *PodUnpauseLibpodParams, opts ...ClientOption) (*PodUnpauseLibpodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPodUnpauseLibpodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PodUnpauseLibpod",
		Method:             "POST",
		PathPattern:        "/libpod/pods/{name}/unpause",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PodUnpauseLibpodReader{formats: a.formats},
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
	success, ok := result.(*PodUnpauseLibpodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PodUnpauseLibpod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
