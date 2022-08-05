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

// NewPodStopLibpodParams creates a new PodStopLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPodStopLibpodParams() *PodStopLibpodParams {
	return &PodStopLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPodStopLibpodParamsWithTimeout creates a new PodStopLibpodParams object
// with the ability to set a timeout on a request.
func NewPodStopLibpodParamsWithTimeout(timeout time.Duration) *PodStopLibpodParams {
	return &PodStopLibpodParams{
		timeout: timeout,
	}
}

// NewPodStopLibpodParamsWithContext creates a new PodStopLibpodParams object
// with the ability to set a context for a request.
func NewPodStopLibpodParamsWithContext(ctx context.Context) *PodStopLibpodParams {
	return &PodStopLibpodParams{
		Context: ctx,
	}
}

// NewPodStopLibpodParamsWithHTTPClient creates a new PodStopLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewPodStopLibpodParamsWithHTTPClient(client *http.Client) *PodStopLibpodParams {
	return &PodStopLibpodParams{
		HTTPClient: client,
	}
}

/* PodStopLibpodParams contains all the parameters to send to the API endpoint
   for the pod stop libpod operation.

   Typically these are written to a http.Request.
*/
type PodStopLibpodParams struct {

	/* Name.

	   the name or ID of the pod
	*/
	Name string

	/* T.

	   timeout
	*/
	T *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pod stop libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodStopLibpodParams) WithDefaults() *PodStopLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pod stop libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodStopLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pod stop libpod params
func (o *PodStopLibpodParams) WithTimeout(timeout time.Duration) *PodStopLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod stop libpod params
func (o *PodStopLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod stop libpod params
func (o *PodStopLibpodParams) WithContext(ctx context.Context) *PodStopLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod stop libpod params
func (o *PodStopLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod stop libpod params
func (o *PodStopLibpodParams) WithHTTPClient(client *http.Client) *PodStopLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod stop libpod params
func (o *PodStopLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the pod stop libpod params
func (o *PodStopLibpodParams) WithName(name string) *PodStopLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pod stop libpod params
func (o *PodStopLibpodParams) SetName(name string) {
	o.Name = name
}

// WithT adds the t to the pod stop libpod params
func (o *PodStopLibpodParams) WithT(t *int64) *PodStopLibpodParams {
	o.SetT(t)
	return o
}

// SetT adds the t to the pod stop libpod params
func (o *PodStopLibpodParams) SetT(t *int64) {
	o.T = t
}

// WriteToRequest writes these params to a swagger request
func (o *PodStopLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.T != nil {

		// query param t
		var qrT int64

		if o.T != nil {
			qrT = *o.T
		}
		qT := swag.FormatInt64(qrT)
		if qT != "" {

			if err := r.SetQueryParam("t", qT); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
