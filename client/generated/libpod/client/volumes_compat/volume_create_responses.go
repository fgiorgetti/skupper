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

// VolumeCreateReader is a Reader for the VolumeCreate structure.
type VolumeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVolumeCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewVolumeCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeCreateCreated creates a VolumeCreateCreated with default headers values
func NewVolumeCreateCreated() *VolumeCreateCreated {
	return &VolumeCreateCreated{}
}

/* VolumeCreateCreated describes a response with status code 201, with default header values.

This response definition is used for both the create and inspect endpoints
*/
type VolumeCreateCreated struct {
	Payload *VolumeCreateCreatedBody
}

func (o *VolumeCreateCreated) Error() string {
	return fmt.Sprintf("[POST /volumes/create][%d] volumeCreateCreated  %+v", 201, o.Payload)
}
func (o *VolumeCreateCreated) GetPayload() *VolumeCreateCreatedBody {
	return o.Payload
}

func (o *VolumeCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeCreateCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeCreateInternalServerError creates a VolumeCreateInternalServerError with default headers values
func NewVolumeCreateInternalServerError() *VolumeCreateInternalServerError {
	return &VolumeCreateInternalServerError{}
}

/* VolumeCreateInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type VolumeCreateInternalServerError struct {
	Payload *VolumeCreateInternalServerErrorBody
}

func (o *VolumeCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /volumes/create][%d] volumeCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *VolumeCreateInternalServerError) GetPayload() *VolumeCreateInternalServerErrorBody {
	return o.Payload
}

func (o *VolumeCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VolumeCreateInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*VolumeCreateCreatedBody volume create created body
swagger:model VolumeCreateCreatedBody
*/
type VolumeCreateCreatedBody struct {

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

// Validate validates this volume create created body
func (o *VolumeCreateCreatedBody) Validate(formats strfmt.Registry) error {
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

func (o *VolumeCreateCreatedBody) validateDriver(formats strfmt.Registry) error {

	if err := validate.Required("volumeCreateCreated"+"."+"Driver", "body", o.Driver); err != nil {
		return err
	}

	return nil
}

func (o *VolumeCreateCreatedBody) validateLabels(formats strfmt.Registry) error {

	if err := validate.Required("volumeCreateCreated"+"."+"Labels", "body", o.Labels); err != nil {
		return err
	}

	return nil
}

func (o *VolumeCreateCreatedBody) validateMountpoint(formats strfmt.Registry) error {

	if err := validate.Required("volumeCreateCreated"+"."+"Mountpoint", "body", o.Mountpoint); err != nil {
		return err
	}

	return nil
}

func (o *VolumeCreateCreatedBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("volumeCreateCreated"+"."+"Name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *VolumeCreateCreatedBody) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("volumeCreateCreated"+"."+"Options", "body", o.Options); err != nil {
		return err
	}

	return nil
}

func (o *VolumeCreateCreatedBody) validateScope(formats strfmt.Registry) error {

	if err := validate.Required("volumeCreateCreated"+"."+"Scope", "body", o.Scope); err != nil {
		return err
	}

	return nil
}

func (o *VolumeCreateCreatedBody) validateUsageData(formats strfmt.Registry) error {
	if swag.IsZero(o.UsageData) { // not required
		return nil
	}

	if o.UsageData != nil {
		if err := o.UsageData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volumeCreateCreated" + "." + "UsageData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volumeCreateCreated" + "." + "UsageData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this volume create created body based on the context it is used
func (o *VolumeCreateCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUsageData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VolumeCreateCreatedBody) contextValidateUsageData(ctx context.Context, formats strfmt.Registry) error {

	if o.UsageData != nil {
		if err := o.UsageData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volumeCreateCreated" + "." + "UsageData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volumeCreateCreated" + "." + "UsageData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VolumeCreateCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeCreateCreatedBody) UnmarshalBinary(b []byte) error {
	var res VolumeCreateCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VolumeCreateInternalServerErrorBody volume create internal server error body
swagger:model VolumeCreateInternalServerErrorBody
*/
type VolumeCreateInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this volume create internal server error body
func (o *VolumeCreateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this volume create internal server error body based on context it is used
func (o *VolumeCreateInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VolumeCreateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VolumeCreateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res VolumeCreateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
