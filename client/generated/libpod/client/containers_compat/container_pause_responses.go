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

// ContainerPauseReader is a Reader for the ContainerPause structure.
type ContainerPauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerPauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerPauseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerPauseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerPauseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerPauseNoContent creates a ContainerPauseNoContent with default headers values
func NewContainerPauseNoContent() *ContainerPauseNoContent {
	return &ContainerPauseNoContent{}
}

/* ContainerPauseNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerPauseNoContent struct {
}

func (o *ContainerPauseNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/pause][%d] containerPauseNoContent ", 204)
}

func (o *ContainerPauseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerPauseNotFound creates a ContainerPauseNotFound with default headers values
func NewContainerPauseNotFound() *ContainerPauseNotFound {
	return &ContainerPauseNotFound{}
}

/* ContainerPauseNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerPauseNotFound struct {
	Payload *ContainerPauseNotFoundBody
}

func (o *ContainerPauseNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/pause][%d] containerPauseNotFound  %+v", 404, o.Payload)
}
func (o *ContainerPauseNotFound) GetPayload() *ContainerPauseNotFoundBody {
	return o.Payload
}

func (o *ContainerPauseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerPauseNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerPauseInternalServerError creates a ContainerPauseInternalServerError with default headers values
func NewContainerPauseInternalServerError() *ContainerPauseInternalServerError {
	return &ContainerPauseInternalServerError{}
}

/* ContainerPauseInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerPauseInternalServerError struct {
	Payload *ContainerPauseInternalServerErrorBody
}

func (o *ContainerPauseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{name}/pause][%d] containerPauseInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerPauseInternalServerError) GetPayload() *ContainerPauseInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerPauseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerPauseInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerPauseInternalServerErrorBody container pause internal server error body
swagger:model ContainerPauseInternalServerErrorBody
*/
type ContainerPauseInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container pause internal server error body
func (o *ContainerPauseInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container pause internal server error body based on context it is used
func (o *ContainerPauseInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerPauseInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerPauseInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerPauseInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerPauseNotFoundBody container pause not found body
swagger:model ContainerPauseNotFoundBody
*/
type ContainerPauseNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container pause not found body
func (o *ContainerPauseNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container pause not found body based on context it is used
func (o *ContainerPauseNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerPauseNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerPauseNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerPauseNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
