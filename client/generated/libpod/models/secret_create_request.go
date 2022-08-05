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

// SecretCreateRequest secret create request
//
// swagger:model SecretCreateRequest
type SecretCreateRequest struct {

	// Base64-url-safe-encoded (RFC 4648) data to store as secret.
	Data string `json:"Data,omitempty"`

	// driver
	Driver *SecretDriverSpec `json:"Driver,omitempty"`

	// User-defined name of the secret.
	Name string `json:"Name,omitempty"`
}

// Validate validates this secret create request
func (m *SecretCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriver(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretCreateRequest) validateDriver(formats strfmt.Registry) error {
	if swag.IsZero(m.Driver) { // not required
		return nil
	}

	if m.Driver != nil {
		if err := m.Driver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Driver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Driver")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this secret create request based on the context it is used
func (m *SecretCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDriver(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretCreateRequest) contextValidateDriver(ctx context.Context, formats strfmt.Registry) error {

	if m.Driver != nil {
		if err := m.Driver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Driver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Driver")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecretCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretCreateRequest) UnmarshalBinary(b []byte) error {
	var res SecretCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
