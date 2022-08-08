// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewNetworkListLibpodParams creates a new NetworkListLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkListLibpodParams() *NetworkListLibpodParams {
	return &NetworkListLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkListLibpodParamsWithTimeout creates a new NetworkListLibpodParams object
// with the ability to set a timeout on a request.
func NewNetworkListLibpodParamsWithTimeout(timeout time.Duration) *NetworkListLibpodParams {
	return &NetworkListLibpodParams{
		timeout: timeout,
	}
}

// NewNetworkListLibpodParamsWithContext creates a new NetworkListLibpodParams object
// with the ability to set a context for a request.
func NewNetworkListLibpodParamsWithContext(ctx context.Context) *NetworkListLibpodParams {
	return &NetworkListLibpodParams{
		Context: ctx,
	}
}

// NewNetworkListLibpodParamsWithHTTPClient creates a new NetworkListLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkListLibpodParamsWithHTTPClient(client *http.Client) *NetworkListLibpodParams {
	return &NetworkListLibpodParams{
		HTTPClient: client,
	}
}

/* NetworkListLibpodParams contains all the parameters to send to the API endpoint
   for the network list libpod operation.

   Typically these are written to a http.Request.
*/
type NetworkListLibpodParams struct {

	/* Filters.

	   JSON encoded value of the filters (a `map[string][]string`) to process on the network list. Available filters:
	- `name=[name]` Matches network name (accepts regex).
	- `id=[id]` Matches for full or partial ID.
	- `driver=[driver]` Only bridge is supported.
	- `label=[key]` or `label=[key=value]` Matches networks based on the presence of a label alone or a label and a value.
	- `until=[timestamp]` Matches all networks that were create before the given timestamp.

	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network list libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkListLibpodParams) WithDefaults() *NetworkListLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network list libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkListLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network list libpod params
func (o *NetworkListLibpodParams) WithTimeout(timeout time.Duration) *NetworkListLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network list libpod params
func (o *NetworkListLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network list libpod params
func (o *NetworkListLibpodParams) WithContext(ctx context.Context) *NetworkListLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network list libpod params
func (o *NetworkListLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network list libpod params
func (o *NetworkListLibpodParams) WithHTTPClient(client *http.Client) *NetworkListLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network list libpod params
func (o *NetworkListLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the network list libpod params
func (o *NetworkListLibpodParams) WithFilters(filters *string) *NetworkListLibpodParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the network list libpod params
func (o *NetworkListLibpodParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkListLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
