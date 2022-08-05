// Code generated by go-swagger; DO NOT EDIT.

package images

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

// NewImageDeleteLibpodParams creates a new ImageDeleteLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageDeleteLibpodParams() *ImageDeleteLibpodParams {
	return &ImageDeleteLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageDeleteLibpodParamsWithTimeout creates a new ImageDeleteLibpodParams object
// with the ability to set a timeout on a request.
func NewImageDeleteLibpodParamsWithTimeout(timeout time.Duration) *ImageDeleteLibpodParams {
	return &ImageDeleteLibpodParams{
		timeout: timeout,
	}
}

// NewImageDeleteLibpodParamsWithContext creates a new ImageDeleteLibpodParams object
// with the ability to set a context for a request.
func NewImageDeleteLibpodParamsWithContext(ctx context.Context) *ImageDeleteLibpodParams {
	return &ImageDeleteLibpodParams{
		Context: ctx,
	}
}

// NewImageDeleteLibpodParamsWithHTTPClient creates a new ImageDeleteLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageDeleteLibpodParamsWithHTTPClient(client *http.Client) *ImageDeleteLibpodParams {
	return &ImageDeleteLibpodParams{
		HTTPClient: client,
	}
}

/* ImageDeleteLibpodParams contains all the parameters to send to the API endpoint
   for the image delete libpod operation.

   Typically these are written to a http.Request.
*/
type ImageDeleteLibpodParams struct {

	/* Force.

	   remove the image even if used by containers or has other tags
	*/
	Force *bool

	/* Name.

	   name or ID of image to remove
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image delete libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageDeleteLibpodParams) WithDefaults() *ImageDeleteLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image delete libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageDeleteLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image delete libpod params
func (o *ImageDeleteLibpodParams) WithTimeout(timeout time.Duration) *ImageDeleteLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image delete libpod params
func (o *ImageDeleteLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image delete libpod params
func (o *ImageDeleteLibpodParams) WithContext(ctx context.Context) *ImageDeleteLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image delete libpod params
func (o *ImageDeleteLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image delete libpod params
func (o *ImageDeleteLibpodParams) WithHTTPClient(client *http.Client) *ImageDeleteLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image delete libpod params
func (o *ImageDeleteLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the image delete libpod params
func (o *ImageDeleteLibpodParams) WithForce(force *bool) *ImageDeleteLibpodParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the image delete libpod params
func (o *ImageDeleteLibpodParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the image delete libpod params
func (o *ImageDeleteLibpodParams) WithName(name string) *ImageDeleteLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the image delete libpod params
func (o *ImageDeleteLibpodParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ImageDeleteLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
