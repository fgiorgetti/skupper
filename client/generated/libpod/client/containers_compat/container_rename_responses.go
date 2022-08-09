// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerRenameReader is a Reader for the ContainerRename structure.
type ContainerRenameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRenameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerRenameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerRenameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewContainerRenameConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerRenameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerRenameNoContent creates a ContainerRenameNoContent with default headers values
func NewContainerRenameNoContent() *ContainerRenameNoContent {
	return &ContainerRenameNoContent{}
}

/* ContainerRenameNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerRenameNoContent struct {
}

func (o *ContainerRenameNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/rename][%d] containerRenameNoContent ", 204)
}

func (o *ContainerRenameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerRenameNotFound creates a ContainerRenameNotFound with default headers values
func NewContainerRenameNotFound() *ContainerRenameNotFound {
	return &ContainerRenameNotFound{}
}

/* ContainerRenameNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerRenameNotFound struct {
	Payload *ContainerRenameNotFoundBody
}

func (o *ContainerRenameNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/rename][%d] containerRenameNotFound  %+v", 404, o.Payload)
}
func (o *ContainerRenameNotFound) GetPayload() *ContainerRenameNotFoundBody {
	return o.Payload
}

func (o *ContainerRenameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerRenameNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRenameConflict creates a ContainerRenameConflict with default headers values
func NewContainerRenameConflict() *ContainerRenameConflict {
	return &ContainerRenameConflict{}
}

/* ContainerRenameConflict describes a response with status code 409, with default header values.

Conflict error in operation
*/
type ContainerRenameConflict struct {
	Payload *ContainerRenameConflictBody
}

func (o *ContainerRenameConflict) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/rename][%d] containerRenameConflict  %+v", 409, o.Payload)
}
func (o *ContainerRenameConflict) GetPayload() *ContainerRenameConflictBody {
	return o.Payload
}

func (o *ContainerRenameConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerRenameConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRenameInternalServerError creates a ContainerRenameInternalServerError with default headers values
func NewContainerRenameInternalServerError() *ContainerRenameInternalServerError {
	return &ContainerRenameInternalServerError{}
}

/* ContainerRenameInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerRenameInternalServerError struct {
	Payload *ContainerRenameInternalServerErrorBody
}

func (o *ContainerRenameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/rename][%d] containerRenameInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerRenameInternalServerError) GetPayload() *ContainerRenameInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerRenameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerRenameInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerRenameConflictBody container rename conflict body
swagger:model ContainerRenameConflictBody
*/
type ContainerRenameConflictBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container rename conflict body
func (o *ContainerRenameConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container rename conflict body based on context it is used
func (o *ContainerRenameConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerRenameConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerRenameConflictBody) UnmarshalBinary(b []byte) error {
	var res ContainerRenameConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerRenameInternalServerErrorBody container rename internal server error body
swagger:model ContainerRenameInternalServerErrorBody
*/
type ContainerRenameInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container rename internal server error body
func (o *ContainerRenameInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container rename internal server error body based on context it is used
func (o *ContainerRenameInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerRenameInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerRenameInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerRenameInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerRenameNotFoundBody container rename not found body
swagger:model ContainerRenameNotFoundBody
*/
type ContainerRenameNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container rename not found body
func (o *ContainerRenameNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container rename not found body based on context it is used
func (o *ContainerRenameNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerRenameNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerRenameNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerRenameNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}