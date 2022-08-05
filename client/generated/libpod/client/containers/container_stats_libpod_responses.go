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

// ContainerStatsLibpodReader is a Reader for the ContainerStatsLibpod structure.
type ContainerStatsLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerStatsLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerStatsLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerStatsLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewContainerStatsLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerStatsLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerStatsLibpodOK creates a ContainerStatsLibpodOK with default headers values
func NewContainerStatsLibpodOK() *ContainerStatsLibpodOK {
	return &ContainerStatsLibpodOK{}
}

/* ContainerStatsLibpodOK describes a response with status code 200, with default header values.

no error
*/
type ContainerStatsLibpodOK struct {
}

func (o *ContainerStatsLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/stats][%d] containerStatsLibpodOK ", 200)
}

func (o *ContainerStatsLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerStatsLibpodNotFound creates a ContainerStatsLibpodNotFound with default headers values
func NewContainerStatsLibpodNotFound() *ContainerStatsLibpodNotFound {
	return &ContainerStatsLibpodNotFound{}
}

/* ContainerStatsLibpodNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerStatsLibpodNotFound struct {
	Payload *ContainerStatsLibpodNotFoundBody
}

func (o *ContainerStatsLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/stats][%d] containerStatsLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ContainerStatsLibpodNotFound) GetPayload() *ContainerStatsLibpodNotFoundBody {
	return o.Payload
}

func (o *ContainerStatsLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerStatsLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerStatsLibpodConflict creates a ContainerStatsLibpodConflict with default headers values
func NewContainerStatsLibpodConflict() *ContainerStatsLibpodConflict {
	return &ContainerStatsLibpodConflict{}
}

/* ContainerStatsLibpodConflict describes a response with status code 409, with default header values.

Conflict error in operation
*/
type ContainerStatsLibpodConflict struct {
	Payload *ContainerStatsLibpodConflictBody
}

func (o *ContainerStatsLibpodConflict) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/stats][%d] containerStatsLibpodConflict  %+v", 409, o.Payload)
}
func (o *ContainerStatsLibpodConflict) GetPayload() *ContainerStatsLibpodConflictBody {
	return o.Payload
}

func (o *ContainerStatsLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerStatsLibpodConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerStatsLibpodInternalServerError creates a ContainerStatsLibpodInternalServerError with default headers values
func NewContainerStatsLibpodInternalServerError() *ContainerStatsLibpodInternalServerError {
	return &ContainerStatsLibpodInternalServerError{}
}

/* ContainerStatsLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerStatsLibpodInternalServerError struct {
	Payload *ContainerStatsLibpodInternalServerErrorBody
}

func (o *ContainerStatsLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/containers/{name}/stats][%d] containerStatsLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerStatsLibpodInternalServerError) GetPayload() *ContainerStatsLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerStatsLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerStatsLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerStatsLibpodConflictBody container stats libpod conflict body
swagger:model ContainerStatsLibpodConflictBody
*/
type ContainerStatsLibpodConflictBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container stats libpod conflict body
func (o *ContainerStatsLibpodConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container stats libpod conflict body based on context it is used
func (o *ContainerStatsLibpodConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerStatsLibpodConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerStatsLibpodConflictBody) UnmarshalBinary(b []byte) error {
	var res ContainerStatsLibpodConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerStatsLibpodInternalServerErrorBody container stats libpod internal server error body
swagger:model ContainerStatsLibpodInternalServerErrorBody
*/
type ContainerStatsLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container stats libpod internal server error body
func (o *ContainerStatsLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container stats libpod internal server error body based on context it is used
func (o *ContainerStatsLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerStatsLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerStatsLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerStatsLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerStatsLibpodNotFoundBody container stats libpod not found body
swagger:model ContainerStatsLibpodNotFoundBody
*/
type ContainerStatsLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container stats libpod not found body
func (o *ContainerStatsLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container stats libpod not found body based on context it is used
func (o *ContainerStatsLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerStatsLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerStatsLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerStatsLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
