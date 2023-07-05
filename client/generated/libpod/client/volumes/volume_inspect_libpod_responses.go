// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumeInspectLibpodReader is a Reader for the VolumeInspectLibpod structure.
type VolumeInspectLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInspectLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInspectLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewVolumeInspectLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInspectLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeInspectLibpodOK creates a VolumeInspectLibpodOK with default headers values
func NewVolumeInspectLibpodOK() *VolumeInspectLibpodOK {
	return &VolumeInspectLibpodOK{}
}

/*
VolumeInspectLibpodOK describes a response with status code 200, with default header values.

Volume create response
*/
type VolumeInspectLibpodOK struct {
	Payload *VolumeInspectLibpodOKBody
}

// IsSuccess returns true when this volume inspect libpod o k response has a 2xx status code
func (o *VolumeInspectLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume inspect libpod o k response has a 3xx status code
func (o *VolumeInspectLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume inspect libpod o k response has a 4xx status code
func (o *VolumeInspectLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume inspect libpod o k response has a 5xx status code
func (o *VolumeInspectLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume inspect libpod o k response a status code equal to that given
func (o *VolumeInspectLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *VolumeInspectLibpodOK) Error() string {
	return fmt.Sprintf("[GET /libpod/volumes/{name}/json][%d] volumeInspectLibpodOK  %+v", 200, o.Payload)
}

func (o *VolumeInspectLibpodOK) String() string {
	return fmt.Sprintf("[GET /libpod/volumes/{name}/json][%d] volumeInspectLibpodOK  %+v", 200, o.Payload)
}

func (o *VolumeInspectLibpodOK) GetPayload() *VolumeInspectLibpodOKBody {
	return o.Payload
}

func (o *VolumeInspectLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeInspectLibpodOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInspectLibpodNotFound creates a VolumeInspectLibpodNotFound with default headers values
func NewVolumeInspectLibpodNotFound() *VolumeInspectLibpodNotFound {
	return &VolumeInspectLibpodNotFound{}
}

/*
VolumeInspectLibpodNotFound describes a response with status code 404, with default header values.

No such volume
*/
type VolumeInspectLibpodNotFound struct {
	Payload *VolumeInspectLibpodNotFoundBody
}

// IsSuccess returns true when this volume inspect libpod not found response has a 2xx status code
func (o *VolumeInspectLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume inspect libpod not found response has a 3xx status code
func (o *VolumeInspectLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume inspect libpod not found response has a 4xx status code
func (o *VolumeInspectLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume inspect libpod not found response has a 5xx status code
func (o *VolumeInspectLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this volume inspect libpod not found response a status code equal to that given
func (o *VolumeInspectLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *VolumeInspectLibpodNotFound) Error() string {
	return fmt.Sprintf("[GET /libpod/volumes/{name}/json][%d] volumeInspectLibpodNotFound  %+v", 404, o.Payload)
}

func (o *VolumeInspectLibpodNotFound) String() string {
	return fmt.Sprintf("[GET /libpod/volumes/{name}/json][%d] volumeInspectLibpodNotFound  %+v", 404, o.Payload)
}

func (o *VolumeInspectLibpodNotFound) GetPayload() *VolumeInspectLibpodNotFoundBody {
	return o.Payload
}

func (o *VolumeInspectLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeInspectLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInspectLibpodInternalServerError creates a VolumeInspectLibpodInternalServerError with default headers values
func NewVolumeInspectLibpodInternalServerError() *VolumeInspectLibpodInternalServerError {
	return &VolumeInspectLibpodInternalServerError{}
}

/*
VolumeInspectLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type VolumeInspectLibpodInternalServerError struct {
	Payload *VolumeInspectLibpodInternalServerErrorBody
}

// IsSuccess returns true when this volume inspect libpod internal server error response has a 2xx status code
func (o *VolumeInspectLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume inspect libpod internal server error response has a 3xx status code
func (o *VolumeInspectLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume inspect libpod internal server error response has a 4xx status code
func (o *VolumeInspectLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume inspect libpod internal server error response has a 5xx status code
func (o *VolumeInspectLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume inspect libpod internal server error response a status code equal to that given
func (o *VolumeInspectLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *VolumeInspectLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /libpod/volumes/{name}/json][%d] volumeInspectLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInspectLibpodInternalServerError) String() string {
	return fmt.Sprintf("[GET /libpod/volumes/{name}/json][%d] volumeInspectLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInspectLibpodInternalServerError) GetPayload() *VolumeInspectLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *VolumeInspectLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeInspectLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
VolumeInspectLibpodInternalServerErrorBody volume inspect libpod internal server error body
swagger:model VolumeInspectLibpodInternalServerErrorBody
*/
type VolumeInspectLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume inspect libpod internal server error body
func (o *VolumeInspectLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume inspect libpod internal server error body based on context it is used
func (o *VolumeInspectLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeInspectLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeInspectLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res VolumeInspectLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VolumeInspectLibpodNotFoundBody volume inspect libpod not found body
swagger:model VolumeInspectLibpodNotFoundBody
*/
type VolumeInspectLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume inspect libpod not found body
func (o *VolumeInspectLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume inspect libpod not found body based on context it is used
func (o *VolumeInspectLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeInspectLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeInspectLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res VolumeInspectLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
VolumeInspectLibpodOKBody volume inspect libpod o k body
swagger:model VolumeInspectLibpodOKBody
*/
type VolumeInspectLibpodOKBody struct {

	// Anonymous indicates that the volume was created as an anonymous
	// volume for a specific container, and will be be removed when any
	// container using it is removed.
	Anonymous bool `json:"Anonymous,omitempty"`

	// CreatedAt is the date and time the volume was created at. This is not
	// stored for older Libpod volumes; if so, it will be omitted.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"CreatedAt,omitempty"`

	// Driver is the driver used to create the volume.
	// If set to "local" or "", the Local driver (Podman built-in code) is
	// used to service the volume; otherwise, a volume plugin with the given
	// name is used to mount and manage the volume.
	Driver string `json:"Driver,omitempty"`

	// GID is the GID that the volume was created with.
	GID int64 `json:"GID,omitempty"`

	// Labels includes the volume's configured labels, key:value pairs that
	// can be passed during volume creation to provide information for third
	// party tools.
	Labels map[string]string `json:"Labels,omitempty"`

	// MountCount is the number of times this volume has been mounted.
	MountCount uint64 `json:"MountCount,omitempty"`

	// Mountpoint is the path on the host where the volume is mounted.
	Mountpoint string `json:"Mountpoint,omitempty"`

	// Name is the name of the volume.
	Name string `json:"Name,omitempty"`

	// NeedsChown indicates that the next time the volume is mounted into
	// a container, the container will chown the volume to the container process
	// UID/GID.
	NeedsChown bool `json:"NeedsChown,omitempty"`

	// NeedsCopyUp indicates that the next time the volume is mounted into
	NeedsCopyUp bool `json:"NeedsCopyUp,omitempty"`

	// Options is a set of options that were used when creating the volume.
	// For the Local driver, these are mount options that will be used to
	// determine how a local filesystem is mounted; they are handled as
	// parameters to Mount in a manner described in the volume create
	// manpage.
	// For non-local drivers, these are passed as-is to the volume plugin.
	Options map[string]string `json:"Options,omitempty"`

	// Scope is unused and provided solely for Docker compatibility. It is
	// unconditionally set to "local".
	Scope string `json:"Scope,omitempty"`

	// Status is used to return information on the volume's current state,
	// if the volume was created using a volume plugin (uses a Driver that
	// is not the local driver).
	// Status is provided to us by an external program, so no guarantees are
	// made about its format or contents. Further, it is an optional field,
	// so it may not be set even in cases where a volume plugin is in use.
	Status map[string]interface{} `json:"Status,omitempty"`

	// UID is the UID that the volume was created with.
	UID int64 `json:"UID,omitempty"`
}

// Validate validates this volume inspect libpod o k body
func (o *VolumeInspectLibpodOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VolumeInspectLibpodOKBody) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("volumeInspectLibpodOK"+"."+"CreatedAt", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this volume inspect libpod o k body based on context it is used
func (o *VolumeInspectLibpodOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeInspectLibpodOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeInspectLibpodOKBody) UnmarshalBinary(b []byte) error {
	var res VolumeInspectLibpodOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
