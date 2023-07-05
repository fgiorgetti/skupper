// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Network Network describes the Network attributes.
//
// swagger:model Network
type Network struct {

	// Created contains the timestamp when this network was created.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// DNSEnabled is whether name resolution is active for container on
	// this Network.
	DNSEnabled bool `json:"dns_enabled,omitempty"`

	// Driver for this Network, e.g. bridge, macvlan...
	Driver string `json:"driver,omitempty"`

	// ID of the Network.
	ID string `json:"id,omitempty"`

	// IPAMOptions contains options used for the ip assignment.
	IPAMOptions map[string]string `json:"ipam_options,omitempty"`

	// IPv6Enabled if set to true an ipv6 subnet should be created for this net.
	IPV6Enabled bool `json:"ipv6_enabled,omitempty"`

	// Internal is whether the Network should not have external routes
	// to public or other Networks.
	Internal bool `json:"internal,omitempty"`

	// Labels is a set of key-value labels that have been applied to the
	// Network.
	Labels map[string]string `json:"labels,omitempty"`

	// Name of the Network.
	Name string `json:"name,omitempty"`

	// NetworkInterface is the network interface name on the host.
	NetworkInterface string `json:"network_interface,omitempty"`

	// Options is a set of key-value options that have been applied to
	// the Network.
	Options map[string]string `json:"options,omitempty"`

	// Subnets to use for this network.
	Subnets []*Subnet `json:"subnets"`
}

// Validate validates this network
func (m *Network) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Network) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateSubnets(formats strfmt.Registry) error {
	if swag.IsZero(m.Subnets) { // not required
		return nil
	}

	for i := 0; i < len(m.Subnets); i++ {
		if swag.IsZero(m.Subnets[i]) { // not required
			continue
		}

		if m.Subnets[i] != nil {
			if err := m.Subnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network based on the context it is used
func (m *Network) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubnets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Network) contextValidateSubnets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Subnets); i++ {

		if m.Subnets[i] != nil {
			if err := m.Subnets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Network) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Network) UnmarshalBinary(b []byte) error {
	var res Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
