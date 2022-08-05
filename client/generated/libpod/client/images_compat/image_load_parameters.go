// Code generated by go-swagger; DO NOT EDIT.

package images_compat

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

// NewImageLoadParams creates a new ImageLoadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageLoadParams() *ImageLoadParams {
	return &ImageLoadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageLoadParamsWithTimeout creates a new ImageLoadParams object
// with the ability to set a timeout on a request.
func NewImageLoadParamsWithTimeout(timeout time.Duration) *ImageLoadParams {
	return &ImageLoadParams{
		timeout: timeout,
	}
}

// NewImageLoadParamsWithContext creates a new ImageLoadParams object
// with the ability to set a context for a request.
func NewImageLoadParamsWithContext(ctx context.Context) *ImageLoadParams {
	return &ImageLoadParams{
		Context: ctx,
	}
}

// NewImageLoadParamsWithHTTPClient creates a new ImageLoadParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageLoadParamsWithHTTPClient(client *http.Client) *ImageLoadParams {
	return &ImageLoadParams{
		HTTPClient: client,
	}
}

/* ImageLoadParams contains all the parameters to send to the API endpoint
   for the image load operation.

   Typically these are written to a http.Request.
*/
type ImageLoadParams struct {

	/* Quiet.

	   not supported
	*/
	Quiet *bool

	/* Request.

	   tarball of container image
	*/
	Request string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image load params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageLoadParams) WithDefaults() *ImageLoadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image load params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageLoadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image load params
func (o *ImageLoadParams) WithTimeout(timeout time.Duration) *ImageLoadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image load params
func (o *ImageLoadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image load params
func (o *ImageLoadParams) WithContext(ctx context.Context) *ImageLoadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image load params
func (o *ImageLoadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image load params
func (o *ImageLoadParams) WithHTTPClient(client *http.Client) *ImageLoadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image load params
func (o *ImageLoadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuiet adds the quiet to the image load params
func (o *ImageLoadParams) WithQuiet(quiet *bool) *ImageLoadParams {
	o.SetQuiet(quiet)
	return o
}

// SetQuiet adds the quiet to the image load params
func (o *ImageLoadParams) SetQuiet(quiet *bool) {
	o.Quiet = quiet
}

// WithRequest adds the request to the image load params
func (o *ImageLoadParams) WithRequest(request string) *ImageLoadParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the image load params
func (o *ImageLoadParams) SetRequest(request string) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *ImageLoadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Quiet != nil {

		// query param quiet
		var qrQuiet bool

		if o.Quiet != nil {
			qrQuiet = *o.Quiet
		}
		qQuiet := swag.FormatBool(qrQuiet)
		if qQuiet != "" {

			if err := r.SetQueryParam("quiet", qQuiet); err != nil {
				return err
			}
		}
	}
	if err := r.SetBodyParam(o.Request); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
