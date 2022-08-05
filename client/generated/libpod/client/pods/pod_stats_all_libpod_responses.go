// Code generated by go-swagger; DO NOT EDIT.

package pods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PodStatsAllLibpodReader is a Reader for the PodStatsAllLibpod structure.
type PodStatsAllLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PodStatsAllLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPodStatsAllLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPodStatsAllLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPodStatsAllLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPodStatsAllLibpodOK creates a PodStatsAllLibpodOK with default headers values
func NewPodStatsAllLibpodOK() *PodStatsAllLibpodOK {
	return &PodStatsAllLibpodOK{}
}

/* PodStatsAllLibpodOK describes a response with status code 200, with default header values.

List processes in pod
*/
type PodStatsAllLibpodOK struct {
	Payload *PodStatsAllLibpodOKBody
}

func (o *PodStatsAllLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/pods/stats][%d] podStatsAllLibpodOK  %+v", 200, o.Payload)
}
func (o *PodStatsAllLibpodOK) GetPayload() *PodStatsAllLibpodOKBody {
	return o.Payload
}

func (o *PodStatsAllLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodStatsAllLibpodOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodStatsAllLibpodNotFound creates a PodStatsAllLibpodNotFound with default headers values
func NewPodStatsAllLibpodNotFound() *PodStatsAllLibpodNotFound {
	return &PodStatsAllLibpodNotFound{}
}

/* PodStatsAllLibpodNotFound describes a response with status code 404, with default header values.

No such pod
*/
type PodStatsAllLibpodNotFound struct {
	Payload *PodStatsAllLibpodNotFoundBody
}

func (o *PodStatsAllLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/pods/stats][%d] podStatsAllLibpodNotFound  %+v", 404, o.Payload)
}
func (o *PodStatsAllLibpodNotFound) GetPayload() *PodStatsAllLibpodNotFoundBody {
	return o.Payload
}

func (o *PodStatsAllLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodStatsAllLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodStatsAllLibpodInternalServerError creates a PodStatsAllLibpodInternalServerError with default headers values
func NewPodStatsAllLibpodInternalServerError() *PodStatsAllLibpodInternalServerError {
	return &PodStatsAllLibpodInternalServerError{}
}

/* PodStatsAllLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PodStatsAllLibpodInternalServerError struct {
	Payload *PodStatsAllLibpodInternalServerErrorBody
}

func (o *PodStatsAllLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/pods/stats][%d] podStatsAllLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *PodStatsAllLibpodInternalServerError) GetPayload() *PodStatsAllLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *PodStatsAllLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodStatsAllLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PodStatsAllLibpodInternalServerErrorBody pod stats all libpod internal server error body
swagger:model PodStatsAllLibpodInternalServerErrorBody
*/
type PodStatsAllLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod stats all libpod internal server error body
func (o *PodStatsAllLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod stats all libpod internal server error body based on context it is used
func (o *PodStatsAllLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodStatsAllLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodStatsAllLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PodStatsAllLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PodStatsAllLibpodNotFoundBody pod stats all libpod not found body
swagger:model PodStatsAllLibpodNotFoundBody
*/
type PodStatsAllLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod stats all libpod not found body
func (o *PodStatsAllLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod stats all libpod not found body based on context it is used
func (o *PodStatsAllLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodStatsAllLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodStatsAllLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PodStatsAllLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PodStatsAllLibpodOKBody pod stats all libpod o k body
swagger:model PodStatsAllLibpodOKBody
*/
type PodStatsAllLibpodOKBody struct {

	// Each process running in the container, where each is process
	// is an array of values corresponding to the titles.
	// Required: true
	Processes [][]string `json:"Processes"`

	// The ps column titles
	// Required: true
	Titles []string `json:"Titles"`
}

// Validate validates this pod stats all libpod o k body
func (o *PodStatsAllLibpodOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProcesses(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PodStatsAllLibpodOKBody) validateProcesses(formats strfmt.Registry) error {

	if err := validate.Required("podStatsAllLibpodOK"+"."+"Processes", "body", o.Processes); err != nil {
		return err
	}

	return nil
}

func (o *PodStatsAllLibpodOKBody) validateTitles(formats strfmt.Registry) error {

	if err := validate.Required("podStatsAllLibpodOK"+"."+"Titles", "body", o.Titles); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pod stats all libpod o k body based on context it is used
func (o *PodStatsAllLibpodOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodStatsAllLibpodOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodStatsAllLibpodOKBody) UnmarshalBinary(b []byte) error {
	var res PodStatsAllLibpodOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
