// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LinuxBlockIO LinuxBlockIO for Linux cgroup 'blkio' resource management
//
// swagger:model LinuxBlockIO
type LinuxBlockIO struct {

	// Specifies tasks' weight in the given cgroup while competing with the cgroup's child cgroups, CFQ scheduler only
	LeafWeight uint16 `json:"leafWeight,omitempty"`

	// IO read rate limit per cgroup per device, bytes per second
	ThrottleReadBpsDevice []*LinuxThrottleDevice `json:"throttleReadBpsDevice"`

	// IO read rate limit per cgroup per device, IO per second
	ThrottleReadIOPSDevice []*LinuxThrottleDevice `json:"throttleReadIOPSDevice"`

	// IO write rate limit per cgroup per device, bytes per second
	ThrottleWriteBpsDevice []*LinuxThrottleDevice `json:"throttleWriteBpsDevice"`

	// IO write rate limit per cgroup per device, IO per second
	ThrottleWriteIOPSDevice []*LinuxThrottleDevice `json:"throttleWriteIOPSDevice"`

	// Specifies per cgroup weight
	Weight uint16 `json:"weight,omitempty"`

	// Weight per cgroup per device, can override BlkioWeight
	WeightDevice []*LinuxWeightDevice `json:"weightDevice"`
}

// Validate validates this linux block i o
func (m *LinuxBlockIO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateThrottleReadBpsDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThrottleReadIOPSDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThrottleWriteBpsDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThrottleWriteIOPSDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinuxBlockIO) validateThrottleReadBpsDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.ThrottleReadBpsDevice) { // not required
		return nil
	}

	for i := 0; i < len(m.ThrottleReadBpsDevice); i++ {
		if swag.IsZero(m.ThrottleReadBpsDevice[i]) { // not required
			continue
		}

		if m.ThrottleReadBpsDevice[i] != nil {
			if err := m.ThrottleReadBpsDevice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("throttleReadBpsDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("throttleReadBpsDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LinuxBlockIO) validateThrottleReadIOPSDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.ThrottleReadIOPSDevice) { // not required
		return nil
	}

	for i := 0; i < len(m.ThrottleReadIOPSDevice); i++ {
		if swag.IsZero(m.ThrottleReadIOPSDevice[i]) { // not required
			continue
		}

		if m.ThrottleReadIOPSDevice[i] != nil {
			if err := m.ThrottleReadIOPSDevice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("throttleReadIOPSDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("throttleReadIOPSDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LinuxBlockIO) validateThrottleWriteBpsDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.ThrottleWriteBpsDevice) { // not required
		return nil
	}

	for i := 0; i < len(m.ThrottleWriteBpsDevice); i++ {
		if swag.IsZero(m.ThrottleWriteBpsDevice[i]) { // not required
			continue
		}

		if m.ThrottleWriteBpsDevice[i] != nil {
			if err := m.ThrottleWriteBpsDevice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("throttleWriteBpsDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("throttleWriteBpsDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LinuxBlockIO) validateThrottleWriteIOPSDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.ThrottleWriteIOPSDevice) { // not required
		return nil
	}

	for i := 0; i < len(m.ThrottleWriteIOPSDevice); i++ {
		if swag.IsZero(m.ThrottleWriteIOPSDevice[i]) { // not required
			continue
		}

		if m.ThrottleWriteIOPSDevice[i] != nil {
			if err := m.ThrottleWriteIOPSDevice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("throttleWriteIOPSDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("throttleWriteIOPSDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LinuxBlockIO) validateWeightDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.WeightDevice) { // not required
		return nil
	}

	for i := 0; i < len(m.WeightDevice); i++ {
		if swag.IsZero(m.WeightDevice[i]) { // not required
			continue
		}

		if m.WeightDevice[i] != nil {
			if err := m.WeightDevice[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weightDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("weightDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this linux block i o based on the context it is used
func (m *LinuxBlockIO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateThrottleReadBpsDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThrottleReadIOPSDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThrottleWriteBpsDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThrottleWriteIOPSDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWeightDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinuxBlockIO) contextValidateThrottleReadBpsDevice(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ThrottleReadBpsDevice); i++ {

		if m.ThrottleReadBpsDevice[i] != nil {
			if err := m.ThrottleReadBpsDevice[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("throttleReadBpsDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("throttleReadBpsDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LinuxBlockIO) contextValidateThrottleReadIOPSDevice(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ThrottleReadIOPSDevice); i++ {

		if m.ThrottleReadIOPSDevice[i] != nil {
			if err := m.ThrottleReadIOPSDevice[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("throttleReadIOPSDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("throttleReadIOPSDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LinuxBlockIO) contextValidateThrottleWriteBpsDevice(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ThrottleWriteBpsDevice); i++ {

		if m.ThrottleWriteBpsDevice[i] != nil {
			if err := m.ThrottleWriteBpsDevice[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("throttleWriteBpsDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("throttleWriteBpsDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LinuxBlockIO) contextValidateThrottleWriteIOPSDevice(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ThrottleWriteIOPSDevice); i++ {

		if m.ThrottleWriteIOPSDevice[i] != nil {
			if err := m.ThrottleWriteIOPSDevice[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("throttleWriteIOPSDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("throttleWriteIOPSDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LinuxBlockIO) contextValidateWeightDevice(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WeightDevice); i++ {

		if m.WeightDevice[i] != nil {
			if err := m.WeightDevice[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("weightDevice" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("weightDevice" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinuxBlockIO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinuxBlockIO) UnmarshalBinary(b []byte) error {
	var res LinuxBlockIO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
