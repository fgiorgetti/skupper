// Code generated by go-swagger; DO NOT EDIT.

package manifests

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

// NewManifestInspectLibpodParams creates a new ManifestInspectLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManifestInspectLibpodParams() *ManifestInspectLibpodParams {
	return &ManifestInspectLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManifestInspectLibpodParamsWithTimeout creates a new ManifestInspectLibpodParams object
// with the ability to set a timeout on a request.
func NewManifestInspectLibpodParamsWithTimeout(timeout time.Duration) *ManifestInspectLibpodParams {
	return &ManifestInspectLibpodParams{
		timeout: timeout,
	}
}

// NewManifestInspectLibpodParamsWithContext creates a new ManifestInspectLibpodParams object
// with the ability to set a context for a request.
func NewManifestInspectLibpodParamsWithContext(ctx context.Context) *ManifestInspectLibpodParams {
	return &ManifestInspectLibpodParams{
		Context: ctx,
	}
}

// NewManifestInspectLibpodParamsWithHTTPClient creates a new ManifestInspectLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewManifestInspectLibpodParamsWithHTTPClient(client *http.Client) *ManifestInspectLibpodParams {
	return &ManifestInspectLibpodParams{
		HTTPClient: client,
	}
}

/* ManifestInspectLibpodParams contains all the parameters to send to the API endpoint
   for the manifest inspect libpod operation.

   Typically these are written to a http.Request.
*/
type ManifestInspectLibpodParams struct {

	/* Name.

	   the name or ID of the manifest
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the manifest inspect libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManifestInspectLibpodParams) WithDefaults() *ManifestInspectLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the manifest inspect libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManifestInspectLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the manifest inspect libpod params
func (o *ManifestInspectLibpodParams) WithTimeout(timeout time.Duration) *ManifestInspectLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the manifest inspect libpod params
func (o *ManifestInspectLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the manifest inspect libpod params
func (o *ManifestInspectLibpodParams) WithContext(ctx context.Context) *ManifestInspectLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the manifest inspect libpod params
func (o *ManifestInspectLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the manifest inspect libpod params
func (o *ManifestInspectLibpodParams) WithHTTPClient(client *http.Client) *ManifestInspectLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the manifest inspect libpod params
func (o *ManifestInspectLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the manifest inspect libpod params
func (o *ManifestInspectLibpodParams) WithName(name string) *ManifestInspectLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the manifest inspect libpod params
func (o *ManifestInspectLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ManifestInspectLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
