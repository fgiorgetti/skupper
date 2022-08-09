// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

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

// ContainerTopReader is a Reader for the ContainerTop structure.
type ContainerTopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerTopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerTopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerTopNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerTopInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerTopOK creates a ContainerTopOK with default headers values
func NewContainerTopOK() *ContainerTopOK {
	return &ContainerTopOK{}
}

/* ContainerTopOK describes a response with status code 200, with default header values.

List processes in container
*/
type ContainerTopOK struct {
	Payload *ContainerTopOKBody
}

func (o *ContainerTopOK) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/top][%d] containerTopOK  %+v", 200, o.Payload)
}
func (o *ContainerTopOK) GetPayload() *ContainerTopOKBody {
	return o.Payload
}

func (o *ContainerTopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerTopOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerTopNotFound creates a ContainerTopNotFound with default headers values
func NewContainerTopNotFound() *ContainerTopNotFound {
	return &ContainerTopNotFound{}
}

/* ContainerTopNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerTopNotFound struct {
	Payload *ContainerTopNotFoundBody
}

func (o *ContainerTopNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/top][%d] containerTopNotFound  %+v", 404, o.Payload)
}
func (o *ContainerTopNotFound) GetPayload() *ContainerTopNotFoundBody {
	return o.Payload
}

func (o *ContainerTopNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerTopNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerTopInternalServerError creates a ContainerTopInternalServerError with default headers values
func NewContainerTopInternalServerError() *ContainerTopInternalServerError {
	return &ContainerTopInternalServerError{}
}

/* ContainerTopInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerTopInternalServerError struct {
	Payload *ContainerTopInternalServerErrorBody
}

func (o *ContainerTopInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/top][%d] containerTopInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerTopInternalServerError) GetPayload() *ContainerTopInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerTopInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerTopInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerTopInternalServerErrorBody container top internal server error body
swagger:model ContainerTopInternalServerErrorBody
*/
type ContainerTopInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container top internal server error body
func (o *ContainerTopInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container top internal server error body based on context it is used
func (o *ContainerTopInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerTopInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerTopInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerTopInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerTopNotFoundBody container top not found body
swagger:model ContainerTopNotFoundBody
*/
type ContainerTopNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container top not found body
func (o *ContainerTopNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container top not found body based on context it is used
func (o *ContainerTopNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerTopNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerTopNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerTopNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerTopOKBody container top o k body
swagger:model ContainerTopOKBody
*/
type ContainerTopOKBody struct {

	// Each process running in the container, where each is process
	// is an array of values corresponding to the titles.
	// Required: true
	Processes [][]string `json:"Processes"`

	// The ps column titles
	// Required: true
	Titles []string `json:"Titles"`
}

// Validate validates this container top o k body
func (o *ContainerTopOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ContainerTopOKBody) validateProcesses(formats strfmt.Registry) error {

	if err := validate.Required("containerTopOK"+"."+"Processes", "body", o.Processes); err != nil {
		return err
	}

	return nil
}

func (o *ContainerTopOKBody) validateTitles(formats strfmt.Registry) error {

	if err := validate.Required("containerTopOK"+"."+"Titles", "body", o.Titles); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this container top o k body based on context it is used
func (o *ContainerTopOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerTopOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerTopOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerTopOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}