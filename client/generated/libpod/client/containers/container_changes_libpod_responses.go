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

// ContainerChangesLibpodReader is a Reader for the ContainerChangesLibpod structure.
type ContainerChangesLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerChangesLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerChangesLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerChangesLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerChangesLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerChangesLibpodOK creates a ContainerChangesLibpodOK with default headers values
func NewContainerChangesLibpodOK() *ContainerChangesLibpodOK {
	return &ContainerChangesLibpodOK{}
}

/* ContainerChangesLibpodOK describes a response with status code 200, with default header values.

Array of Changes
*/
type ContainerChangesLibpodOK struct {
}

func (o *ContainerChangesLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/changes][%d] containerChangesLibpodOK ", 200)
}

func (o *ContainerChangesLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerChangesLibpodNotFound creates a ContainerChangesLibpodNotFound with default headers values
func NewContainerChangesLibpodNotFound() *ContainerChangesLibpodNotFound {
	return &ContainerChangesLibpodNotFound{}
}

/* ContainerChangesLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerChangesLibpodNotFound struct {
	Payload *ContainerChangesLibpodNotFoundBody
}

func (o *ContainerChangesLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/changes][%d] containerChangesLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ContainerChangesLibpodNotFound) GetPayload() *ContainerChangesLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerChangesLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerChangesLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerChangesLibpodInternalServerError creates a ContainerChangesLibpodInternalServerError with default headers values
func NewContainerChangesLibpodInternalServerError() *ContainerChangesLibpodInternalServerError {
	return &ContainerChangesLibpodInternalServerError{}
}

/* ContainerChangesLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerChangesLibpodInternalServerError struct {
	Payload *ContainerChangesLibpodInternalServerErrorBody
}

func (o *ContainerChangesLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/changes][%d] containerChangesLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerChangesLibpodInternalServerError) GetPayload() *ContainerChangesLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerChangesLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerChangesLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerChangesLibpodInternalServerErrorBody container changes libpod internal server error body
swagger:model ContainerChangesLibpodInternalServerErrorBody
*/
type ContainerChangesLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container changes libpod internal server error body
func (o *ContainerChangesLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container changes libpod internal server error body based on context it is used
func (o *ContainerChangesLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerChangesLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerChangesLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerChangesLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerChangesLibpodNotFoundBody container changes libpod not found body
swagger:model ContainerChangesLibpodNotFoundBody
*/
type ContainerChangesLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container changes libpod not found body
func (o *ContainerChangesLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container changes libpod not found body based on context it is used
func (o *ContainerChangesLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerChangesLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerChangesLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerChangesLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
