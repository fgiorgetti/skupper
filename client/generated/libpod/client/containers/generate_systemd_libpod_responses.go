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

// GenerateSystemdLibpodReader is a Reader for the GenerateSystemdLibpod structure.
type GenerateSystemdLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateSystemdLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateSystemdLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGenerateSystemdLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateSystemdLibpodOK creates a GenerateSystemdLibpodOK with default headers values
func NewGenerateSystemdLibpodOK() *GenerateSystemdLibpodOK {
	return &GenerateSystemdLibpodOK{}
}

/* GenerateSystemdLibpodOK describes a response with status code 200, with default header values.

no error
*/
type GenerateSystemdLibpodOK struct {
	Payload map[string]string
}

func (o *GenerateSystemdLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/generate/{name}/systemd][%d] generateSystemdLibpodOK  %+v", 200, o.Payload)
}
func (o *GenerateSystemdLibpodOK) GetPayload() map[string]string {
	return o.Payload
}

func (o *GenerateSystemdLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateSystemdLibpodInternalServerError creates a GenerateSystemdLibpodInternalServerError with default headers values
func NewGenerateSystemdLibpodInternalServerError() *GenerateSystemdLibpodInternalServerError {
	return &GenerateSystemdLibpodInternalServerError{}
}

/* GenerateSystemdLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GenerateSystemdLibpodInternalServerError struct {
	Payload *GenerateSystemdLibpodInternalServerErrorBody
}

func (o *GenerateSystemdLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/generate/{name}/systemd][%d] generateSystemdLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *GenerateSystemdLibpodInternalServerError) GetPayload() *GenerateSystemdLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *GenerateSystemdLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GenerateSystemdLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GenerateSystemdLibpodInternalServerErrorBody generate systemd libpod internal server error body
swagger:model GenerateSystemdLibpodInternalServerErrorBody
*/
type GenerateSystemdLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this generate systemd libpod internal server error body
func (o *GenerateSystemdLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this generate systemd libpod internal server error body based on context it is used
func (o *GenerateSystemdLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GenerateSystemdLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GenerateSystemdLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GenerateSystemdLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
