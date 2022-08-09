// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SystemDfVolumeReport SystemDfVolumeReport describes a volume and its size
//
// swagger:model SystemDfVolumeReport
type SystemDfVolumeReport struct {

	// links
	Links int64 `json:"Links,omitempty"`

	// reclaimable size
	ReclaimableSize int64 `json:"ReclaimableSize,omitempty"`

	// size
	Size int64 `json:"Size,omitempty"`

	// volume name
	VolumeName string `json:"VolumeName,omitempty"`
}

// Validate validates this system df volume report
func (m *SystemDfVolumeReport) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this system df volume report based on context it is used
func (m *SystemDfVolumeReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SystemDfVolumeReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemDfVolumeReport) UnmarshalBinary(b []byte) error {
	var res SystemDfVolumeReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}