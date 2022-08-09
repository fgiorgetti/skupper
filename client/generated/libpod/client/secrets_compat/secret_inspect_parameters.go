// Code generated by go-swagger; DO NOT EDIT.

package secrets_compat

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

// NewSecretInspectParams creates a new SecretInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretInspectParams() *SecretInspectParams {
	return &SecretInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretInspectParamsWithTimeout creates a new SecretInspectParams object
// with the ability to set a timeout on a request.
func NewSecretInspectParamsWithTimeout(timeout time.Duration) *SecretInspectParams {
	return &SecretInspectParams{
		timeout: timeout,
	}
}

// NewSecretInspectParamsWithContext creates a new SecretInspectParams object
// with the ability to set a context for a request.
func NewSecretInspectParamsWithContext(ctx context.Context) *SecretInspectParams {
	return &SecretInspectParams{
		Context: ctx,
	}
}

// NewSecretInspectParamsWithHTTPClient creates a new SecretInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretInspectParamsWithHTTPClient(client *http.Client) *SecretInspectParams {
	return &SecretInspectParams{
		HTTPClient: client,
	}
}

/* SecretInspectParams contains all the parameters to send to the API endpoint
   for the secret inspect operation.

   Typically these are written to a http.Request.
*/
type SecretInspectParams struct {

	/* Name.

	   the name or ID of the secret
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secret inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretInspectParams) WithDefaults() *SecretInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secret inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secret inspect params
func (o *SecretInspectParams) WithTimeout(timeout time.Duration) *SecretInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secret inspect params
func (o *SecretInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secret inspect params
func (o *SecretInspectParams) WithContext(ctx context.Context) *SecretInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secret inspect params
func (o *SecretInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secret inspect params
func (o *SecretInspectParams) WithHTTPClient(client *http.Client) *SecretInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secret inspect params
func (o *SecretInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the secret inspect params
func (o *SecretInspectParams) WithName(name string) *SecretInspectParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the secret inspect params
func (o *SecretInspectParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SecretInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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