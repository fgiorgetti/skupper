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

// NewContainerExportParams creates a new ContainerExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerExportParams() *ContainerExportParams {
	return &ContainerExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerExportParamsWithTimeout creates a new ContainerExportParams object
// with the ability to set a timeout on a request.
func NewContainerExportParamsWithTimeout(timeout time.Duration) *ContainerExportParams {
	return &ContainerExportParams{
		timeout: timeout,
	}
}

// NewContainerExportParamsWithContext creates a new ContainerExportParams object
// with the ability to set a context for a request.
func NewContainerExportParamsWithContext(ctx context.Context) *ContainerExportParams {
	return &ContainerExportParams{
		Context: ctx,
	}
}

// NewContainerExportParamsWithHTTPClient creates a new ContainerExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerExportParamsWithHTTPClient(client *http.Client) *ContainerExportParams {
	return &ContainerExportParams{
		HTTPClient: client,
	}
}

/*
ContainerExportParams contains all the parameters to send to the API endpoint

	for the container export operation.

	Typically these are written to a http.Request.
*/
type ContainerExportParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerExportParams) WithDefaults() *ContainerExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerExportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container export params
func (o *ContainerExportParams) WithTimeout(timeout time.Duration) *ContainerExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container export params
func (o *ContainerExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container export params
func (o *ContainerExportParams) WithContext(ctx context.Context) *ContainerExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container export params
func (o *ContainerExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container export params
func (o *ContainerExportParams) WithHTTPClient(client *http.Client) *ContainerExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container export params
func (o *ContainerExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the container export params
func (o *ContainerExportParams) WithName(name string) *ContainerExportParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container export params
func (o *ContainerExportParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
