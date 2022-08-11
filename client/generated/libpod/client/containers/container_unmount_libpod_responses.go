// Code generated by go-swagger; DO NOT EDIT.

package containers

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

// ContainerUnmountLibpodReader is a Reader for the ContainerUnmountLibpod structure.
type ContainerUnmountLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerUnmountLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerUnmountLibpodNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerUnmountLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerUnmountLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerUnmountLibpodNoContent creates a ContainerUnmountLibpodNoContent with default headers values
func NewContainerUnmountLibpodNoContent() *ContainerUnmountLibpodNoContent {
	return &ContainerUnmountLibpodNoContent{}
}

/* ContainerUnmountLibpodNoContent describes a response with status code 204, with default header values.

ok
*/
type ContainerUnmountLibpodNoContent struct {
}

func (o *ContainerUnmountLibpodNoContent) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/unmount][%d] containerUnmountLibpodNoContent ", 204)
}

func (o *ContainerUnmountLibpodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerUnmountLibpodNotFound creates a ContainerUnmountLibpodNotFound with default headers values
func NewContainerUnmountLibpodNotFound() *ContainerUnmountLibpodNotFound {
	return &ContainerUnmountLibpodNotFound{}
}

/* ContainerUnmountLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerUnmountLibpodNotFound struct {
	Payload *ContainerUnmountLibpodNotFoundBody
}

func (o *ContainerUnmountLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/unmount][%d] containerUnmountLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ContainerUnmountLibpodNotFound) GetPayload() *ContainerUnmountLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerUnmountLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerUnmountLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerUnmountLibpodInternalServerError creates a ContainerUnmountLibpodInternalServerError with default headers values
func NewContainerUnmountLibpodInternalServerError() *ContainerUnmountLibpodInternalServerError {
	return &ContainerUnmountLibpodInternalServerError{}
}

/* ContainerUnmountLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerUnmountLibpodInternalServerError struct {
	Payload *ContainerUnmountLibpodInternalServerErrorBody
}

func (o *ContainerUnmountLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/unmount][%d] containerUnmountLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerUnmountLibpodInternalServerError) GetPayload() *ContainerUnmountLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerUnmountLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerUnmountLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerUnmountLibpodInternalServerErrorBody container unmount libpod internal server error body
swagger:model ContainerUnmountLibpodInternalServerErrorBody
*/
type ContainerUnmountLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container unmount libpod internal server error body
func (o *ContainerUnmountLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container unmount libpod internal server error body based on context it is used
func (o *ContainerUnmountLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerUnmountLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerUnmountLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerUnmountLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerUnmountLibpodNotFoundBody container unmount libpod not found body
swagger:model ContainerUnmountLibpodNotFoundBody
*/
type ContainerUnmountLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container unmount libpod not found body
func (o *ContainerUnmountLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container unmount libpod not found body based on context it is used
func (o *ContainerUnmountLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerUnmountLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerUnmountLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerUnmountLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}