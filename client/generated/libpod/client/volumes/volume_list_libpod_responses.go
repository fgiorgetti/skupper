// Code generated by go-swagger; DO NOT EDIT.

package volumes

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

// VolumeListLibpodReader is a Reader for the VolumeListLibpod structure.
type VolumeListLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeListLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeListLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewVolumeListLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeListLibpodOK creates a VolumeListLibpodOK with default headers values
func NewVolumeListLibpodOK() *VolumeListLibpodOK {
	return &VolumeListLibpodOK{}
}

/* VolumeListLibpodOK describes a response with status code 200, with default header values.

Volume list
*/
type VolumeListLibpodOK struct {
	Payload []*models.VolumeConfigResponse
}

func (o *VolumeListLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/volumes/json][%d] volumeListLibpodOK  %+v", 200, o.Payload)
}
func (o *VolumeListLibpodOK) GetPayload() []*models.VolumeConfigResponse {
	return o.Payload
}

func (o *VolumeListLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeListLibpodInternalServerError creates a VolumeListLibpodInternalServerError with default headers values
func NewVolumeListLibpodInternalServerError() *VolumeListLibpodInternalServerError {
	return &VolumeListLibpodInternalServerError{}
}

/* VolumeListLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type VolumeListLibpodInternalServerError struct {
	Payload *VolumeListLibpodInternalServerErrorBody
}

func (o *VolumeListLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/volumes/json][%d] volumeListLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *VolumeListLibpodInternalServerError) GetPayload() *VolumeListLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *VolumeListLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeListLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*VolumeListLibpodInternalServerErrorBody volume list libpod internal server error body
swagger:model VolumeListLibpodInternalServerErrorBody
*/
type VolumeListLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume list libpod internal server error body
func (o *VolumeListLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume list libpod internal server error body based on context it is used
func (o *VolumeListLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeListLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeListLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res VolumeListLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}