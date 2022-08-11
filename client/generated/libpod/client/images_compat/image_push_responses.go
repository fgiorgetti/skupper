// Code generated by go-swagger; DO NOT EDIT.

package images_compat

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

// ImagePushReader is a Reader for the ImagePush structure.
type ImagePushReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *ImagePushReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagePushOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImagePushNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImagePushInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagePushOK creates a ImagePushOK with default headers values
func NewImagePushOK(writer io.Writer) *ImagePushOK {
	return &ImagePushOK{

		Payload: writer,
	}
}

/* ImagePushOK describes a response with status code 200, with default header values.

no error
*/
type ImagePushOK struct {
	Payload io.Writer
}

func (o *ImagePushOK) Error() string {
	return fmt.Sprintf("[POST /images/{name}/push][%d] imagePushOK  %+v", 200, o.Payload)
}
func (o *ImagePushOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *ImagePushOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagePushNotFound creates a ImagePushNotFound with default headers values
func NewImagePushNotFound() *ImagePushNotFound {
	return &ImagePushNotFound{}
}

/* ImagePushNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImagePushNotFound struct {
	Payload *ImagePushNotFoundBody
}

func (o *ImagePushNotFound) Error() string {
	return fmt.Sprintf("[POST /images/{name}/push][%d] imagePushNotFound  %+v", 404, o.Payload)
}
func (o *ImagePushNotFound) GetPayload() *ImagePushNotFoundBody {
	return o.Payload
}

func (o *ImagePushNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagePushNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagePushInternalServerError creates a ImagePushInternalServerError with default headers values
func NewImagePushInternalServerError() *ImagePushInternalServerError {
	return &ImagePushInternalServerError{}
}

/* ImagePushInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImagePushInternalServerError struct {
	Payload *ImagePushInternalServerErrorBody
}

func (o *ImagePushInternalServerError) Error() string {
	return fmt.Sprintf("[POST /images/{name}/push][%d] imagePushInternalServerError  %+v", 500, o.Payload)
}
func (o *ImagePushInternalServerError) GetPayload() *ImagePushInternalServerErrorBody {
	return o.Payload
}

func (o *ImagePushInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagePushInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ImagePushInternalServerErrorBody image push internal server error body
swagger:model ImagePushInternalServerErrorBody
*/
type ImagePushInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image push internal server error body
func (o *ImagePushInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image push internal server error body based on context it is used
func (o *ImagePushInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagePushInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagePushInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImagePushInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ImagePushNotFoundBody image push not found body
swagger:model ImagePushNotFoundBody
*/
type ImagePushNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image push not found body
func (o *ImagePushNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image push not found body based on context it is used
func (o *ImagePushNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImagePushNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagePushNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ImagePushNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}