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

	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

// PlayKubeLibpodReader is a Reader for the PlayKubeLibpod structure.
type PlayKubeLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlayKubeLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlayKubeLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPlayKubeLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPlayKubeLibpodOK creates a PlayKubeLibpodOK with default headers values
func NewPlayKubeLibpodOK() *PlayKubeLibpodOK {
	return &PlayKubeLibpodOK{}
}

/* PlayKubeLibpodOK describes a response with status code 200, with default header values.

PlayKube response
*/
type PlayKubeLibpodOK struct {
	Payload *models.PlayKubeReport
}

func (o *PlayKubeLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/play/kube][%d] playKubeLibpodOK  %+v", 200, o.Payload)
}
func (o *PlayKubeLibpodOK) GetPayload() *models.PlayKubeReport {
	return o.Payload
}

func (o *PlayKubeLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlayKubeReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlayKubeLibpodInternalServerError creates a PlayKubeLibpodInternalServerError with default headers values
func NewPlayKubeLibpodInternalServerError() *PlayKubeLibpodInternalServerError {
	return &PlayKubeLibpodInternalServerError{}
}

/* PlayKubeLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PlayKubeLibpodInternalServerError struct {
	Payload *PlayKubeLibpodInternalServerErrorBody
}

func (o *PlayKubeLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/play/kube][%d] playKubeLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *PlayKubeLibpodInternalServerError) GetPayload() *PlayKubeLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *PlayKubeLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PlayKubeLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PlayKubeLibpodInternalServerErrorBody play kube libpod internal server error body
swagger:model PlayKubeLibpodInternalServerErrorBody
*/
type PlayKubeLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this play kube libpod internal server error body
func (o *PlayKubeLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this play kube libpod internal server error body based on context it is used
func (o *PlayKubeLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PlayKubeLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PlayKubeLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PlayKubeLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
