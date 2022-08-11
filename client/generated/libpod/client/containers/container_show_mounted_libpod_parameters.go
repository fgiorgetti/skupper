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

// NewContainerShowMountedLibpodParams creates a new ContainerShowMountedLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerShowMountedLibpodParams() *ContainerShowMountedLibpodParams {
	return &ContainerShowMountedLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerShowMountedLibpodParamsWithTimeout creates a new ContainerShowMountedLibpodParams object
// with the ability to set a timeout on a request.
func NewContainerShowMountedLibpodParamsWithTimeout(timeout time.Duration) *ContainerShowMountedLibpodParams {
	return &ContainerShowMountedLibpodParams{
		timeout: timeout,
	}
}

// NewContainerShowMountedLibpodParamsWithContext creates a new ContainerShowMountedLibpodParams object
// with the ability to set a context for a request.
func NewContainerShowMountedLibpodParamsWithContext(ctx context.Context) *ContainerShowMountedLibpodParams {
	return &ContainerShowMountedLibpodParams{
		Context: ctx,
	}
}

// NewContainerShowMountedLibpodParamsWithHTTPClient creates a new ContainerShowMountedLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerShowMountedLibpodParamsWithHTTPClient(client *http.Client) *ContainerShowMountedLibpodParams {
	return &ContainerShowMountedLibpodParams{
		HTTPClient: client,
	}
}

/* ContainerShowMountedLibpodParams contains all the parameters to send to the API endpoint
   for the container show mounted libpod operation.

   Typically these are written to a http.Request.
*/
type ContainerShowMountedLibpodParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container show mounted libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerShowMountedLibpodParams) WithDefaults() *ContainerShowMountedLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container show mounted libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerShowMountedLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container show mounted libpod params
func (o *ContainerShowMountedLibpodParams) WithTimeout(timeout time.Duration) *ContainerShowMountedLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container show mounted libpod params
func (o *ContainerShowMountedLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container show mounted libpod params
func (o *ContainerShowMountedLibpodParams) WithContext(ctx context.Context) *ContainerShowMountedLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container show mounted libpod params
func (o *ContainerShowMountedLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container show mounted libpod params
func (o *ContainerShowMountedLibpodParams) WithHTTPClient(client *http.Client) *ContainerShowMountedLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container show mounted libpod params
func (o *ContainerShowMountedLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerShowMountedLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}