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
	"github.com/go-openapi/swag"
)

// NewPodDeleteLibpodParams creates a new PodDeleteLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPodDeleteLibpodParams() *PodDeleteLibpodParams {
	return &PodDeleteLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPodDeleteLibpodParamsWithTimeout creates a new PodDeleteLibpodParams object
// with the ability to set a timeout on a request.
func NewPodDeleteLibpodParamsWithTimeout(timeout time.Duration) *PodDeleteLibpodParams {
	return &PodDeleteLibpodParams{
		timeout: timeout,
	}
}

// NewPodDeleteLibpodParamsWithContext creates a new PodDeleteLibpodParams object
// with the ability to set a context for a request.
func NewPodDeleteLibpodParamsWithContext(ctx context.Context) *PodDeleteLibpodParams {
	return &PodDeleteLibpodParams{
		Context: ctx,
	}
}

// NewPodDeleteLibpodParamsWithHTTPClient creates a new PodDeleteLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewPodDeleteLibpodParamsWithHTTPClient(client *http.Client) *PodDeleteLibpodParams {
	return &PodDeleteLibpodParams{
		HTTPClient: client,
	}
}

/* PodDeleteLibpodParams contains all the parameters to send to the API endpoint
   for the pod delete libpod operation.

   Typically these are written to a http.Request.
*/
type PodDeleteLibpodParams struct {

	/* Force.

	   force removal of a running pod by first stopping all containers, then removing all containers in the pod
	*/
	Force *bool

	/* Name.

	   the name or ID of the pod
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pod delete libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodDeleteLibpodParams) WithDefaults() *PodDeleteLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pod delete libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodDeleteLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pod delete libpod params
func (o *PodDeleteLibpodParams) WithTimeout(timeout time.Duration) *PodDeleteLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod delete libpod params
func (o *PodDeleteLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod delete libpod params
func (o *PodDeleteLibpodParams) WithContext(ctx context.Context) *PodDeleteLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod delete libpod params
func (o *PodDeleteLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod delete libpod params
func (o *PodDeleteLibpodParams) WithHTTPClient(client *http.Client) *PodDeleteLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod delete libpod params
func (o *PodDeleteLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the pod delete libpod params
func (o *PodDeleteLibpodParams) WithForce(force *bool) *PodDeleteLibpodParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the pod delete libpod params
func (o *PodDeleteLibpodParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the pod delete libpod params
func (o *PodDeleteLibpodParams) WithName(name string) *PodDeleteLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pod delete libpod params
func (o *PodDeleteLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PodDeleteLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
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
