// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewSystemDataUsageLibpodParams creates a new SystemDataUsageLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemDataUsageLibpodParams() *SystemDataUsageLibpodParams {
	return &SystemDataUsageLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemDataUsageLibpodParamsWithTimeout creates a new SystemDataUsageLibpodParams object
// with the ability to set a timeout on a request.
func NewSystemDataUsageLibpodParamsWithTimeout(timeout time.Duration) *SystemDataUsageLibpodParams {
	return &SystemDataUsageLibpodParams{
		timeout: timeout,
	}
}

// NewSystemDataUsageLibpodParamsWithContext creates a new SystemDataUsageLibpodParams object
// with the ability to set a context for a request.
func NewSystemDataUsageLibpodParamsWithContext(ctx context.Context) *SystemDataUsageLibpodParams {
	return &SystemDataUsageLibpodParams{
		Context: ctx,
	}
}

// NewSystemDataUsageLibpodParamsWithHTTPClient creates a new SystemDataUsageLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemDataUsageLibpodParamsWithHTTPClient(client *http.Client) *SystemDataUsageLibpodParams {
	return &SystemDataUsageLibpodParams{
		HTTPClient: client,
	}
}

/* SystemDataUsageLibpodParams contains all the parameters to send to the API endpoint
   for the system data usage libpod operation.

   Typically these are written to a http.Request.
*/
type SystemDataUsageLibpodParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system data usage libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemDataUsageLibpodParams) WithDefaults() *SystemDataUsageLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system data usage libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemDataUsageLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system data usage libpod params
func (o *SystemDataUsageLibpodParams) WithTimeout(timeout time.Duration) *SystemDataUsageLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system data usage libpod params
func (o *SystemDataUsageLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system data usage libpod params
func (o *SystemDataUsageLibpodParams) WithContext(ctx context.Context) *SystemDataUsageLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system data usage libpod params
func (o *SystemDataUsageLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system data usage libpod params
func (o *SystemDataUsageLibpodParams) WithHTTPClient(client *http.Client) *SystemDataUsageLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system data usage libpod params
func (o *SystemDataUsageLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SystemDataUsageLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
