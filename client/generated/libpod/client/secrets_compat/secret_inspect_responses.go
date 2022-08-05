// Code generated by go-swagger; DO NOT EDIT.

package secrets_compat

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

// SecretInspectReader is a Reader for the SecretInspect structure.
type SecretInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSecretInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSecretInspectOK creates a SecretInspectOK with default headers values
func NewSecretInspectOK() *SecretInspectOK {
	return &SecretInspectOK{}
}

/* SecretInspectOK describes a response with status code 200, with default header values.

Secret inspect compat
*/
type SecretInspectOK struct {
	Payload *models.SecretInfoReportCompat
}

func (o *SecretInspectOK) Error() string {
	return fmt.Sprintf("[GET /secrets/{name}][%d] secretInspectOK  %+v", 200, o.Payload)
}
func (o *SecretInspectOK) GetPayload() *models.SecretInfoReportCompat {
	return o.Payload
}

func (o *SecretInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecretInfoReportCompat)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretInspectNotFound creates a SecretInspectNotFound with default headers values
func NewSecretInspectNotFound() *SecretInspectNotFound {
	return &SecretInspectNotFound{}
}

/* SecretInspectNotFound describes a response with status code 404, with default header values.

No such secret
*/
type SecretInspectNotFound struct {
	Payload *SecretInspectNotFoundBody
}

func (o *SecretInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /secrets/{name}][%d] secretInspectNotFound  %+v", 404, o.Payload)
}
func (o *SecretInspectNotFound) GetPayload() *SecretInspectNotFoundBody {
	return o.Payload
}

func (o *SecretInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SecretInspectNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretInspectInternalServerError creates a SecretInspectInternalServerError with default headers values
func NewSecretInspectInternalServerError() *SecretInspectInternalServerError {
	return &SecretInspectInternalServerError{}
}

/* SecretInspectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SecretInspectInternalServerError struct {
	Payload *SecretInspectInternalServerErrorBody
}

func (o *SecretInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /secrets/{name}][%d] secretInspectInternalServerError  %+v", 500, o.Payload)
}
func (o *SecretInspectInternalServerError) GetPayload() *SecretInspectInternalServerErrorBody {
	return o.Payload
}

func (o *SecretInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SecretInspectInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SecretInspectInternalServerErrorBody secret inspect internal server error body
swagger:model SecretInspectInternalServerErrorBody
*/
type SecretInspectInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this secret inspect internal server error body
func (o *SecretInspectInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret inspect internal server error body based on context it is used
func (o *SecretInspectInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SecretInspectInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecretInspectInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res SecretInspectInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SecretInspectNotFoundBody secret inspect not found body
swagger:model SecretInspectNotFoundBody
*/
type SecretInspectNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this secret inspect not found body
func (o *SecretInspectNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this secret inspect not found body based on context it is used
func (o *SecretInspectNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SecretInspectNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecretInspectNotFoundBody) UnmarshalBinary(b []byte) error {
	var res SecretInspectNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
