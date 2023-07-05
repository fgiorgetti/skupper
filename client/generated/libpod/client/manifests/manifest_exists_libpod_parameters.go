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

// NewManifestExistsLibpodParams creates a new ManifestExistsLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManifestExistsLibpodParams() *ManifestExistsLibpodParams {
	return &ManifestExistsLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManifestExistsLibpodParamsWithTimeout creates a new ManifestExistsLibpodParams object
// with the ability to set a timeout on a request.
func NewManifestExistsLibpodParamsWithTimeout(timeout time.Duration) *ManifestExistsLibpodParams {
	return &ManifestExistsLibpodParams{
		timeout: timeout,
	}
}

// NewManifestExistsLibpodParamsWithContext creates a new ManifestExistsLibpodParams object
// with the ability to set a context for a request.
func NewManifestExistsLibpodParamsWithContext(ctx context.Context) *ManifestExistsLibpodParams {
	return &ManifestExistsLibpodParams{
		Context: ctx,
	}
}

// NewManifestExistsLibpodParamsWithHTTPClient creates a new ManifestExistsLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewManifestExistsLibpodParamsWithHTTPClient(client *http.Client) *ManifestExistsLibpodParams {
	return &ManifestExistsLibpodParams{
		HTTPClient: client,
	}
}

/*
ManifestExistsLibpodParams contains all the parameters to send to the API endpoint

	for the manifest exists libpod operation.

	Typically these are written to a http.Request.
*/
type ManifestExistsLibpodParams struct {

	/* Name.

	   the name or ID of the manifest list
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the manifest exists libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManifestExistsLibpodParams) WithDefaults() *ManifestExistsLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the manifest exists libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManifestExistsLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the manifest exists libpod params
func (o *ManifestExistsLibpodParams) WithTimeout(timeout time.Duration) *ManifestExistsLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the manifest exists libpod params
func (o *ManifestExistsLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the manifest exists libpod params
func (o *ManifestExistsLibpodParams) WithContext(ctx context.Context) *ManifestExistsLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the manifest exists libpod params
func (o *ManifestExistsLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the manifest exists libpod params
func (o *ManifestExistsLibpodParams) WithHTTPClient(client *http.Client) *ManifestExistsLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the manifest exists libpod params
func (o *ManifestExistsLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the manifest exists libpod params
func (o *ManifestExistsLibpodParams) WithName(name string) *ManifestExistsLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the manifest exists libpod params
func (o *ManifestExistsLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ManifestExistsLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
