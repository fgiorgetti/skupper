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

// ContainerCheckpointLibpodReader is a Reader for the ContainerCheckpointLibpod structure.
type ContainerCheckpointLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerCheckpointLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerCheckpointLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerCheckpointLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerCheckpointLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerCheckpointLibpodOK creates a ContainerCheckpointLibpodOK with default headers values
func NewContainerCheckpointLibpodOK() *ContainerCheckpointLibpodOK {
	return &ContainerCheckpointLibpodOK{}
}

/* ContainerCheckpointLibpodOK describes a response with status code 200, with default header values.

tarball is returned in body if exported
*/
type ContainerCheckpointLibpodOK struct {
}

func (o *ContainerCheckpointLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/checkpoint][%d] containerCheckpointLibpodOK ", 200)
}

func (o *ContainerCheckpointLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerCheckpointLibpodNotFound creates a ContainerCheckpointLibpodNotFound with default headers values
func NewContainerCheckpointLibpodNotFound() *ContainerCheckpointLibpodNotFound {
	return &ContainerCheckpointLibpodNotFound{}
}

/* ContainerCheckpointLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerCheckpointLibpodNotFound struct {
	Payload *ContainerCheckpointLibpodNotFoundBody
}

func (o *ContainerCheckpointLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/checkpoint][%d] containerCheckpointLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ContainerCheckpointLibpodNotFound) GetPayload() *ContainerCheckpointLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerCheckpointLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerCheckpointLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerCheckpointLibpodInternalServerError creates a ContainerCheckpointLibpodInternalServerError with default headers values
func NewContainerCheckpointLibpodInternalServerError() *ContainerCheckpointLibpodInternalServerError {
	return &ContainerCheckpointLibpodInternalServerError{}
}

/* ContainerCheckpointLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerCheckpointLibpodInternalServerError struct {
	Payload *ContainerCheckpointLibpodInternalServerErrorBody
}

func (o *ContainerCheckpointLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/containers/{name}/checkpoint][%d] containerCheckpointLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerCheckpointLibpodInternalServerError) GetPayload() *ContainerCheckpointLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerCheckpointLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerCheckpointLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerCheckpointLibpodInternalServerErrorBody container checkpoint libpod internal server error body
swagger:model ContainerCheckpointLibpodInternalServerErrorBody
*/
type ContainerCheckpointLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container checkpoint libpod internal server error body
func (o *ContainerCheckpointLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container checkpoint libpod internal server error body based on context it is used
func (o *ContainerCheckpointLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerCheckpointLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerCheckpointLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerCheckpointLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerCheckpointLibpodNotFoundBody container checkpoint libpod not found body
swagger:model ContainerCheckpointLibpodNotFoundBody
*/
type ContainerCheckpointLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container checkpoint libpod not found body
func (o *ContainerCheckpointLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container checkpoint libpod not found body based on context it is used
func (o *ContainerCheckpointLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerCheckpointLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerCheckpointLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerCheckpointLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}