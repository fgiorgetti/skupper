// Code generated by go-swagger; DO NOT EDIT.

package pods

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

// PodRestartLibpodReader is a Reader for the PodRestartLibpod structure.
type PodRestartLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PodRestartLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPodRestartLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPodRestartLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPodRestartLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPodRestartLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPodRestartLibpodOK creates a PodRestartLibpodOK with default headers values
func NewPodRestartLibpodOK() *PodRestartLibpodOK {
	return &PodRestartLibpodOK{}
}

/* PodRestartLibpodOK describes a response with status code 200, with default header values.

Restart pod
*/
type PodRestartLibpodOK struct {
	Payload *models.PodRestartReport
}

func (o *PodRestartLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/restart][%d] podRestartLibpodOK  %+v", 200, o.Payload)
}
func (o *PodRestartLibpodOK) GetPayload() *models.PodRestartReport {
	return o.Payload
}

func (o *PodRestartLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodRestartReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodRestartLibpodNotFound creates a PodRestartLibpodNotFound with default headers values
func NewPodRestartLibpodNotFound() *PodRestartLibpodNotFound {
	return &PodRestartLibpodNotFound{}
}

/* PodRestartLibpodNotFound describes a response with status code 404, with default header values.

No such pod
*/
type PodRestartLibpodNotFound struct {
	Payload *PodRestartLibpodNotFoundBody
}

func (o *PodRestartLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/restart][%d] podRestartLibpodNotFound  %+v", 404, o.Payload)
}
func (o *PodRestartLibpodNotFound) GetPayload() *PodRestartLibpodNotFoundBody {
	return o.Payload
}

func (o *PodRestartLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodRestartLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodRestartLibpodConflict creates a PodRestartLibpodConflict with default headers values
func NewPodRestartLibpodConflict() *PodRestartLibpodConflict {
	return &PodRestartLibpodConflict{}
}

/* PodRestartLibpodConflict describes a response with status code 409, with default header values.

Restart pod
*/
type PodRestartLibpodConflict struct {
	Payload *models.PodRestartReport
}

func (o *PodRestartLibpodConflict) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/restart][%d] podRestartLibpodConflict  %+v", 409, o.Payload)
}
func (o *PodRestartLibpodConflict) GetPayload() *models.PodRestartReport {
	return o.Payload
}

func (o *PodRestartLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodRestartReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodRestartLibpodInternalServerError creates a PodRestartLibpodInternalServerError with default headers values
func NewPodRestartLibpodInternalServerError() *PodRestartLibpodInternalServerError {
	return &PodRestartLibpodInternalServerError{}
}

/* PodRestartLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PodRestartLibpodInternalServerError struct {
	Payload *PodRestartLibpodInternalServerErrorBody
}

func (o *PodRestartLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/restart][%d] podRestartLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *PodRestartLibpodInternalServerError) GetPayload() *PodRestartLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *PodRestartLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodRestartLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PodRestartLibpodInternalServerErrorBody pod restart libpod internal server error body
swagger:model PodRestartLibpodInternalServerErrorBody
*/
type PodRestartLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod restart libpod internal server error body
func (o *PodRestartLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod restart libpod internal server error body based on context it is used
func (o *PodRestartLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodRestartLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodRestartLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PodRestartLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PodRestartLibpodNotFoundBody pod restart libpod not found body
swagger:model PodRestartLibpodNotFoundBody
*/
type PodRestartLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod restart libpod not found body
func (o *PodRestartLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod restart libpod not found body based on context it is used
func (o *PodRestartLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodRestartLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodRestartLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PodRestartLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}