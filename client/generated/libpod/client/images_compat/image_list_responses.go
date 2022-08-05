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

	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

// ImageListReader is a Reader for the ImageList structure.
type ImageListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewImageListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageListOK creates a ImageListOK with default headers values
func NewImageListOK() *ImageListOK {
	return &ImageListOK{}
}

/* ImageListOK describes a response with status code 200, with default header values.

Image summary for compat API
*/
type ImageListOK struct {
	Payload []*models.ImageSummary
}

func (o *ImageListOK) Error() string {
	return fmt.Sprintf("[GET /images/json][%d] imageListOK  %+v", 200, o.Payload)
}
func (o *ImageListOK) GetPayload() []*models.ImageSummary {
	return o.Payload
}

func (o *ImageListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageListInternalServerError creates a ImageListInternalServerError with default headers values
func NewImageListInternalServerError() *ImageListInternalServerError {
	return &ImageListInternalServerError{}
}

/* ImageListInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ImageListInternalServerError struct {
	Payload *ImageListInternalServerErrorBody
}

func (o *ImageListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/json][%d] imageListInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageListInternalServerError) GetPayload() *ImageListInternalServerErrorBody {
	return o.Payload
}

func (o *ImageListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImageListInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ImageListInternalServerErrorBody image list internal server error body
swagger:model ImageListInternalServerErrorBody
*/
type ImageListInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this image list internal server error body
func (o *ImageListInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this image list internal server error body based on context it is used
func (o *ImageListInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ImageListInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImageListInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ImageListInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
