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

// ImageTagLibpodReader is a Reader for the ImageTagLibpod structure.
type ImageTagLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageTagLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewImageTagLibpodCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageTagLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageTagLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImageTagLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageTagLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageTagLibpodCreated creates a ImageTagLibpodCreated with default headers values
func NewImageTagLibpodCreated() *ImageTagLibpodCreated {
	return &ImageTagLibpodCreated{}
}

/* ImageTagLibpodCreated describes a response with status code 201, with default header values.

no error
*/
type ImageTagLibpodCreated struct {
}

func (o *ImageTagLibpodCreated) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/tag][%d] imageTagLibpodCreated ", 201)
}

func (o *ImageTagLibpodCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImageTagLibpodBadRequest creates a ImageTagLibpodBadRequest with default headers values
func NewImageTagLibpodBadRequest() *ImageTagLibpodBadRequest {
	return &ImageTagLibpodBadRequest{}
}

/* ImageTagLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type ImageTagLibpodBadRequest struct {
	Payload *ImageTagLibpodBadRequestBody
}

func (o *ImageTagLibpodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/tag][%d] imageTagLibpodBadRequest  %+v", 400, o.Payload)
}
func (o *ImageTagLibpodBadRequest) GetPayload() *ImageTagLibpodBadRequestBody {
	return o.Payload
}

func (o *ImageTagLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageTagLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageTagLibpodNotFound creates a ImageTagLibpodNotFound with default headers values
func NewImageTagLibpodNotFound() *ImageTagLibpodNotFound {
	return &ImageTagLibpodNotFound{}
}

/* ImageTagLibpodNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImageTagLibpodNotFound struct {
	Payload *ImageTagLibpodNotFoundBody
}

func (o *ImageTagLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/tag][%d] imageTagLibpodNotFound  %+v", 404, o.Payload)
}
func (o *ImageTagLibpodNotFound) GetPayload() *ImageTagLibpodNotFoundBody {
	return o.Payload
}

func (o *ImageTagLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageTagLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageTagLibpodConflict creates a ImageTagLibpodConflict with default headers values
func NewImageTagLibpodConflict() *ImageTagLibpodConflict {
	return &ImageTagLibpodConflict{}
}

/* ImageTagLibpodConflict describes a response with status code 409, with default header values.

Conflict error in operation
*/
type ImageTagLibpodConflict struct {
	Payload *ImageTagLibpodConflictBody
}

func (o *ImageTagLibpodConflict) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/tag][%d] imageTagLibpodConflict  %+v", 409, o.Payload)
}
func (o *ImageTagLibpodConflict) GetPayload() *ImageTagLibpodConflictBody {
	return o.Payload
}

func (o *ImageTagLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageTagLibpodConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageTagLibpodInternalServerError creates a ImageTagLibpodInternalServerError with default headers values
func NewImageTagLibpodInternalServerError() *ImageTagLibpodInternalServerError {
	return &ImageTagLibpodInternalServerError{}
}

/* ImageTagLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageTagLibpodInternalServerError struct {
	Payload *ImageTagLibpodInternalServerErrorBody
}

func (o *ImageTagLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/images/{name}/tag][%d] imageTagLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageTagLibpodInternalServerError) GetPayload() *ImageTagLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImageTagLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageTagLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ImageTagLibpodBadRequestBody image tag libpod bad request body
swagger:model ImageTagLibpodBadRequestBody
*/
type ImageTagLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image tag libpod bad request body
func (o *ImageTagLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image tag libpod bad request body based on context it is used
func (o *ImageTagLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageTagLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageTagLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ImageTagLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ImageTagLibpodConflictBody image tag libpod conflict body
swagger:model ImageTagLibpodConflictBody
*/
type ImageTagLibpodConflictBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image tag libpod conflict body
func (o *ImageTagLibpodConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image tag libpod conflict body based on context it is used
func (o *ImageTagLibpodConflictBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageTagLibpodConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageTagLibpodConflictBody) UnmarshalBinary(b []byte) error {
	var res ImageTagLibpodConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ImageTagLibpodInternalServerErrorBody image tag libpod internal server error body
swagger:model ImageTagLibpodInternalServerErrorBody
*/
type ImageTagLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image tag libpod internal server error body
func (o *ImageTagLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image tag libpod internal server error body based on context it is used
func (o *ImageTagLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageTagLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageTagLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImageTagLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ImageTagLibpodNotFoundBody image tag libpod not found body
swagger:model ImageTagLibpodNotFoundBody
*/
type ImageTagLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image tag libpod not found body
func (o *ImageTagLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image tag libpod not found body based on context it is used
func (o *ImageTagLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageTagLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageTagLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ImageTagLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}