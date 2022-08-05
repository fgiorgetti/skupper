// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SlirpInfo SlirpInfo describes the slirp executable that
// is being being used.
//
// swagger:model SlirpInfo
type SlirpInfo struct {

	// executable
	Executable string `json:"executable,omitempty"`

	// package
	Package string `json:"package,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this slirp info
func (m *SlirpInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this slirp info based on context it is used
func (m *SlirpInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SlirpInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlirpInfo) UnmarshalBinary(b []byte) error {
	var res SlirpInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
