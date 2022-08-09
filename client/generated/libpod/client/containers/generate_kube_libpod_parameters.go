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

// NewGenerateKubeLibpodParams creates a new GenerateKubeLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateKubeLibpodParams() *GenerateKubeLibpodParams {
	return &GenerateKubeLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateKubeLibpodParamsWithTimeout creates a new GenerateKubeLibpodParams object
// with the ability to set a timeout on a request.
func NewGenerateKubeLibpodParamsWithTimeout(timeout time.Duration) *GenerateKubeLibpodParams {
	return &GenerateKubeLibpodParams{
		timeout: timeout,
	}
}

// NewGenerateKubeLibpodParamsWithContext creates a new GenerateKubeLibpodParams object
// with the ability to set a context for a request.
func NewGenerateKubeLibpodParamsWithContext(ctx context.Context) *GenerateKubeLibpodParams {
	return &GenerateKubeLibpodParams{
		Context: ctx,
	}
}

// NewGenerateKubeLibpodParamsWithHTTPClient creates a new GenerateKubeLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateKubeLibpodParamsWithHTTPClient(client *http.Client) *GenerateKubeLibpodParams {
	return &GenerateKubeLibpodParams{
		HTTPClient: client,
	}
}

/* GenerateKubeLibpodParams contains all the parameters to send to the API endpoint
   for the generate kube libpod operation.

   Typically these are written to a http.Request.
*/
type GenerateKubeLibpodParams struct {

	/* Names.

	   Name or ID of the container or pod.
	*/
	Names []string

	/* Service.

	   Generate YAML for a Kubernetes service object.
	*/
	Service *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate kube libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateKubeLibpodParams) WithDefaults() *GenerateKubeLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate kube libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateKubeLibpodParams) SetDefaults() {
	var (
		serviceDefault = bool(false)
	)

	val := GenerateKubeLibpodParams{
		Service: &serviceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the generate kube libpod params
func (o *GenerateKubeLibpodParams) WithTimeout(timeout time.Duration) *GenerateKubeLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate kube libpod params
func (o *GenerateKubeLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate kube libpod params
func (o *GenerateKubeLibpodParams) WithContext(ctx context.Context) *GenerateKubeLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate kube libpod params
func (o *GenerateKubeLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate kube libpod params
func (o *GenerateKubeLibpodParams) WithHTTPClient(client *http.Client) *GenerateKubeLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate kube libpod params
func (o *GenerateKubeLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the generate kube libpod params
func (o *GenerateKubeLibpodParams) WithNames(names []string) *GenerateKubeLibpodParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the generate kube libpod params
func (o *GenerateKubeLibpodParams) SetNames(names []string) {
	o.Names = names
}

// WithService adds the service to the generate kube libpod params
func (o *GenerateKubeLibpodParams) WithService(service *bool) *GenerateKubeLibpodParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the generate kube libpod params
func (o *GenerateKubeLibpodParams) SetService(service *bool) {
	o.Service = service
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateKubeLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Names != nil {

		// binding items for names
		joinedNames := o.bindParamNames(reg)

		// query array param names
		if err := r.SetQueryParam("names", joinedNames...); err != nil {
			return err
		}
	}

	if o.Service != nil {

		// query param service
		var qrService bool

		if o.Service != nil {
			qrService = *o.Service
		}
		qService := swag.FormatBool(qrService)
		if qService != "" {

			if err := r.SetQueryParam("service", qService); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGenerateKubeLibpod binds the parameter names
func (o *GenerateKubeLibpodParams) bindParamNames(formats strfmt.Registry) []string {
	namesIR := o.Names

	var namesIC []string
	for _, namesIIR := range namesIR { // explode []string

		namesIIV := namesIIR // string as string
		namesIC = append(namesIC, namesIIV)
	}

	// items.CollectionFormat: ""
	namesIS := swag.JoinByFormat(namesIC, "")

	return namesIS
}