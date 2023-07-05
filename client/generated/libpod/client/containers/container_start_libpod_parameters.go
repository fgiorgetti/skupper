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

// NewContainerStartLibpodParams creates a new ContainerStartLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerStartLibpodParams() *ContainerStartLibpodParams {
	return &ContainerStartLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerStartLibpodParamsWithTimeout creates a new ContainerStartLibpodParams object
// with the ability to set a timeout on a request.
func NewContainerStartLibpodParamsWithTimeout(timeout time.Duration) *ContainerStartLibpodParams {
	return &ContainerStartLibpodParams{
		timeout: timeout,
	}
}

// NewContainerStartLibpodParamsWithContext creates a new ContainerStartLibpodParams object
// with the ability to set a context for a request.
func NewContainerStartLibpodParamsWithContext(ctx context.Context) *ContainerStartLibpodParams {
	return &ContainerStartLibpodParams{
		Context: ctx,
	}
}

// NewContainerStartLibpodParamsWithHTTPClient creates a new ContainerStartLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerStartLibpodParamsWithHTTPClient(client *http.Client) *ContainerStartLibpodParams {
	return &ContainerStartLibpodParams{
		HTTPClient: client,
	}
}

/*
ContainerStartLibpodParams contains all the parameters to send to the API endpoint

	for the container start libpod operation.

	Typically these are written to a http.Request.
*/
type ContainerStartLibpodParams struct {

	/* DetachKeys.

	   Override the key sequence for detaching a container. Format is a single character [a-Z] or ctrl-<value> where <value> is one of: a-z, @, ^, [, , or _.

	   Default: "ctrl-p,ctrl-q"
	*/
	DetachKeys *string

	/* Name.

	   the name or ID of the container
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container start libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerStartLibpodParams) WithDefaults() *ContainerStartLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container start libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerStartLibpodParams) SetDefaults() {
	var (
		detachKeysDefault = string("ctrl-p,ctrl-q")
	)

	val := ContainerStartLibpodParams{
		DetachKeys: &detachKeysDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the container start libpod params
func (o *ContainerStartLibpodParams) WithTimeout(timeout time.Duration) *ContainerStartLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container start libpod params
func (o *ContainerStartLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container start libpod params
func (o *ContainerStartLibpodParams) WithContext(ctx context.Context) *ContainerStartLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container start libpod params
func (o *ContainerStartLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container start libpod params
func (o *ContainerStartLibpodParams) WithHTTPClient(client *http.Client) *ContainerStartLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container start libpod params
func (o *ContainerStartLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDetachKeys adds the detachKeys to the container start libpod params
func (o *ContainerStartLibpodParams) WithDetachKeys(detachKeys *string) *ContainerStartLibpodParams {
	o.SetDetachKeys(detachKeys)
	return o
}

// SetDetachKeys adds the detachKeys to the container start libpod params
func (o *ContainerStartLibpodParams) SetDetachKeys(detachKeys *string) {
	o.DetachKeys = detachKeys
}

// WithName adds the name to the container start libpod params
func (o *ContainerStartLibpodParams) WithName(name string) *ContainerStartLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the container start libpod params
func (o *ContainerStartLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerStartLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DetachKeys != nil {

		// query param detachKeys
		var qrDetachKeys string

		if o.DetachKeys != nil {
			qrDetachKeys = *o.DetachKeys
		}
		qDetachKeys := qrDetachKeys
		if qDetachKeys != "" {

			if err := r.SetQueryParam("detachKeys", qDetachKeys); err != nil {
				return err
			}
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
