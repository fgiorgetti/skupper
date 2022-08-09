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

// ContainerExportReader is a Reader for the ContainerExport structure.
type ContainerExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerExportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerExportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerExportOK creates a ContainerExportOK with default headers values
func NewContainerExportOK() *ContainerExportOK {
	return &ContainerExportOK{}
}

/* ContainerExportOK describes a response with status code 200, with default header values.

tarball is returned in body
*/
type ContainerExportOK struct {
}

func (o *ContainerExportOK) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/export][%d] containerExportOK ", 200)
}

func (o *ContainerExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerExportNotFound creates a ContainerExportNotFound with default headers values
func NewContainerExportNotFound() *ContainerExportNotFound {
	return &ContainerExportNotFound{}
}

/* ContainerExportNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerExportNotFound struct {
	Payload *ContainerExportNotFoundBody
}

func (o *ContainerExportNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/export][%d] containerExportNotFound  %+v", 404, o.Payload)
}
func (o *ContainerExportNotFound) GetPayload() *ContainerExportNotFoundBody {
	return o.Payload
}

func (o *ContainerExportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerExportNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerExportInternalServerError creates a ContainerExportInternalServerError with default headers values
func NewContainerExportInternalServerError() *ContainerExportInternalServerError {
	return &ContainerExportInternalServerError{}
}

/* ContainerExportInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerExportInternalServerError struct {
	Payload *ContainerExportInternalServerErrorBody
}

func (o *ContainerExportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/export][%d] containerExportInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerExportInternalServerError) GetPayload() *ContainerExportInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerExportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerExportInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerExportInternalServerErrorBody container export internal server error body
swagger:model ContainerExportInternalServerErrorBody
*/
type ContainerExportInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container export internal server error body
func (o *ContainerExportInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container export internal server error body based on context it is used
func (o *ContainerExportInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerExportInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerExportInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerExportInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerExportNotFoundBody container export not found body
swagger:model ContainerExportNotFoundBody
*/
type ContainerExportNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container export not found body
func (o *ContainerExportNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container export not found body based on context it is used
func (o *ContainerExportNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerExportNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerExportNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerExportNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}