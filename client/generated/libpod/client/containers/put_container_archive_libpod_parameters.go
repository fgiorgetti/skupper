// Code generated by go-swagger; DO NOT EDIT.

package containers

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

// NewPutContainerArchiveLibpodParams creates a new PutContainerArchiveLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutContainerArchiveLibpodParams() *PutContainerArchiveLibpodParams {
	return &PutContainerArchiveLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutContainerArchiveLibpodParamsWithTimeout creates a new PutContainerArchiveLibpodParams object
// with the ability to set a timeout on a request.
func NewPutContainerArchiveLibpodParamsWithTimeout(timeout time.Duration) *PutContainerArchiveLibpodParams {
	return &PutContainerArchiveLibpodParams{
		timeout: timeout,
	}
}

// NewPutContainerArchiveLibpodParamsWithContext creates a new PutContainerArchiveLibpodParams object
// with the ability to set a context for a request.
func NewPutContainerArchiveLibpodParamsWithContext(ctx context.Context) *PutContainerArchiveLibpodParams {
	return &PutContainerArchiveLibpodParams{
		Context: ctx,
	}
}

// NewPutContainerArchiveLibpodParamsWithHTTPClient creates a new PutContainerArchiveLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutContainerArchiveLibpodParamsWithHTTPClient(client *http.Client) *PutContainerArchiveLibpodParams {
	return &PutContainerArchiveLibpodParams{
		HTTPClient: client,
	}
}

/* PutContainerArchiveLibpodParams contains all the parameters to send to the API endpoint
   for the put container archive libpod operation.

   Typically these are written to a http.Request.
*/
type PutContainerArchiveLibpodParams struct {

	/* Name.

	   container name or id
	*/
	Name string

	/* Path.

	   Path to a directory in the container to extract
	*/
	Path string

	/* Pause.

	   pause the container while copying (defaults to true)

	   Default: true
	*/
	Pause *bool

	/* Request.

	   tarfile of files to copy into the container
	*/
	Request string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put container archive libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutContainerArchiveLibpodParams) WithDefaults() *PutContainerArchiveLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put container archive libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutContainerArchiveLibpodParams) SetDefaults() {
	var (
		pauseDefault = bool(true)
	)

	val := PutContainerArchiveLibpodParams{
		Pause: &pauseDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) WithTimeout(timeout time.Duration) *PutContainerArchiveLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) WithContext(ctx context.Context) *PutContainerArchiveLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) WithHTTPClient(client *http.Client) *PutContainerArchiveLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) WithName(name string) *PutContainerArchiveLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) SetName(name string) {
	o.Name = name
}

// WithPath adds the path to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) WithPath(path string) *PutContainerArchiveLibpodParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) SetPath(path string) {
	o.Path = path
}

// WithPause adds the pause to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) WithPause(pause *bool) *PutContainerArchiveLibpodParams {
	o.SetPause(pause)
	return o
}

// SetPause adds the pause to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) SetPause(pause *bool) {
	o.Pause = pause
}

// WithRequest adds the request to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) WithRequest(request string) *PutContainerArchiveLibpodParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put container archive libpod params
func (o *PutContainerArchiveLibpodParams) SetRequest(request string) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PutContainerArchiveLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {

		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	if o.Pause != nil {

		// query param pause
		var qrPause bool

		if o.Pause != nil {
			qrPause = *o.Pause
		}
		qPause := swag.FormatBool(qrPause)
		if qPause != "" {

			if err := r.SetQueryParam("pause", qPause); err != nil {
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
