// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

// SecretInspectLibpodReader is a Reader for the SecretInspectLibpod structure.
type SecretInspectLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretInspectLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretInspectLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSecretInspectLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretInspectLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSecretInspectLibpodOK creates a SecretInspectLibpodOK with default headers values
func NewSecretInspectLibpodOK() *SecretInspectLibpodOK {
	return &SecretInspectLibpodOK{}
}

/*
SecretInspectLibpodOK describes a response with status code 200, with default header values.

Secret inspect response
*/
type SecretInspectLibpodOK struct {
	Payload *models.SecretInfoReport
}

// IsSuccess returns true when this secret inspect libpod o k response has a 2xx status code
func (o *SecretInspectLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this secret inspect libpod o k response has a 3xx status code
func (o *SecretInspectLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret inspect libpod o k response has a 4xx status code
func (o *SecretInspectLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this secret inspect libpod o k response has a 5xx status code
func (o *SecretInspectLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this secret inspect libpod o k response a status code equal to that given
func (o *SecretInspectLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *SecretInspectLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/secrets/{name}/json][%d] secretInspectLibpodOK  %+v", 200, o.Payload)
}

func (o *SecretInspectLibpodOK) String() string {
	return fmt.Sprintf("[GET /libpod/secrets/{name}/json][%d] secretInspectLibpodOK  %+v", 200, o.Payload)
}

func (o *SecretInspectLibpodOK) GetPayload() *models.SecretInfoReport {
	return o.Payload
}

func (o *SecretInspectLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretInfoReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretInspectLibpodNotFound creates a SecretInspectLibpodNotFound with default headers values
func NewSecretInspectLibpodNotFound() *SecretInspectLibpodNotFound {
	return &SecretInspectLibpodNotFound{}
}

/*
SecretInspectLibpodNotFound describes a response with status code 404, with default header values.

No such secret
*/
type SecretInspectLibpodNotFound struct {
	Payload *SecretInspectLibpodNotFoundBody
}

// IsSuccess returns true when this secret inspect libpod not found response has a 2xx status code
func (o *SecretInspectLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secret inspect libpod not found response has a 3xx status code
func (o *SecretInspectLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret inspect libpod not found response has a 4xx status code
func (o *SecretInspectLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this secret inspect libpod not found response has a 5xx status code
func (o *SecretInspectLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this secret inspect libpod not found response a status code equal to that given
func (o *SecretInspectLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SecretInspectLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/secrets/{name}/json][%d] secretInspectLibpodNotFound  %+v", 404, o.Payload)
}

func (o *SecretInspectLibpodNotFound) String() string {
	return fmt.Sprintf("[GET /libpod/secrets/{name}/json][%d] secretInspectLibpodNotFound  %+v", 404, o.Payload)
}

func (o *SecretInspectLibpodNotFound) GetPayload() *SecretInspectLibpodNotFoundBody {
	return o.Payload
}

func (o *SecretInspectLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SecretInspectLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretInspectLibpodInternalServerError creates a SecretInspectLibpodInternalServerError with default headers values
func NewSecretInspectLibpodInternalServerError() *SecretInspectLibpodInternalServerError {
	return &SecretInspectLibpodInternalServerError{}
}

/*
SecretInspectLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretInspectLibpodInternalServerError struct {
	Payload *SecretInspectLibpodInternalServerErrorBody
}

// IsSuccess returns true when this secret inspect libpod internal server error response has a 2xx status code
func (o *SecretInspectLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this secret inspect libpod internal server error response has a 3xx status code
func (o *SecretInspectLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this secret inspect libpod internal server error response has a 4xx status code
func (o *SecretInspectLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this secret inspect libpod internal server error response has a 5xx status code
func (o *SecretInspectLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this secret inspect libpod internal server error response a status code equal to that given
func (o *SecretInspectLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SecretInspectLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/secrets/{name}/json][%d] secretInspectLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretInspectLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/secrets/{name}/json][%d] secretInspectLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretInspectLibpodInternalServerError) GetPayload() *SecretInspectLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *SecretInspectLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SecretInspectLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SecretInspectLibpodInternalServerErrorBody secret inspect libpod internal server error body
swagger:model SecretInspectLibpodInternalServerErrorBody
*/
type SecretInspectLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this secret inspect libpod internal server error body
func (o *SecretInspectLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret inspect libpod internal server error body based on context it is used
func (o *SecretInspectLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SecretInspectLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecretInspectLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SecretInspectLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SecretInspectLibpodNotFoundBody secret inspect libpod not found body
swagger:model SecretInspectLibpodNotFoundBody
*/
type SecretInspectLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this secret inspect libpod not found body
func (o *SecretInspectLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret inspect libpod not found body based on context it is used
func (o *SecretInspectLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SecretInspectLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecretInspectLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SecretInspectLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
