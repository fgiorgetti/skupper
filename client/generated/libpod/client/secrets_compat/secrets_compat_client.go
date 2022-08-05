// Code generated by go-swagger; DO NOT EDIT.

package secrets_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new secrets compat API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for secrets compat API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SecretCreate(params *SecretCreateParams, opts ...ClientOption) (*SecretCreateCreated, error)

	SecretDelete(params *SecretDeleteParams, opts ...ClientOption) (*SecretDeleteNoContent, error)

	SecretInspect(params *SecretInspectParams, opts ...ClientOption) (*SecretInspectOK, error)

	SecretList(params *SecretListParams, opts ...ClientOption) (*SecretListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SecretCreate creates a secret
*/
func (a *Client) SecretCreate(params *SecretCreateParams, opts ...ClientOption) (*SecretCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SecretCreate",
		Method:             "POST",
		PathPattern:        "/secrets/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SecretCreateReader{formats: a.formats},
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
	success, ok := result.(*SecretCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SecretCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretDelete removes secret
*/
func (a *Client) SecretDelete(params *SecretDeleteParams, opts ...ClientOption) (*SecretDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SecretDelete",
		Method:             "DELETE",
		PathPattern:        "/secrets/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SecretDeleteReader{formats: a.formats},
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
	success, ok := result.(*SecretDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SecretDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretInspect inspects secret
*/
func (a *Client) SecretInspect(params *SecretInspectParams, opts ...ClientOption) (*SecretInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SecretInspect",
		Method:             "GET",
		PathPattern:        "/secrets/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SecretInspectReader{formats: a.formats},
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
	success, ok := result.(*SecretInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SecretInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SecretList lists secrets

  Returns a list of secrets
*/
func (a *Client) SecretList(params *SecretListParams, opts ...ClientOption) (*SecretListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SecretList",
		Method:             "GET",
		PathPattern:        "/secrets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SecretListReader{formats: a.formats},
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
	success, ok := result.(*SecretListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SecretList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
