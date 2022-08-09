// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NetworkDisconnectLibpodReader is a Reader for the NetworkDisconnectLibpod structure.
type NetworkDisconnectLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkDisconnectLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkDisconnectLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewNetworkDisconnectLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNetworkDisconnectLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNetworkDisconnectLibpodOK creates a NetworkDisconnectLibpodOK with default headers values
func NewNetworkDisconnectLibpodOK() *NetworkDisconnectLibpodOK {
	return &NetworkDisconnectLibpodOK{}
}

/* NetworkDisconnectLibpodOK describes a response with status code 200, with default header values.

OK
*/
type NetworkDisconnectLibpodOK struct {
}

func (o *NetworkDisconnectLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/networks/{name}/disconnect][%d] networkDisconnectLibpodOK ", 200)
}

func (o *NetworkDisconnectLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkDisconnectLibpodNotFound creates a NetworkDisconnectLibpodNotFound with default headers values
func NewNetworkDisconnectLibpodNotFound() *NetworkDisconnectLibpodNotFound {
	return &NetworkDisconnectLibpodNotFound{}
}

/* NetworkDisconnectLibpodNotFound describes a response with status code 404, with default header values.

No such network
*/
type NetworkDisconnectLibpodNotFound struct {
	Payload *NetworkDisconnectLibpodNotFoundBody
}

func (o *NetworkDisconnectLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/networks/{name}/disconnect][%d] networkDisconnectLibpodNotFound  %+v", 404, o.Payload)
}
func (o *NetworkDisconnectLibpodNotFound) GetPayload() *NetworkDisconnectLibpodNotFoundBody {
	return o.Payload
}

func (o *NetworkDisconnectLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkDisconnectLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkDisconnectLibpodInternalServerError creates a NetworkDisconnectLibpodInternalServerError with default headers values
func NewNetworkDisconnectLibpodInternalServerError() *NetworkDisconnectLibpodInternalServerError {
	return &NetworkDisconnectLibpodInternalServerError{}
}

/* NetworkDisconnectLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type NetworkDisconnectLibpodInternalServerError struct {
	Payload *NetworkDisconnectLibpodInternalServerErrorBody
}

func (o *NetworkDisconnectLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/networks/{name}/disconnect][%d] networkDisconnectLibpodInternalServerError  %+v", 500, o.Payload)
}
func (o *NetworkDisconnectLibpodInternalServerError) GetPayload() *NetworkDisconnectLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *NetworkDisconnectLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(NetworkDisconnectLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*NetworkDisconnectLibpodInternalServerErrorBody network disconnect libpod internal server error body
swagger:model NetworkDisconnectLibpodInternalServerErrorBody
*/
type NetworkDisconnectLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network disconnect libpod internal server error body
func (o *NetworkDisconnectLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network disconnect libpod internal server error body based on context it is used
func (o *NetworkDisconnectLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkDisconnectLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkDisconnectLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res NetworkDisconnectLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*NetworkDisconnectLibpodNotFoundBody network disconnect libpod not found body
swagger:model NetworkDisconnectLibpodNotFoundBody
*/
type NetworkDisconnectLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this network disconnect libpod not found body
func (o *NetworkDisconnectLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network disconnect libpod not found body based on context it is used
func (o *NetworkDisconnectLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkDisconnectLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkDisconnectLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res NetworkDisconnectLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}