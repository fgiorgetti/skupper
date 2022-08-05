// Code generated by go-swagger; DO NOT EDIT.

package volumes

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

// NewVolumeExistsLibpodParams creates a new VolumeExistsLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeExistsLibpodParams() *VolumeExistsLibpodParams {
	return &VolumeExistsLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeExistsLibpodParamsWithTimeout creates a new VolumeExistsLibpodParams object
// with the ability to set a timeout on a request.
func NewVolumeExistsLibpodParamsWithTimeout(timeout time.Duration) *VolumeExistsLibpodParams {
	return &VolumeExistsLibpodParams{
		timeout: timeout,
	}
}

// NewVolumeExistsLibpodParamsWithContext creates a new VolumeExistsLibpodParams object
// with the ability to set a context for a request.
func NewVolumeExistsLibpodParamsWithContext(ctx context.Context) *VolumeExistsLibpodParams {
	return &VolumeExistsLibpodParams{
		Context: ctx,
	}
}

// NewVolumeExistsLibpodParamsWithHTTPClient creates a new VolumeExistsLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeExistsLibpodParamsWithHTTPClient(client *http.Client) *VolumeExistsLibpodParams {
	return &VolumeExistsLibpodParams{
		HTTPClient: client,
	}
}

/* VolumeExistsLibpodParams contains all the parameters to send to the API endpoint
   for the volume exists libpod operation.

   Typically these are written to a http.Request.
*/
type VolumeExistsLibpodParams struct {

	/* Name.

	   the name of the volume
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume exists libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeExistsLibpodParams) WithDefaults() *VolumeExistsLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume exists libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeExistsLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume exists libpod params
func (o *VolumeExistsLibpodParams) WithTimeout(timeout time.Duration) *VolumeExistsLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume exists libpod params
func (o *VolumeExistsLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume exists libpod params
func (o *VolumeExistsLibpodParams) WithContext(ctx context.Context) *VolumeExistsLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume exists libpod params
func (o *VolumeExistsLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume exists libpod params
func (o *VolumeExistsLibpodParams) WithHTTPClient(client *http.Client) *VolumeExistsLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume exists libpod params
func (o *VolumeExistsLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the volume exists libpod params
func (o *VolumeExistsLibpodParams) WithName(name string) *VolumeExistsLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the volume exists libpod params
func (o *VolumeExistsLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeExistsLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
