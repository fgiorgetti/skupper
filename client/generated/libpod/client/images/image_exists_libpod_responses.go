// Code generated by go-swagger; DO NOT EDIT.

package images

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

// ImageExistsLibpodReader is a Reader for the ImageExistsLibpod structure.
type ImageExistsLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageExistsLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewImageExistsLibpodNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImageExistsLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageExistsLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageExistsLibpodNoContent creates a ImageExistsLibpodNoContent with default headers values
func NewImageExistsLibpodNoContent() *ImageExistsLibpodNoContent {
	return &ImageExistsLibpodNoContent{}
}

/* ImageExistsLibpodNoContent describes a response with status code 204, with default header values.

image exists
*/
type ImageExistsLibpodNoContent struct {
}

func (o *ImageExistsLibpodNoContent) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/exists][%d] imageExistsLibpodNoContent ", 204)
}

func (o *ImageExistsLibpodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImageExistsLibpodNotFound creates a ImageExistsLibpodNotFound with default headers values
func NewImageExistsLibpodNotFound() *ImageExistsLibpodNotFound {
	return &ImageExistsLibpodNotFound{}
}

/* ImageExistsLibpodNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImageExistsLibpodNotFound struct {
	Payload *ImageExistsLibpodNotFoundBody
}

func (o *ImageExistsLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/exists][%d] imageExistsLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ImageExistsLibpodNotFound) GetPayload() *ImageExistsLibpodNotFoundBody {
	return o.Payload
}

func (o *ImageExistsLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageExistsLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageExistsLibpodInternalServerError creates a ImageExistsLibpodInternalServerError with default headers values
func NewImageExistsLibpodInternalServerError() *ImageExistsLibpodInternalServerError {
	return &ImageExistsLibpodInternalServerError{}
}

/* ImageExistsLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageExistsLibpodInternalServerError struct {
	Payload *ImageExistsLibpodInternalServerErrorBody
}

func (o *ImageExistsLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/images/{name}/exists][%d] imageExistsLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageExistsLibpodInternalServerError) GetPayload() *ImageExistsLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImageExistsLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageExistsLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ImageExistsLibpodInternalServerErrorBody image exists libpod internal server error body
swagger:model ImageExistsLibpodInternalServerErrorBody
*/
type ImageExistsLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image exists libpod internal server error body
func (o *ImageExistsLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image exists libpod internal server error body based on context it is used
func (o *ImageExistsLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageExistsLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageExistsLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImageExistsLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ImageExistsLibpodNotFoundBody image exists libpod not found body
swagger:model ImageExistsLibpodNotFoundBody
*/
type ImageExistsLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image exists libpod not found body
func (o *ImageExistsLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image exists libpod not found body based on context it is used
func (o *ImageExistsLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageExistsLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageExistsLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ImageExistsLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}