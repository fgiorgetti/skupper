// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InspectRestartPolicy InspectRestartPolicy holds information about the container's restart policy.
//
// swagger:model InspectRestartPolicy
type InspectRestartPolicy struct {

	// MaximumRetryCount is the maximum number of retries allowed if the
	// "on-failure" restart policy is in use. Not used if "on-failure" is
	// not set.
	MaximumRetryCount uint64 `json:"MaximumRetryCount,omitempty"`

	// Name contains the container's restart policy.
	// Allowable values are "no" or "" (take no action),
	// "on-failure" (restart on non-zero exit code, with an optional max
	// retry count), and "always" (always restart on container stop, unless
	// explicitly requested by API).
	// Note that this is NOT actually a name of any sort - the poor naming
	// is for Docker compatibility.
	Name string `json:"Name,omitempty"`
}

// Validate validates this inspect restart policy
func (m *InspectRestartPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inspect restart policy based on context it is used
func (m *InspectRestartPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InspectRestartPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectRestartPolicy) UnmarshalBinary(b []byte) error {
	var res InspectRestartPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
