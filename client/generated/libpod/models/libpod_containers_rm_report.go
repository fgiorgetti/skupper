// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LibpodContainersRmReport libpod containers rm report
//
// swagger:model LibpodContainersRmReport
type LibpodContainersRmReport struct {

	// ID
	ID string `json:"Id,omitempty"`

	// Error which occurred during Rm operation (if any).
	// This field is optional and may be omitted if no error occurred.
	RmError *string `json:"Err,omitempty"`
}

// Validate validates this libpod containers rm report
func (m *LibpodContainersRmReport) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this libpod containers rm report based on context it is used
func (m *LibpodContainersRmReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LibpodContainersRmReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LibpodContainersRmReport) UnmarshalBinary(b []byte) error {
	var res LibpodContainersRmReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
