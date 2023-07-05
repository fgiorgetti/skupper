// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

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

// PutContainerArchiveReader is a Reader for the PutContainerArchive structure.
type PutContainerArchiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutContainerArchiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutContainerArchiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutContainerArchiveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutContainerArchiveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutContainerArchiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutContainerArchiveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutContainerArchiveOK creates a PutContainerArchiveOK with default headers values
func NewPutContainerArchiveOK() *PutContainerArchiveOK {
	return &PutContainerArchiveOK{}
}

/*
PutContainerArchiveOK describes a response with status code 200, with default header values.

no error
*/
type PutContainerArchiveOK struct {
}

// IsSuccess returns true when this put container archive o k response has a 2xx status code
func (o *PutContainerArchiveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put container archive o k response has a 3xx status code
func (o *PutContainerArchiveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put container archive o k response has a 4xx status code
func (o *PutContainerArchiveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put container archive o k response has a 5xx status code
func (o *PutContainerArchiveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put container archive o k response a status code equal to that given
func (o *PutContainerArchiveOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutContainerArchiveOK) Error() string {
	return fmt.Sprintf("[PUT /containers/{name}/archive][%d] putContainerArchiveOK ", 200)
}

func (o *PutContainerArchiveOK) String() string {
	return fmt.Sprintf("[PUT /containers/{name}/archive][%d] putContainerArchiveOK ", 200)
}

func (o *PutContainerArchiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutContainerArchiveBadRequest creates a PutContainerArchiveBadRequest with default headers values
func NewPutContainerArchiveBadRequest() *PutContainerArchiveBadRequest {
	return &PutContainerArchiveBadRequest{}
}

/*
PutContainerArchiveBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type PutContainerArchiveBadRequest struct {
	Payload *PutContainerArchiveBadRequestBody
}

// IsSuccess returns true when this put container archive bad request response has a 2xx status code
func (o *PutContainerArchiveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put container archive bad request response has a 3xx status code
func (o *PutContainerArchiveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put container archive bad request response has a 4xx status code
func (o *PutContainerArchiveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put container archive bad request response has a 5xx status code
func (o *PutContainerArchiveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put container archive bad request response a status code equal to that given
func (o *PutContainerArchiveBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutContainerArchiveBadRequest) Error() string {
	return fmt.Sprintf("[PUT /containers/{name}/archive][%d] putContainerArchiveBadRequest  %+v", 400, o.Payload)
}

func (o *PutContainerArchiveBadRequest) String() string {
	return fmt.Sprintf("[PUT /containers/{name}/archive][%d] putContainerArchiveBadRequest  %+v", 400, o.Payload)
}

func (o *PutContainerArchiveBadRequest) GetPayload() *PutContainerArchiveBadRequestBody {
	return o.Payload
}

func (o *PutContainerArchiveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutContainerArchiveBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContainerArchiveForbidden creates a PutContainerArchiveForbidden with default headers values
func NewPutContainerArchiveForbidden() *PutContainerArchiveForbidden {
	return &PutContainerArchiveForbidden{}
}

/*
PutContainerArchiveForbidden describes a response with status code 403, with default header values.

the container rootfs is read-only
*/
type PutContainerArchiveForbidden struct {
}

// IsSuccess returns true when this put container archive forbidden response has a 2xx status code
func (o *PutContainerArchiveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put container archive forbidden response has a 3xx status code
func (o *PutContainerArchiveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put container archive forbidden response has a 4xx status code
func (o *PutContainerArchiveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put container archive forbidden response has a 5xx status code
func (o *PutContainerArchiveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put container archive forbidden response a status code equal to that given
func (o *PutContainerArchiveForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutContainerArchiveForbidden) Error() string {
	return fmt.Sprintf("[PUT /containers/{name}/archive][%d] putContainerArchiveForbidden ", 403)
}

func (o *PutContainerArchiveForbidden) String() string {
	return fmt.Sprintf("[PUT /containers/{name}/archive][%d] putContainerArchiveForbidden ", 403)
}

func (o *PutContainerArchiveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutContainerArchiveNotFound creates a PutContainerArchiveNotFound with default headers values
func NewPutContainerArchiveNotFound() *PutContainerArchiveNotFound {
	return &PutContainerArchiveNotFound{}
}

/*
PutContainerArchiveNotFound describes a response with status code 404, with default header values.

No such container
*/
type PutContainerArchiveNotFound struct {
	Payload *PutContainerArchiveNotFoundBody
}

// IsSuccess returns true when this put container archive not found response has a 2xx status code
func (o *PutContainerArchiveNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put container archive not found response has a 3xx status code
func (o *PutContainerArchiveNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put container archive not found response has a 4xx status code
func (o *PutContainerArchiveNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put container archive not found response has a 5xx status code
func (o *PutContainerArchiveNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put container archive not found response a status code equal to that given
func (o *PutContainerArchiveNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutContainerArchiveNotFound) Error() string {
	return fmt.Sprintf("[PUT /containers/{name}/archive][%d] putContainerArchiveNotFound  %+v", 404, o.Payload)
}

func (o *PutContainerArchiveNotFound) String() string {
	return fmt.Sprintf("[PUT /containers/{name}/archive][%d] putContainerArchiveNotFound  %+v", 404, o.Payload)
}

func (o *PutContainerArchiveNotFound) GetPayload() *PutContainerArchiveNotFoundBody {
	return o.Payload
}

func (o *PutContainerArchiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutContainerArchiveNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContainerArchiveInternalServerError creates a PutContainerArchiveInternalServerError with default headers values
func NewPutContainerArchiveInternalServerError() *PutContainerArchiveInternalServerError {
	return &PutContainerArchiveInternalServerError{}
}

/*
PutContainerArchiveInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PutContainerArchiveInternalServerError struct {
	Payload *PutContainerArchiveInternalServerErrorBody
}

// IsSuccess returns true when this put container archive internal server error response has a 2xx status code
func (o *PutContainerArchiveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put container archive internal server error response has a 3xx status code
func (o *PutContainerArchiveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put container archive internal server error response has a 4xx status code
func (o *PutContainerArchiveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put container archive internal server error response has a 5xx status code
func (o *PutContainerArchiveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put container archive internal server error response a status code equal to that given
func (o *PutContainerArchiveInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutContainerArchiveInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /containers/{name}/archive][%d] putContainerArchiveInternalServerError  %+v", 500, o.Payload)
}

func (o *PutContainerArchiveInternalServerError) String() string {
	return fmt.Sprintf("[PUT /containers/{name}/archive][%d] putContainerArchiveInternalServerError  %+v", 500, o.Payload)
}

func (o *PutContainerArchiveInternalServerError) GetPayload() *PutContainerArchiveInternalServerErrorBody {
	return o.Payload
}

func (o *PutContainerArchiveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutContainerArchiveInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutContainerArchiveBadRequestBody put container archive bad request body
swagger:model PutContainerArchiveBadRequestBody
*/
type PutContainerArchiveBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this put container archive bad request body
func (o *PutContainerArchiveBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put container archive bad request body based on context it is used
func (o *PutContainerArchiveBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutContainerArchiveBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutContainerArchiveBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutContainerArchiveBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutContainerArchiveInternalServerErrorBody put container archive internal server error body
swagger:model PutContainerArchiveInternalServerErrorBody
*/
type PutContainerArchiveInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this put container archive internal server error body
func (o *PutContainerArchiveInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put container archive internal server error body based on context it is used
func (o *PutContainerArchiveInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutContainerArchiveInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutContainerArchiveInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PutContainerArchiveInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PutContainerArchiveNotFoundBody put container archive not found body
swagger:model PutContainerArchiveNotFoundBody
*/
type PutContainerArchiveNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this put container archive not found body
func (o *PutContainerArchiveNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put container archive not found body based on context it is used
func (o *PutContainerArchiveNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutContainerArchiveNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutContainerArchiveNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutContainerArchiveNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
