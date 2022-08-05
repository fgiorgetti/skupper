// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LinuxBlockIODevice linuxBlockIODevice holds major:minor format supported in blkio cgroup
//
// swagger:model linuxBlockIODevice
type LinuxBlockIODevice struct {

	// Major is the device's major number.
	Major int64 `json:"major,omitempty"`

	// Minor is the device's minor number.
	Minor int64 `json:"minor,omitempty"`
}

// Validate validates this linux block i o device
func (m *LinuxBlockIODevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this linux block i o device based on context it is used
func (m *LinuxBlockIODevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinuxBlockIODevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinuxBlockIODevice) UnmarshalBinary(b []byte) error {
	var res LinuxBlockIODevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
