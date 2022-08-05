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

// PodKillLibpodReader is a Reader for the PodKillLibpod structure.
type PodKillLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PodKillLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPodKillLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPodKillLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPodKillLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPodKillLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPodKillLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPodKillLibpodOK creates a PodKillLibpodOK with default headers values
func NewPodKillLibpodOK() *PodKillLibpodOK {
	return &PodKillLibpodOK{}
}

/* PodKillLibpodOK describes a response with status code 200, with default header values.

Kill Pod
*/
type PodKillLibpodOK struct {
	Payload *models.PodKillReport
}

func (o *PodKillLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/kill][%d] podKillLibpodOK  %+v", 200, o.Payload)
}
func (o *PodKillLibpodOK) GetPayload() *models.PodKillReport {
	return o.Payload
}

func (o *PodKillLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodKillReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodKillLibpodBadRequest creates a PodKillLibpodBadRequest with default headers values
func NewPodKillLibpodBadRequest() *PodKillLibpodBadRequest {
	return &PodKillLibpodBadRequest{}
}

/* PodKillLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type PodKillLibpodBadRequest struct {
	Payload *PodKillLibpodBadRequestBody
}

func (o *PodKillLibpodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/kill][%d] podKillLibpodBadRequest  %+v", 400, o.Payload)
}
func (o *PodKillLibpodBadRequest) GetPayload() *PodKillLibpodBadRequestBody {
	return o.Payload
}

func (o *PodKillLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodKillLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodKillLibpodNotFound creates a PodKillLibpodNotFound with default headers values
func NewPodKillLibpodNotFound() *PodKillLibpodNotFound {
	return &PodKillLibpodNotFound{}
}

/* PodKillLibpodNotFound describes a response with status code 404, with default header values.

No such pod
*/
type PodKillLibpodNotFound struct {
	Payload *PodKillLibpodNotFoundBody
}

func (o *PodKillLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/kill][%d] podKillLibpodNotFound  %+v", 404, o.Payload)
}
func (o *PodKillLibpodNotFound) GetPayload() *PodKillLibpodNotFoundBody {
	return o.Payload
}

func (o *PodKillLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodKillLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodKillLibpodConflict creates a PodKillLibpodConflict with default headers values
func NewPodKillLibpodConflict() *PodKillLibpodConflict {
	return &PodKillLibpodConflict{}
}

/* PodKillLibpodConflict describes a response with status code 409, with default header values.

Kill Pod
*/
type PodKillLibpodConflict struct {
	Payload *models.PodKillReport
}

func (o *PodKillLibpodConflict) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/kill][%d] podKillLibpodConflict  %+v", 409, o.Payload)
}
func (o *PodKillLibpodConflict) GetPayload() *models.PodKillReport {
	return o.Payload
}

func (o *PodKillLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodKillReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodKillLibpodInternalServerError creates a PodKillLibpodInternalServerError with default headers values
func NewPodKillLibpodInternalServerError() *PodKillLibpodInternalServerError {
	return &PodKillLibpodInternalServerError{}
}

/* PodKillLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PodKillLibpodInternalServerError struct {
	Payload *PodKillLibpodInternalServerErrorBody
}

func (o *PodKillLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/kill][%d] podKillLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *PodKillLibpodInternalServerError) GetPayload() *PodKillLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *PodKillLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodKillLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PodKillLibpodBadRequestBody pod kill libpod bad request body
swagger:model PodKillLibpodBadRequestBody
*/
type PodKillLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod kill libpod bad request body
func (o *PodKillLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod kill libpod bad request body based on context it is used
func (o *PodKillLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodKillLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodKillLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PodKillLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PodKillLibpodInternalServerErrorBody pod kill libpod internal server error body
swagger:model PodKillLibpodInternalServerErrorBody
*/
type PodKillLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod kill libpod internal server error body
func (o *PodKillLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod kill libpod internal server error body based on context it is used
func (o *PodKillLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodKillLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodKillLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PodKillLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PodKillLibpodNotFoundBody pod kill libpod not found body
swagger:model PodKillLibpodNotFoundBody
*/
type PodKillLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod kill libpod not found body
func (o *PodKillLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod kill libpod not found body based on context it is used
func (o *PodKillLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodKillLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodKillLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PodKillLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
