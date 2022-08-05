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

// NewPodKillLibpodParams creates a new PodKillLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPodKillLibpodParams() *PodKillLibpodParams {
	return &PodKillLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPodKillLibpodParamsWithTimeout creates a new PodKillLibpodParams object
// with the ability to set a timeout on a request.
func NewPodKillLibpodParamsWithTimeout(timeout time.Duration) *PodKillLibpodParams {
	return &PodKillLibpodParams{
		timeout: timeout,
	}
}

// NewPodKillLibpodParamsWithContext creates a new PodKillLibpodParams object
// with the ability to set a context for a request.
func NewPodKillLibpodParamsWithContext(ctx context.Context) *PodKillLibpodParams {
	return &PodKillLibpodParams{
		Context: ctx,
	}
}

// NewPodKillLibpodParamsWithHTTPClient creates a new PodKillLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewPodKillLibpodParamsWithHTTPClient(client *http.Client) *PodKillLibpodParams {
	return &PodKillLibpodParams{
		HTTPClient: client,
	}
}

/* PodKillLibpodParams contains all the parameters to send to the API endpoint
   for the pod kill libpod operation.

   Typically these are written to a http.Request.
*/
type PodKillLibpodParams struct {

	/* Name.

	   the name or ID of the pod
	*/
	Name string

	/* Signal.

	   signal to be sent to pod

	   Default: "SIGKILL"
	*/
	Signal *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pod kill libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodKillLibpodParams) WithDefaults() *PodKillLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pod kill libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PodKillLibpodParams) SetDefaults() {
	var (
		signalDefault = string("SIGKILL")
	)

	val := PodKillLibpodParams{
		Signal: &signalDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the pod kill libpod params
func (o *PodKillLibpodParams) WithTimeout(timeout time.Duration) *PodKillLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pod kill libpod params
func (o *PodKillLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pod kill libpod params
func (o *PodKillLibpodParams) WithContext(ctx context.Context) *PodKillLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pod kill libpod params
func (o *PodKillLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pod kill libpod params
func (o *PodKillLibpodParams) WithHTTPClient(client *http.Client) *PodKillLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pod kill libpod params
func (o *PodKillLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the pod kill libpod params
func (o *PodKillLibpodParams) WithName(name string) *PodKillLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the pod kill libpod params
func (o *PodKillLibpodParams) SetName(name string) {
	o.Name = name
}

// WithSignal adds the signal to the pod kill libpod params
func (o *PodKillLibpodParams) WithSignal(signal *string) *PodKillLibpodParams {
	o.SetSignal(signal)
	return o
}

// SetSignal adds the signal to the pod kill libpod params
func (o *PodKillLibpodParams) SetSignal(signal *string) {
	o.Signal = signal
}

// WriteToRequest writes these params to a swagger request
func (o *PodKillLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Signal != nil {

		// query param signal
		var qrSignal string

		if o.Signal != nil {
			qrSignal = *o.Signal
		}
		qSignal := qrSignal
		if qSignal != "" {

			if err := r.SetQueryParam("signal", qSignal); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
