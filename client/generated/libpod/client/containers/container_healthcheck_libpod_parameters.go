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

// NewContainerHealthcheckLibpodParams creates a new ContainerHealthcheckLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerHealthcheckLibpodParams() *ContainerHealthcheckLibpodParams {
	return &ContainerHealthcheckLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerHealthcheckLibpodParamsWithTimeout creates a new ContainerHealthcheckLibpodParams object
// with the ability to set a timeout on a request.
func NewContainerHealthcheckLibpodParamsWithTimeout(timeout time.Duration) *ContainerHealthcheckLibpodParams {
	return &ContainerHealthcheckLibpodParams{
		timeout: timeout,
	}
}

// NewContainerHealthcheckLibpodParamsWithContext creates a new ContainerHealthcheckLibpodParams object
// with the ability to set a context for a request.
func NewContainerHealthcheckLibpodParamsWithContext(ctx context.Context) *ContainerHealthcheckLibpodParams {
	return &ContainerHealthcheckLibpodParams{
		Context: ctx,
	}
}

// NewContainerHealthcheckLibpodParamsWithHTTPClient creates a new ContainerHealthcheckLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerHealthcheckLibpodParamsWithHTTPClient(client *http.Client) *ContainerHealthcheckLibpodParams {
	return &ContainerHealthcheckLibpodParams{
		HTTPClient: client,
	}
}

/*
ContainerHealthcheckLibpodParams contains all the parameters to send to the API endpoint

	for the container healthcheck libpod operation.

	Typically these are written to a http.Request.
*/
type ContainerHealthcheckLibpodParams struct {

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container healthcheck libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerHealthcheckLibpodParams) WithDefaults() *ContainerHealthcheckLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container healthcheck libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerHealthcheckLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container healthcheck libpod params
func (o *ContainerHealthcheckLibpodParams) WithTimeout(timeout time.Duration) *ContainerHealthcheckLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container healthcheck libpod params
func (o *ContainerHealthcheckLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container healthcheck libpod params
func (o *ContainerHealthcheckLibpodParams) WithContext(ctx context.Context) *ContainerHealthcheckLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container healthcheck libpod params
func (o *ContainerHealthcheckLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container healthcheck libpod params
func (o *ContainerHealthcheckLibpodParams) WithHTTPClient(client *http.Client) *ContainerHealthcheckLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container healthcheck libpod params
func (o *ContainerHealthcheckLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the container healthcheck libpod params
func (o *ContainerHealthcheckLibpodParams) WithName(name string) *ContainerHealthcheckLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container healthcheck libpod params
func (o *ContainerHealthcheckLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerHealthcheckLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
