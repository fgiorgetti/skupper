// Code generated by go-swagger; DO NOT EDIT.

package networks_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

// NewNetworkDisconnectParams creates a new NetworkDisconnectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkDisconnectParams() *NetworkDisconnectParams {
	return &NetworkDisconnectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkDisconnectParamsWithTimeout creates a new NetworkDisconnectParams object
// with the ability to set a timeout on a request.
func NewNetworkDisconnectParamsWithTimeout(timeout time.Duration) *NetworkDisconnectParams {
	return &NetworkDisconnectParams{
		timeout: timeout,
	}
}

// NewNetworkDisconnectParamsWithContext creates a new NetworkDisconnectParams object
// with the ability to set a context for a request.
func NewNetworkDisconnectParamsWithContext(ctx context.Context) *NetworkDisconnectParams {
	return &NetworkDisconnectParams{
		Context: ctx,
	}
}

// NewNetworkDisconnectParamsWithHTTPClient creates a new NetworkDisconnectParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkDisconnectParamsWithHTTPClient(client *http.Client) *NetworkDisconnectParams {
	return &NetworkDisconnectParams{
		HTTPClient: client,
	}
}

/* NetworkDisconnectParams contains all the parameters to send to the API endpoint
   for the network disconnect operation.

   Typically these are written to a http.Request.
*/
type NetworkDisconnectParams struct {

	/* Create.

	   attributes for disconnecting a container from a network
	*/
	Create *models.SwagCompatNetworkDisconnectRequest

	/* Name.

	   the name of the network
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network disconnect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkDisconnectParams) WithDefaults() *NetworkDisconnectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network disconnect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkDisconnectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network disconnect params
func (o *NetworkDisconnectParams) WithTimeout(timeout time.Duration) *NetworkDisconnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network disconnect params
func (o *NetworkDisconnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network disconnect params
func (o *NetworkDisconnectParams) WithContext(ctx context.Context) *NetworkDisconnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network disconnect params
func (o *NetworkDisconnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network disconnect params
func (o *NetworkDisconnectParams) WithHTTPClient(client *http.Client) *NetworkDisconnectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network disconnect params
func (o *NetworkDisconnectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreate adds the create to the network disconnect params
func (o *NetworkDisconnectParams) WithCreate(create *models.SwagCompatNetworkDisconnectRequest) *NetworkDisconnectParams {
	o.SetCreate(create)
	return o
}

// SetCreate adds the create to the network disconnect params
func (o *NetworkDisconnectParams) SetCreate(create *models.SwagCompatNetworkDisconnectRequest) {
	o.Create = create
}

// WithName adds the name to the network disconnect params
func (o *NetworkDisconnectParams) WithName(name string) *NetworkDisconnectParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the network disconnect params
func (o *NetworkDisconnectParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkDisconnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Create != nil {
		if err := r.SetBodyParam(o.Create); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
