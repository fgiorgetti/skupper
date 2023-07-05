// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Subnet subnet
//
// swagger:model Subnet
type Subnet struct {

	// Gateway IP for this Network.
	Gateway string `json:"gateway,omitempty"`

	// Subnet for this Network in CIDR form.
	Subnet string `json:"subnet,omitempty"`

	// lease range
	LeaseRange *LeaseRange `json:"lease_range,omitempty"`
}

// Validate validates this subnet
func (m *Subnet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLeaseRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subnet) validateLeaseRange(formats strfmt.Registry) error {
	if swag.IsZero(m.LeaseRange) { // not required
		return nil
	}

	if m.LeaseRange != nil {
		if err := m.LeaseRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lease_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lease_range")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this subnet based on the context it is used
func (m *Subnet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLeaseRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subnet) contextValidateLeaseRange(ctx context.Context, formats strfmt.Registry) error {

	if m.LeaseRange != nil {
		if err := m.LeaseRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lease_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lease_range")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subnet) UnmarshalBinary(b []byte) error {
	var res Subnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
