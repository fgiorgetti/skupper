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

// NewImagePushParams creates a new ImagePushParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImagePushParams() *ImagePushParams {
	return &ImagePushParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImagePushParamsWithTimeout creates a new ImagePushParams object
// with the ability to set a timeout on a request.
func NewImagePushParamsWithTimeout(timeout time.Duration) *ImagePushParams {
	return &ImagePushParams{
		timeout: timeout,
	}
}

// NewImagePushParamsWithContext creates a new ImagePushParams object
// with the ability to set a context for a request.
func NewImagePushParamsWithContext(ctx context.Context) *ImagePushParams {
	return &ImagePushParams{
		Context: ctx,
	}
}

// NewImagePushParamsWithHTTPClient creates a new ImagePushParams object
// with the ability to set a custom HTTPClient for a request.
func NewImagePushParamsWithHTTPClient(client *http.Client) *ImagePushParams {
	return &ImagePushParams{
		HTTPClient: client,
	}
}

/* ImagePushParams contains all the parameters to send to the API endpoint
   for the image push operation.

   Typically these are written to a http.Request.
*/
type ImagePushParams struct {

	/* XRegistryAuth.

	   A base64-encoded auth configuration.
	*/
	XRegistryAuth *string

	/* All.

	   All indicates whether to push all images related to the image list
	*/
	All *bool

	/* Compress.

	   use compression on image
	*/
	Compress *bool

	/* Destination.

	   destination name for the image being pushed
	*/
	Destination *string

	/* Name.

	   Name of image to push.
	*/
	Name string

	/* Tag.

	   The tag to associate with the image on the registry.
	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image push params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImagePushParams) WithDefaults() *ImagePushParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image push params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImagePushParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image push params
func (o *ImagePushParams) WithTimeout(timeout time.Duration) *ImagePushParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image push params
func (o *ImagePushParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image push params
func (o *ImagePushParams) WithContext(ctx context.Context) *ImagePushParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image push params
func (o *ImagePushParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image push params
func (o *ImagePushParams) WithHTTPClient(client *http.Client) *ImagePushParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image push params
func (o *ImagePushParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRegistryAuth adds the xRegistryAuth to the image push params
func (o *ImagePushParams) WithXRegistryAuth(xRegistryAuth *string) *ImagePushParams {
	o.SetXRegistryAuth(xRegistryAuth)
	return o
}

// SetXRegistryAuth adds the xRegistryAuth to the image push params
func (o *ImagePushParams) SetXRegistryAuth(xRegistryAuth *string) {
	o.XRegistryAuth = xRegistryAuth
}

// WithAll adds the all to the image push params
func (o *ImagePushParams) WithAll(all *bool) *ImagePushParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the image push params
func (o *ImagePushParams) SetAll(all *bool) {
	o.All = all
}

// WithCompress adds the compress to the image push params
func (o *ImagePushParams) WithCompress(compress *bool) *ImagePushParams {
	o.SetCompress(compress)
	return o
}

// SetCompress adds the compress to the image push params
func (o *ImagePushParams) SetCompress(compress *bool) {
	o.Compress = compress
}

// WithDestination adds the destination to the image push params
func (o *ImagePushParams) WithDestination(destination *string) *ImagePushParams {
	o.SetDestination(destination)
	return o
}

// SetDestination adds the destination to the image push params
func (o *ImagePushParams) SetDestination(destination *string) {
	o.Destination = destination
}

// WithName adds the name to the image push params
func (o *ImagePushParams) WithName(name string) *ImagePushParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the image push params
func (o *ImagePushParams) SetName(name string) {
	o.Name = name
}

// WithTag adds the tag to the image push params
func (o *ImagePushParams) WithTag(tag *string) *ImagePushParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the image push params
func (o *ImagePushParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *ImagePushParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRegistryAuth != nil {

		// header param X-Registry-Auth
		if err := r.SetHeaderParam("X-Registry-Auth", *o.XRegistryAuth); err != nil {
			return err
		}
	}

	if o.All != nil {

		// query param all
		var qrAll bool

		if o.All != nil {
			qrAll = *o.All
		}
		qAll := swag.FormatBool(qrAll)
		if qAll != "" {

			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}
	}

	if o.Compress != nil {

		// query param compress
		var qrCompress bool

		if o.Compress != nil {
			qrCompress = *o.Compress
		}
		qCompress := swag.FormatBool(qrCompress)
		if qCompress != "" {

			if err := r.SetQueryParam("compress", qCompress); err != nil {
				return err
			}
		}
	}

	if o.Destination != nil {

		// query param destination
		var qrDestination string

		if o.Destination != nil {
			qrDestination = *o.Destination
		}
		qDestination := qrDestination
		if qDestination != "" {

			if err := r.SetQueryParam("destination", qDestination); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
