// Code generated by go-swagger; DO NOT EDIT.

package containers_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/skupperproject/skupper/client/generated/libpod/models"
)

// ContainerInspectReader is a Reader for the ContainerInspect structure.
type ContainerInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerInspectOK creates a ContainerInspectOK with default headers values
func NewContainerInspectOK() *ContainerInspectOK {
	return &ContainerInspectOK{}
}

/* ContainerInspectOK describes a response with status code 200, with default header values.

Inspect container
*/
type ContainerInspectOK struct {
	Payload *ContainerInspectOKBody
}

func (o *ContainerInspectOK) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/json][%d] containerInspectOK  %+v", 200, o.Payload)
}
func (o *ContainerInspectOK) GetPayload() *ContainerInspectOKBody {
	return o.Payload
}

func (o *ContainerInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerInspectOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerInspectNotFound creates a ContainerInspectNotFound with default headers values
func NewContainerInspectNotFound() *ContainerInspectNotFound {
	return &ContainerInspectNotFound{}
}

/* ContainerInspectNotFound describes a response with status code 404, with default header values.

No such container
*/
type ContainerInspectNotFound struct {
	Payload *ContainerInspectNotFoundBody
}

func (o *ContainerInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/json][%d] containerInspectNotFound  %+v", 404, o.Payload)
}
func (o *ContainerInspectNotFound) GetPayload() *ContainerInspectNotFoundBody {
	return o.Payload
}

func (o *ContainerInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerInspectNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerInspectInternalServerError creates a ContainerInspectInternalServerError with default headers values
func NewContainerInspectInternalServerError() *ContainerInspectInternalServerError {
	return &ContainerInspectInternalServerError{}
}

/* ContainerInspectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ContainerInspectInternalServerError struct {
	Payload *ContainerInspectInternalServerErrorBody
}

func (o *ContainerInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{name}/json][%d] containerInspectInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerInspectInternalServerError) GetPayload() *ContainerInspectInternalServerErrorBody {
	return o.Payload
}

func (o *ContainerInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerInspectInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerInspectInternalServerErrorBody container inspect internal server error body
swagger:model ContainerInspectInternalServerErrorBody
*/
type ContainerInspectInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container inspect internal server error body
func (o *ContainerInspectInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container inspect internal server error body based on context it is used
func (o *ContainerInspectInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerInspectInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerInspectInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ContainerInspectInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerInspectNotFoundBody container inspect not found body
swagger:model ContainerInspectNotFoundBody
*/
type ContainerInspectNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this container inspect not found body
func (o *ContainerInspectNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container inspect not found body based on context it is used
func (o *ContainerInspectNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerInspectNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerInspectNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ContainerInspectNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ContainerInspectOKBody container inspect o k body
swagger:model ContainerInspectOKBody
*/
type ContainerInspectOKBody struct {

	// app armor profile
	AppArmorProfile string `json:"AppArmorProfile,omitempty"`

	// args
	Args []string `json:"Args"`

	// config
	Config *models.Config `json:"Config,omitempty"`

	// created
	Created string `json:"Created,omitempty"`

	// driver
	Driver string `json:"Driver,omitempty"`

	// exec i ds
	ExecIDs []string `json:"ExecIDs"`

	// graph driver
	GraphDriver *models.GraphDriverData `json:"GraphDriver,omitempty"`

	// host config
	HostConfig *models.HostConfig `json:"HostConfig,omitempty"`

	// hostname path
	HostnamePath string `json:"HostnamePath,omitempty"`

	// hosts path
	HostsPath string `json:"HostsPath,omitempty"`

	// ID
	ID string `json:"Id,omitempty"`

	// image
	Image string `json:"Image,omitempty"`

	// log path
	LogPath string `json:"LogPath,omitempty"`

	// mount label
	MountLabel string `json:"MountLabel,omitempty"`

	// mounts
	Mounts []*models.MountPoint `json:"Mounts"`

	// name
	Name string `json:"Name,omitempty"`

	// network settings
	NetworkSettings *models.NetworkSettings `json:"NetworkSettings,omitempty"`

	// node
	Node *models.ContainerNode `json:"Node,omitempty"`

	// path
	Path string `json:"Path,omitempty"`

	// platform
	Platform string `json:"Platform,omitempty"`

	// process label
	ProcessLabel string `json:"ProcessLabel,omitempty"`

	// resolv conf path
	ResolvConfPath string `json:"ResolvConfPath,omitempty"`

	// restart count
	RestartCount int64 `json:"RestartCount,omitempty"`

	// size root fs
	SizeRootFs int64 `json:"SizeRootFs,omitempty"`

	// size rw
	SizeRw int64 `json:"SizeRw,omitempty"`

	// state
	State *models.ContainerState `json:"State,omitempty"`
}

// Validate validates this container inspect o k body
func (o *ContainerInspectOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGraphDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHostConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworkSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerInspectOKBody) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.Config) { // not required
		return nil
	}

	if o.Config != nil {
		if err := o.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "Config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "Config")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerInspectOKBody) validateGraphDriver(formats strfmt.Registry) error {
	if swag.IsZero(o.GraphDriver) { // not required
		return nil
	}

	if o.GraphDriver != nil {
		if err := o.GraphDriver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "GraphDriver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "GraphDriver")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerInspectOKBody) validateHostConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.HostConfig) { // not required
		return nil
	}

	if o.HostConfig != nil {
		if err := o.HostConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "HostConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "HostConfig")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerInspectOKBody) validateMounts(formats strfmt.Registry) error {
	if swag.IsZero(o.Mounts) { // not required
		return nil
	}

	for i := 0; i < len(o.Mounts); i++ {
		if swag.IsZero(o.Mounts[i]) { // not required
			continue
		}

		if o.Mounts[i] != nil {
			if err := o.Mounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containerInspectOK" + "." + "Mounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containerInspectOK" + "." + "Mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ContainerInspectOKBody) validateNetworkSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.NetworkSettings) { // not required
		return nil
	}

	if o.NetworkSettings != nil {
		if err := o.NetworkSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "NetworkSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "NetworkSettings")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerInspectOKBody) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(o.Node) { // not required
		return nil
	}

	if o.Node != nil {
		if err := o.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "Node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "Node")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerInspectOKBody) validateState(formats strfmt.Registry) error {
	if swag.IsZero(o.State) { // not required
		return nil
	}

	if o.State != nil {
		if err := o.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "State")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "State")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this container inspect o k body based on the context it is used
func (o *ContainerInspectOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateGraphDriver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateHostConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNetworkSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerInspectOKBody) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.Config != nil {
		if err := o.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "Config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "Config")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerInspectOKBody) contextValidateGraphDriver(ctx context.Context, formats strfmt.Registry) error {

	if o.GraphDriver != nil {
		if err := o.GraphDriver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "GraphDriver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "GraphDriver")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerInspectOKBody) contextValidateHostConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.HostConfig != nil {
		if err := o.HostConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "HostConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "HostConfig")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerInspectOKBody) contextValidateMounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Mounts); i++ {

		if o.Mounts[i] != nil {
			if err := o.Mounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containerInspectOK" + "." + "Mounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("containerInspectOK" + "." + "Mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ContainerInspectOKBody) contextValidateNetworkSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.NetworkSettings != nil {
		if err := o.NetworkSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "NetworkSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "NetworkSettings")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerInspectOKBody) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if o.Node != nil {
		if err := o.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "Node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "Node")
			}
			return err
		}
	}

	return nil
}

func (o *ContainerInspectOKBody) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if o.State != nil {
		if err := o.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("containerInspectOK" + "." + "State")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("containerInspectOK" + "." + "State")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ContainerInspectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerInspectOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerInspectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
