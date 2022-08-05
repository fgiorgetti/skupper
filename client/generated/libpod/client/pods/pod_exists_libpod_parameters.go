// Code generated by go-swagger; DO NOT EDIT.

package pods

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

// NewPodExistsLibpodParams creates a new PodExistsLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPodExistsLibpodParams() *PodExistsLibpodParams {
	return &PodExistsLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPodExistsLibpodParamsWithTimeout creates a new PodExistsLibpodParams object
// with the ability to set a timeout on a request.
func NewPodExistsLibpodParamsWithTimeout(timeout time.Duration) *PodExistsLibpodParams {
	return &PodExistsLibpodParams{
		timeout: timeout,
	}
}

// NewPodExistsLibpodParamsWithContext creates a new PodExistsLibpodParams object
// with the ability to set a context for a request.
func NewPodExistsLibpodParamsWithContext(ctx context.Context) *PodExistsLibpodParams {
	return &PodExistsLibpodParams{
		Context: ctx,
	}
}

// NewPodExistsLibpodParamsWithHTTPClient creates a new PodExistsLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewPodExistsLibpodParamsWithHTTPClient(client *http.Client) *PodExistsLibpodParams {
	return &PodExistsLibpodParams{
		HTTPClient: client,
	}
}

/* PodExistsLibpodParams contains all the parameters to send to the API endpoint
   for the pod exists libpod operation.

   Typically these are written to a http.Request.
*/
type PodExistsLibpodParams struct {

	/* Name.

	   the name or ID of the pod
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pod exists libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodExistsLibpodParams) WithDefaults() *PodExistsLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pod exists libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodExistsLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pod exists libpod params
func (o *PodExistsLibpodParams) WithTimeout(timeout time.Duration) *PodExistsLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod exists libpod params
func (o *PodExistsLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod exists libpod params
func (o *PodExistsLibpodParams) WithContext(ctx context.Context) *PodExistsLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod exists libpod params
func (o *PodExistsLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod exists libpod params
func (o *PodExistsLibpodParams) WithHTTPClient(client *http.Client) *PodExistsLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod exists libpod params
func (o *PodExistsLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the pod exists libpod params
func (o *PodExistsLibpodParams) WithName(name string) *PodExistsLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pod exists libpod params
func (o *PodExistsLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PodExistsLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
