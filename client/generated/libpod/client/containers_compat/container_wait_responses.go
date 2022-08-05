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

	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

// ContainerWaitReader is a Reader for the ContainerWait structure.
type ContainerWaitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerWaitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerWaitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerWaitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerWaitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerWaitOK creates a ContainerWaitOK with default headers values
func NewContainerWaitOK() *ContainerWaitOK {
	return &ContainerWaitOK{}
}

/* ContainerWaitOK describes a response with status code 200, with default header values.

Wait container
*/
type ContainerWaitOK struct {
	Payload *ContainerWaitOKBody
}

func (o *ContainerWaitOK) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/wait][%d] containerWaitOK  %+v", 200, o.Payload)
}
func (o *ContainerWaitOK) GetPayload() *ContainerWaitOKBody {
	return o.Payload
}

func (o *ContainerWaitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerWaitOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerWaitNotFound creates a ContainerWaitNotFound with default headers values
func NewContainerWaitNotFound() *ContainerWaitNotFound {
	return &ContainerWaitNotFound{}
}

/* ContainerWaitNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerWaitNotFound struct {
	Payload *ContainerWaitNotFoundBody
}

func (o *ContainerWaitNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/wait][%d] containerWaitNotFound  %+v", 404, o.Payload)
}
func (o *ContainerWaitNotFound) GetPayload() *ContainerWaitNotFoundBody {
	return o.Payload
}

func (o *ContainerWaitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerWaitNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerWaitInternalServerError creates a ContainerWaitInternalServerError with default headers values
func NewContainerWaitInternalServerError() *ContainerWaitInternalServerError {
	return &ContainerWaitInternalServerError{}
}

/* ContainerWaitInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerWaitInternalServerError struct {
	Payload *ContainerWaitInternalServerErrorBody
}

func (o *ContainerWaitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/wait][%d] containerWaitInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerWaitInternalServerError) GetPayload() *ContainerWaitInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerWaitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerWaitInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerWaitInternalServerErrorBody container wait internal server error body
swagger:model ContainerWaitInternalServerErrorBody
*/
type ContainerWaitInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container wait internal server error body
func (o *ContainerWaitInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container wait internal server error body based on context it is used
func (o *ContainerWaitInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerWaitInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerWaitInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerWaitInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerWaitNotFoundBody container wait not found body
swagger:model ContainerWaitNotFoundBody
*/
type ContainerWaitNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container wait not found body
func (o *ContainerWaitNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container wait not found body based on context it is used
func (o *ContainerWaitNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerWaitNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerWaitNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerWaitNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerWaitOKBody container wait o k body
swagger:model ContainerWaitOKBody
*/
type ContainerWaitOKBody struct {

	// error
	Error *models.ContainerWaitOKBodyError `json:"Error,omitempty"`

	// container exit code
	StatusCode int64 `json:"StatusCode,omitempty"`
}

// Validate validates this container wait o k body
func (o *ContainerWaitOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerWaitOKBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerWaitOK" + "." + "Error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerWaitOK" + "." + "Error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this container wait o k body based on the context it is used
func (o *ContainerWaitOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerWaitOKBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerWaitOK" + "." + "Error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerWaitOK" + "." + "Error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ContainerWaitOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerWaitOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerWaitOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerWaitOKBodyError container wait o k body error
swagger:model ContainerWaitOKBodyError
*/
type ContainerWaitOKBodyError struct {

	// message
	Message string `json:"Message,omitempty"`
}

// Validate validates this container wait o k body error
func (o *ContainerWaitOKBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container wait o k body error based on context it is used
func (o *ContainerWaitOKBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerWaitOKBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerWaitOKBodyError) UnmarshalBinary(b []byte) error {
	var res ContainerWaitOKBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
