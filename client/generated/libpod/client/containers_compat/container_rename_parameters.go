// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

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

// NewContainerRenameParams creates a new ContainerRenameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerRenameParams() *ContainerRenameParams {
	return &ContainerRenameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerRenameParamsWithTimeout creates a new ContainerRenameParams object
// with the ability to set a timeout on a request.
func NewContainerRenameParamsWithTimeout(timeout time.Duration) *ContainerRenameParams {
	return &ContainerRenameParams{
		timeout: timeout,
	}
}

// NewContainerRenameParamsWithContext creates a new ContainerRenameParams object
// with the ability to set a context for a request.
func NewContainerRenameParamsWithContext(ctx context.Context) *ContainerRenameParams {
	return &ContainerRenameParams{
		Context: ctx,
	}
}

// NewContainerRenameParamsWithHTTPClient creates a new ContainerRenameParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerRenameParamsWithHTTPClient(client *http.Client) *ContainerRenameParams {
	return &ContainerRenameParams{
		HTTPClient: client,
	}
}

/* ContainerRenameParams contains all the parameters to send to the API endpoint
   for the container rename operation.

   Typically these are written to a http.Request.
*/
type ContainerRenameParams struct {

	/* Name.

	   New name for the container
	*/
	QueryName string

	/* Name.

	   Full or partial ID or full name of the container to rename
	*/
	PathName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container rename params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerRenameParams) WithDefaults() *ContainerRenameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container rename params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerRenameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container rename params
func (o *ContainerRenameParams) WithTimeout(timeout time.Duration) *ContainerRenameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container rename params
func (o *ContainerRenameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container rename params
func (o *ContainerRenameParams) WithContext(ctx context.Context) *ContainerRenameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container rename params
func (o *ContainerRenameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container rename params
func (o *ContainerRenameParams) WithHTTPClient(client *http.Client) *ContainerRenameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container rename params
func (o *ContainerRenameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQueryName adds the name to the container rename params
func (o *ContainerRenameParams) WithQueryName(name string) *ContainerRenameParams {
	o.SetQueryName(name)
	return o
}

// SetQueryName adds the name to the container rename params
func (o *ContainerRenameParams) SetQueryName(name string) {
	o.QueryName = name
}

// WithPathName adds the name to the container rename params
func (o *ContainerRenameParams) WithPathName(name string) *ContainerRenameParams {
	o.SetPathName(name)
	return o
}

// SetPathName adds the name to the container rename params
func (o *ContainerRenameParams) SetPathName(name string) {
	o.PathName = name
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerRenameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param name
	qrName := o.QueryName
	qName := qrName
	if qName != "" {

		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.PathName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
