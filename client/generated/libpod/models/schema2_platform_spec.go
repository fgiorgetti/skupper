// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Schema2PlatformSpec Schema2PlatformSpec describes the platform which a particular manifest is
// specialized for.
//
// swagger:model Schema2PlatformSpec
type Schema2PlatformSpec struct {

	// architecture
	Architecture string `json:"architecture,omitempty"`

	// features
	Features []string `json:"features"`

	// o s
	OS string `json:"os,omitempty"`

	// o s features
	OSFeatures []string `json:"os.features"`

	// o s version
	OSVersion string `json:"os.version,omitempty"`

	// variant
	Variant string `json:"variant,omitempty"`
}

// Validate validates this schema2 platform spec
func (m *Schema2PlatformSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schema2 platform spec based on context it is used
func (m *Schema2PlatformSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Schema2PlatformSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Schema2PlatformSpec) UnmarshalBinary(b []byte) error {
	var res Schema2PlatformSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
