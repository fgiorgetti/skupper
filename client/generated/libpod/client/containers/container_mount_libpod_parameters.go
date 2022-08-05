// Code generated by go-swagger; DO NOT EDIT.

package containers

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
)

// NewContainerMountLibpodParams creates a new ContainerMountLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerMountLibpodParams() *ContainerMountLibpodParams {
	return &ContainerMountLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerMountLibpodParamsWithTimeout creates a new ContainerMountLibpodParams object
// with the ability to set a timeout on a request.
func NewContainerMountLibpodParamsWithTimeout(timeout time.Duration) *ContainerMountLibpodParams {
	return &ContainerMountLibpodParams{
		timeout: timeout,
	}
}

// NewContainerMountLibpodParamsWithContext creates a new ContainerMountLibpodParams object
// with the ability to set a context for a request.
func NewContainerMountLibpodParamsWithContext(ctx context.Context) *ContainerMountLibpodParams {
	return &ContainerMountLibpodParams{
		Context: ctx,
	}
}

// NewContainerMountLibpodParamsWithHTTPClient creates a new ContainerMountLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerMountLibpodParamsWithHTTPClient(client *http.Client) *ContainerMountLibpodParams {
	return &ContainerMountLibpodParams{
		HTTPClient: client,
	}
}

/* ContainerMountLibpodParams contains all the parameters to send to the API endpoint
   for the container mount libpod operation.

   Typically these are written to a http.Request.
*/
type ContainerMountLibpodParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container mount libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerMountLibpodParams) WithDefaults() *ContainerMountLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container mount libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerMountLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container mount libpod params
func (o *ContainerMountLibpodParams) WithTimeout(timeout time.Duration) *ContainerMountLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container mount libpod params
func (o *ContainerMountLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container mount libpod params
func (o *ContainerMountLibpodParams) WithContext(ctx context.Context) *ContainerMountLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container mount libpod params
func (o *ContainerMountLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container mount libpod params
func (o *ContainerMountLibpodParams) WithHTTPClient(client *http.Client) *ContainerMountLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container mount libpod params
func (o *ContainerMountLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the container mount libpod params
func (o *ContainerMountLibpodParams) WithName(name string) *ContainerMountLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container mount libpod params
func (o *ContainerMountLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerMountLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
