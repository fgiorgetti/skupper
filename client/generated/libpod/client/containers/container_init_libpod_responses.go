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

// ContainerInitLibpodReader is a Reader for the ContainerInitLibpod structure.
type ContainerInitLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerInitLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerInitLibpodNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewContainerInitLibpodNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerInitLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerInitLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerInitLibpodNoContent creates a ContainerInitLibpodNoContent with default headers values
func NewContainerInitLibpodNoContent() *ContainerInitLibpodNoContent {
	return &ContainerInitLibpodNoContent{}
}

/* ContainerInitLibpodNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerInitLibpodNoContent struct {
}

func (o *ContainerInitLibpodNoContent) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/init][%d] containerInitLibpodNoContent ", 204)
}

func (o *ContainerInitLibpodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerInitLibpodNotModified creates a ContainerInitLibpodNotModified with default headers values
func NewContainerInitLibpodNotModified() *ContainerInitLibpodNotModified {
	return &ContainerInitLibpodNotModified{}
}

/* ContainerInitLibpodNotModified describes a response with status code 304, with default header values.

container already initialized
*/
type ContainerInitLibpodNotModified struct {
}

func (o *ContainerInitLibpodNotModified) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/init][%d] containerInitLibpodNotModified ", 304)
}

func (o *ContainerInitLibpodNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerInitLibpodNotFound creates a ContainerInitLibpodNotFound with default headers values
func NewContainerInitLibpodNotFound() *ContainerInitLibpodNotFound {
	return &ContainerInitLibpodNotFound{}
}

/* ContainerInitLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerInitLibpodNotFound struct {
	Payload *ContainerInitLibpodNotFoundBody
}

func (o *ContainerInitLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/init][%d] containerInitLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ContainerInitLibpodNotFound) GetPayload() *ContainerInitLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerInitLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerInitLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerInitLibpodInternalServerError creates a ContainerInitLibpodInternalServerError with default headers values
func NewContainerInitLibpodInternalServerError() *ContainerInitLibpodInternalServerError {
	return &ContainerInitLibpodInternalServerError{}
}

/* ContainerInitLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerInitLibpodInternalServerError struct {
	Payload *ContainerInitLibpodInternalServerErrorBody
}

func (o *ContainerInitLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/init][%d] containerInitLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerInitLibpodInternalServerError) GetPayload() *ContainerInitLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerInitLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerInitLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerInitLibpodInternalServerErrorBody container init libpod internal server error body
swagger:model ContainerInitLibpodInternalServerErrorBody
*/
type ContainerInitLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container init libpod internal server error body
func (o *ContainerInitLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container init libpod internal server error body based on context it is used
func (o *ContainerInitLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerInitLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerInitLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerInitLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerInitLibpodNotFoundBody container init libpod not found body
swagger:model ContainerInitLibpodNotFoundBody
*/
type ContainerInitLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container init libpod not found body
func (o *ContainerInitLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container init libpod not found body based on context it is used
func (o *ContainerInitLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerInitLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerInitLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerInitLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}