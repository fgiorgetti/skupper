// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Ulimit Ulimit is a human friendly version of Rlimit.
//
// swagger:model Ulimit
type Ulimit struct {

	// hard
	Hard int64 `json:"Hard,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// soft
	Soft int64 `json:"Soft,omitempty"`
}

// Validate validates this ulimit
func (m *Ulimit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ulimit based on context it is used
func (m *Ulimit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Ulimit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Ulimit) UnmarshalBinary(b []byte) error {
	var res Ulimit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}