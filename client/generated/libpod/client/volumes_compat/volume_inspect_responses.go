// Code generated by go-swagger; DO NOT EDIT.

package volumes_compat

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

	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

// VolumeInspectReader is a Reader for the VolumeInspect structure.
type VolumeInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewVolumeInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeInspectOK creates a VolumeInspectOK with default headers values
func NewVolumeInspectOK() *VolumeInspectOK {
	return &VolumeInspectOK{}
}

/* VolumeInspectOK describes a response with status code 200, with default header values.

This response definition is used for both the create and inspect endpoints
*/
type VolumeInspectOK struct {
	Payload *VolumeInspectOKBody
}

func (o *VolumeInspectOK) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectOK  %+v", 200, o.Payload)
}
func (o *VolumeInspectOK) GetPayload() *VolumeInspectOKBody {
	return o.Payload
}

func (o *VolumeInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeInspectOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInspectNotFound creates a VolumeInspectNotFound with default headers values
func NewVolumeInspectNotFound() *VolumeInspectNotFound {
	return &VolumeInspectNotFound{}
}

/* VolumeInspectNotFound describes a response with status code 404, with default header values.

No such volume
*/
type VolumeInspectNotFound struct {
	Payload *VolumeInspectNotFoundBody
}

func (o *VolumeInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectNotFound  %+v", 404, o.Payload)
}
func (o *VolumeInspectNotFound) GetPayload() *VolumeInspectNotFoundBody {
	return o.Payload
}

func (o *VolumeInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeInspectNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInspectInternalServerError creates a VolumeInspectInternalServerError with default headers values
func NewVolumeInspectInternalServerError() *VolumeInspectInternalServerError {
	return &VolumeInspectInternalServerError{}
}

/* VolumeInspectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type VolumeInspectInternalServerError struct {
	Payload *VolumeInspectInternalServerErrorBody
}

func (o *VolumeInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectInternalServerError  %+v", 500, o.Payload)
}
func (o *VolumeInspectInternalServerError) GetPayload() *VolumeInspectInternalServerErrorBody {
	return o.Payload
}

func (o *VolumeInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeInspectInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*VolumeInspectInternalServerErrorBody volume inspect internal server error body
swagger:model VolumeInspectInternalServerErrorBody
*/
type VolumeInspectInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume inspect internal server error body
func (o *VolumeInspectInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume inspect internal server error body based on context it is used
func (o *VolumeInspectInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeInspectInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeInspectInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res VolumeInspectInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VolumeInspectNotFoundBody volume inspect not found body
swagger:model VolumeInspectNotFoundBody
*/
type VolumeInspectNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume inspect not found body
func (o *VolumeInspectNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume inspect not found body based on context it is used
func (o *VolumeInspectNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeInspectNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeInspectNotFoundBody) UnmarshalBinary(b []byte) error {
	var res VolumeInspectNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VolumeInspectOKBody volume inspect o k body
swagger:model VolumeInspectOKBody
*/
type VolumeInspectOKBody struct {

	// Date/Time the volume was created.
	CreatedAt string `json:"CreatedAt,omitempty"`

	// Name of the volume driver used by the volume.
	// Required: true
	Driver *string `json:"Driver"`

	// User-defined key/value metadata.
	// Required: true
	Labels map[string]string `json:"Labels"`

	// Mount path of the volume on the host.
	// Required: true
	Mountpoint *string `json:"Mountpoint"`

	// Name of the volume.
	// Required: true
	Name *string `json:"Name"`

	// The driver specific options used when creating the volume.
	// Required: true
	Options map[string]string `json:"Options"`

	// The level at which the volume exists. Either `global` for cluster-wide,
	// or `local` for machine level.
	// Required: true
	Scope *string `json:"Scope"`

	// Low-level details about the volume, provided by the volume driver.
	// Details are returned as a map with key/value pairs:
	// `{"key":"value","key2":"value2"}`.
	//
	// The `Status` field is optional, and is omitted if the volume driver
	// does not support this feature.
	Status map[string]interface{} `json:"Status,omitempty"`

	// usage data
	UsageData *models.VolumeUsageData `json:"UsageData,omitempty"`
}

// Validate validates this volume inspect o k body
func (o *VolumeInspectOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMountpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsageData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VolumeInspectOKBody) validateDriver(formats strfmt.Registry) error {

	if err := validate.Required("volumeInspectOK"+"."+"Driver", "body", o.Driver); err != nil {
		return err
	}

	return nil
}

func (o *VolumeInspectOKBody) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("volumeInspectOK"+"."+"Labels", "body", o.Labels); err != nil {
		return err
	}

	return nil
}

func (o *VolumeInspectOKBody) validateMountpoint(formats strfmt.Registry) error {

	if err := validate.Required("volumeInspectOK"+"."+"Mountpoint", "body", o.Mountpoint); err != nil {
		return err
	}

	return nil
}

func (o *VolumeInspectOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("volumeInspectOK"+"."+"Name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *VolumeInspectOKBody) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("volumeInspectOK"+"."+"Options", "body", o.Options); err != nil {
		return err
	}

	return nil
}

func (o *VolumeInspectOKBody) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("volumeInspectOK"+"."+"Scope", "body", o.Scope); err != nil {
		return err
	}

	return nil
}

func (o *VolumeInspectOKBody) validateUsageData(formats strfmt.Registry) error {
	if swag.IsZero(o.UsageData) { // not required
		return nil
	}

	if o.UsageData != nil {
		if err := o.UsageData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volumeInspectOK" + "." + "UsageData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volumeInspectOK" + "." + "UsageData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this volume inspect o k body based on the context it is used
func (o *VolumeInspectOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUsageData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VolumeInspectOKBody) contextValidateUsageData(ctx context.Context, formats strfmt.Registry) error {

	if o.UsageData != nil {
		if err := o.UsageData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volumeInspectOK" + "." + "UsageData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volumeInspectOK" + "." + "UsageData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VolumeInspectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeInspectOKBody) UnmarshalBinary(b []byte) error {
	var res VolumeInspectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
