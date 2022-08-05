// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LinuxMemory LinuxMemory for Linux cgroup 'memory' resource management
//
// swagger:model LinuxMemory
type LinuxMemory struct {

	// DisableOOMKiller disables the OOM killer for out of memory conditions
	DisableOOMKiller bool `json:"disableOOMKiller,omitempty"`

	// Kernel memory limit (in bytes).
	Kernel int64 `json:"kernel,omitempty"`

	// Kernel memory limit for tcp (in bytes)
	KernelTCP int64 `json:"kernelTCP,omitempty"`

	// Memory limit (in bytes).
	Limit int64 `json:"limit,omitempty"`

	// Memory reservation or soft_limit (in bytes).
	Reservation int64 `json:"reservation,omitempty"`

	// Total memory limit (memory + swap).
	Swap int64 `json:"swap,omitempty"`

	// How aggressive the kernel will swap memory pages.
	Swappiness uint64 `json:"swappiness,omitempty"`

	// Enables hierarchical memory accounting
	UseHierarchy bool `json:"useHierarchy,omitempty"`
}

// Validate validates this linux memory
func (m *LinuxMemory) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this linux memory based on context it is used
func (m *LinuxMemory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinuxMemory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinuxMemory) UnmarshalBinary(b []byte) error {
	var res LinuxMemory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
