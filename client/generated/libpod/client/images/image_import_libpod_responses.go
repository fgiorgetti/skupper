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

	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

// ImageImportLibpodReader is a Reader for the ImageImportLibpod structure.
type ImageImportLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageImportLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageImportLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageImportLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageImportLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageImportLibpodOK creates a ImageImportLibpodOK with default headers values
func NewImageImportLibpodOK() *ImageImportLibpodOK {
	return &ImageImportLibpodOK{}
}

/*
ImageImportLibpodOK describes a response with status code 200, with default header values.

Import response
*/
type ImageImportLibpodOK struct {
	Payload *models.ImageImportReport
}

// IsSuccess returns true when this image import libpod o k response has a 2xx status code
func (o *ImageImportLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image import libpod o k response has a 3xx status code
func (o *ImageImportLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image import libpod o k response has a 4xx status code
func (o *ImageImportLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image import libpod o k response has a 5xx status code
func (o *ImageImportLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image import libpod o k response a status code equal to that given
func (o *ImageImportLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *ImageImportLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/images/import][%d] imageImportLibpodOK  %+v", 200, o.Payload)
}

func (o *ImageImportLibpodOK) String() string {
	return fmt.Sprintf("[POST /libpod/images/import][%d] imageImportLibpodOK  %+v", 200, o.Payload)
}

func (o *ImageImportLibpodOK) GetPayload() *models.ImageImportReport {
	return o.Payload
}

func (o *ImageImportLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageImportReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageImportLibpodBadRequest creates a ImageImportLibpodBadRequest with default headers values
func NewImageImportLibpodBadRequest() *ImageImportLibpodBadRequest {
	return &ImageImportLibpodBadRequest{}
}

/*
ImageImportLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type ImageImportLibpodBadRequest struct {
	Payload *ImageImportLibpodBadRequestBody
}

// IsSuccess returns true when this image import libpod bad request response has a 2xx status code
func (o *ImageImportLibpodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image import libpod bad request response has a 3xx status code
func (o *ImageImportLibpodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image import libpod bad request response has a 4xx status code
func (o *ImageImportLibpodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image import libpod bad request response has a 5xx status code
func (o *ImageImportLibpodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image import libpod bad request response a status code equal to that given
func (o *ImageImportLibpodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ImageImportLibpodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/images/import][%d] imageImportLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ImageImportLibpodBadRequest) String() string {
	return fmt.Sprintf("[POST /libpod/images/import][%d] imageImportLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *ImageImportLibpodBadRequest) GetPayload() *ImageImportLibpodBadRequestBody {
	return o.Payload
}

func (o *ImageImportLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageImportLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageImportLibpodInternalServerError creates a ImageImportLibpodInternalServerError with default headers values
func NewImageImportLibpodInternalServerError() *ImageImportLibpodInternalServerError {
	return &ImageImportLibpodInternalServerError{}
}

/*
ImageImportLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageImportLibpodInternalServerError struct {
	Payload *ImageImportLibpodInternalServerErrorBody
}

// IsSuccess returns true when this image import libpod internal server error response has a 2xx status code
func (o *ImageImportLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image import libpod internal server error response has a 3xx status code
func (o *ImageImportLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image import libpod internal server error response has a 4xx status code
func (o *ImageImportLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image import libpod internal server error response has a 5xx status code
func (o *ImageImportLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image import libpod internal server error response a status code equal to that given
func (o *ImageImportLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ImageImportLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/images/import][%d] imageImportLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageImportLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/images/import][%d] imageImportLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageImportLibpodInternalServerError) GetPayload() *ImageImportLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *ImageImportLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageImportLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ImageImportLibpodBadRequestBody image import libpod bad request body
swagger:model ImageImportLibpodBadRequestBody
*/
type ImageImportLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image import libpod bad request body
func (o *ImageImportLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image import libpod bad request body based on context it is used
func (o *ImageImportLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageImportLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageImportLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ImageImportLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ImageImportLibpodInternalServerErrorBody image import libpod internal server error body
swagger:model ImageImportLibpodInternalServerErrorBody
*/
type ImageImportLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image import libpod internal server error body
func (o *ImageImportLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image import libpod internal server error body based on context it is used
func (o *ImageImportLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageImportLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageImportLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImageImportLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
