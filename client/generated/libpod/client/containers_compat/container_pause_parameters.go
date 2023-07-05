// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

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

// NewContainerPauseParams creates a new ContainerPauseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerPauseParams() *ContainerPauseParams {
	return &ContainerPauseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerPauseParamsWithTimeout creates a new ContainerPauseParams object
// with the ability to set a timeout on a request.
func NewContainerPauseParamsWithTimeout(timeout time.Duration) *ContainerPauseParams {
	return &ContainerPauseParams{
		timeout: timeout,
	}
}

// NewContainerPauseParamsWithContext creates a new ContainerPauseParams object
// with the ability to set a context for a request.
func NewContainerPauseParamsWithContext(ctx context.Context) *ContainerPauseParams {
	return &ContainerPauseParams{
		Context: ctx,
	}
}

// NewContainerPauseParamsWithHTTPClient creates a new ContainerPauseParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerPauseParamsWithHTTPClient(client *http.Client) *ContainerPauseParams {
	return &ContainerPauseParams{
		HTTPClient: client,
	}
}

/*
ContainerPauseParams contains all the parameters to send to the API endpoint

	for the container pause operation.

	Typically these are written to a http.Request.
*/
type ContainerPauseParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container pause params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerPauseParams) WithDefaults() *ContainerPauseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container pause params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerPauseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container pause params
func (o *ContainerPauseParams) WithTimeout(timeout time.Duration) *ContainerPauseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container pause params
func (o *ContainerPauseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container pause params
func (o *ContainerPauseParams) WithContext(ctx context.Context) *ContainerPauseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container pause params
func (o *ContainerPauseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container pause params
func (o *ContainerPauseParams) WithHTTPClient(client *http.Client) *ContainerPauseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container pause params
func (o *ContainerPauseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the container pause params
func (o *ContainerPauseParams) WithName(name string) *ContainerPauseParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container pause params
func (o *ContainerPauseParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerPauseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
