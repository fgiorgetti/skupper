// Code generated by go-swagger; DO NOT EDIT.

package system_compat

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

// NewSystemEventsParams creates a new SystemEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemEventsParams() *SystemEventsParams {
	return &SystemEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemEventsParamsWithTimeout creates a new SystemEventsParams object
// with the ability to set a timeout on a request.
func NewSystemEventsParamsWithTimeout(timeout time.Duration) *SystemEventsParams {
	return &SystemEventsParams{
		timeout: timeout,
	}
}

// NewSystemEventsParamsWithContext creates a new SystemEventsParams object
// with the ability to set a context for a request.
func NewSystemEventsParamsWithContext(ctx context.Context) *SystemEventsParams {
	return &SystemEventsParams{
		Context: ctx,
	}
}

// NewSystemEventsParamsWithHTTPClient creates a new SystemEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemEventsParamsWithHTTPClient(client *http.Client) *SystemEventsParams {
	return &SystemEventsParams{
		HTTPClient: client,
	}
}

/* SystemEventsParams contains all the parameters to send to the API endpoint
   for the system events operation.

   Typically these are written to a http.Request.
*/
type SystemEventsParams struct {

	/* Filters.

	   JSON encoded map[string][]string of constraints
	*/
	Filters *string

	/* Since.

	   start streaming events from this time
	*/
	Since *string

	/* Until.

	   stop streaming events later than this
	*/
	Until *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemEventsParams) WithDefaults() *SystemEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system events params
func (o *SystemEventsParams) WithTimeout(timeout time.Duration) *SystemEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system events params
func (o *SystemEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system events params
func (o *SystemEventsParams) WithContext(ctx context.Context) *SystemEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system events params
func (o *SystemEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system events params
func (o *SystemEventsParams) WithHTTPClient(client *http.Client) *SystemEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system events params
func (o *SystemEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the system events params
func (o *SystemEventsParams) WithFilters(filters *string) *SystemEventsParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the system events params
func (o *SystemEventsParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithSince adds the since to the system events params
func (o *SystemEventsParams) WithSince(since *string) *SystemEventsParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the system events params
func (o *SystemEventsParams) SetSince(since *string) {
	o.Since = since
}

// WithUntil adds the until to the system events params
func (o *SystemEventsParams) WithUntil(until *string) *SystemEventsParams {
	o.SetUntil(until)
	return o
}

// SetUntil adds the until to the system events params
func (o *SystemEventsParams) SetUntil(until *string) {
	o.Until = until
}

// WriteToRequest writes these params to a swagger request
func (o *SystemEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string

		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {

			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}
	}

	if o.Since != nil {

		// query param since
		var qrSince string

		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := qrSince
		if qSince != "" {

			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}
	}

	if o.Until != nil {

		// query param until
		var qrUntil string

		if o.Until != nil {
			qrUntil = *o.Until
		}
		qUntil := qrUntil
		if qUntil != "" {

			if err := r.SetQueryParam("until", qUntil); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
