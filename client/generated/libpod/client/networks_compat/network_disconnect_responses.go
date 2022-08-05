// Code generated by go-swagger; DO NOT EDIT.

package networks_compat

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

// NetworkDisconnectReader is a Reader for the NetworkDisconnect structure.
type NetworkDisconnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkDisconnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkDisconnectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNetworkDisconnectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNetworkDisconnectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNetworkDisconnectOK creates a NetworkDisconnectOK with default headers values
func NewNetworkDisconnectOK() *NetworkDisconnectOK {
	return &NetworkDisconnectOK{}
}

/* NetworkDisconnectOK describes a response with status code 200, with default header values.

OK
*/
type NetworkDisconnectOK struct {
}

func (o *NetworkDisconnectOK) Error() string {
	return fmt.Sprintf("[POST /networks/{name}/disconnect][%d] networkDisconnectOK ", 200)
}

func (o *NetworkDisconnectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkDisconnectBadRequest creates a NetworkDisconnectBadRequest with default headers values
func NewNetworkDisconnectBadRequest() *NetworkDisconnectBadRequest {
	return &NetworkDisconnectBadRequest{}
}

/* NetworkDisconnectBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type NetworkDisconnectBadRequest struct {
	Payload *NetworkDisconnectBadRequestBody
}

func (o *NetworkDisconnectBadRequest) Error() string {
	return fmt.Sprintf("[POST /networks/{name}/disconnect][%d] networkDisconnectBadRequest  %+v", 400, o.Payload)
}
func (o *NetworkDisconnectBadRequest) GetPayload() *NetworkDisconnectBadRequestBody {
	return o.Payload
}

func (o *NetworkDisconnectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkDisconnectBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkDisconnectInternalServerError creates a NetworkDisconnectInternalServerError with default headers values
func NewNetworkDisconnectInternalServerError() *NetworkDisconnectInternalServerError {
	return &NetworkDisconnectInternalServerError{}
}

/* NetworkDisconnectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type NetworkDisconnectInternalServerError struct {
	Payload *NetworkDisconnectInternalServerErrorBody
}

func (o *NetworkDisconnectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /networks/{name}/disconnect][%d] networkDisconnectInternalServerError  %+v", 500, o.Payload)
}
func (o *NetworkDisconnectInternalServerError) GetPayload() *NetworkDisconnectInternalServerErrorBody {
	return o.Payload
}

func (o *NetworkDisconnectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkDisconnectInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*NetworkDisconnectBadRequestBody network disconnect bad request body
swagger:model NetworkDisconnectBadRequestBody
*/
type NetworkDisconnectBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network disconnect bad request body
func (o *NetworkDisconnectBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network disconnect bad request body based on context it is used
func (o *NetworkDisconnectBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkDisconnectBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkDisconnectBadRequestBody) UnmarshalBinary(b []byte) error {
	var res NetworkDisconnectBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*NetworkDisconnectInternalServerErrorBody network disconnect internal server error body
swagger:model NetworkDisconnectInternalServerErrorBody
*/
type NetworkDisconnectInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network disconnect internal server error body
func (o *NetworkDisconnectInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network disconnect internal server error body based on context it is used
func (o *NetworkDisconnectInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkDisconnectInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkDisconnectInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NetworkDisconnectInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
