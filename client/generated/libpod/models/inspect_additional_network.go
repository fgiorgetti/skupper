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
)

// InspectAdditionalNetwork InspectAdditionalNetwork holds information about non-default CNI networks the
// container has been connected to.
// As with InspectNetworkSettings, many fields are unused and maintained only
// for compatibility with Docker.
//
// swagger:model InspectAdditionalNetwork
type InspectAdditionalNetwork struct {

	// AdditionalMacAddresses is a set of additional MAC Addresses beyond
	// the first. CNI may configure more than one interface for a single
	// network, which can cause this.
	AdditionalMacAddresses []string `json:"AdditionalMACAddresses"`

	// Aliases are any network aliases the container has in this network.
	Aliases []string `json:"Aliases"`

	// DriverOpts is presently unused and maintained exclusively for
	// compatibility.
	DriverOpts map[string]string `json:"DriverOpts,omitempty"`

	// EndpointID is unused, maintained exclusively for compatibility.
	EndpointID string `json:"EndpointID,omitempty"`

	// Gateway is the IP address of the gateway this network will use.
	Gateway string `json:"Gateway,omitempty"`

	// GlobalIPv6Address is the global-scope IPv6 Address for this network.
	GlobalIPV6Address string `json:"GlobalIPv6Address,omitempty"`

	// GlobalIPv6PrefixLen is the length of the subnet mask of this network.
	GlobalIPV6PrefixLen int64 `json:"GlobalIPv6PrefixLen,omitempty"`

	// IPAMConfig is presently unused and maintained exclusively for
	// compatibility.
	IPAMConfig map[string]string `json:"IPAMConfig,omitempty"`

	// IPAddress is the IP address for this network.
	IPAddress string `json:"IPAddress,omitempty"`

	// IPPrefixLen is the length of the subnet mask of this network.
	IPPrefixLen int64 `json:"IPPrefixLen,omitempty"`

	// IPv6Gateway is the IPv6 gateway this network will use.
	IPV6Gateway string `json:"IPv6Gateway,omitempty"`

	// Links is presently unused and maintained exclusively for
	// compatibility.
	Links []string `json:"Links"`

	// MacAddress is the MAC address for the interface in this network.
	MacAddress string `json:"MacAddress,omitempty"`

	// Name of the network we're connecting to.
	NetworkID string `json:"NetworkID,omitempty"`

	// SecondaryIPAddresses is a list of extra IP Addresses that the
	// container has been assigned in this network.
	SecondaryIPAddresses []*Address `json:"SecondaryIPAddresses"`

	// SecondaryIPv6Addresses is a list of extra IPv6 Addresses that the
	// container has been assigned in this network.
	SecondaryIPV6Addresses []*Address `json:"SecondaryIPv6Addresses"`
}

// Validate validates this inspect additional network
func (m *InspectAdditionalNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecondaryIPAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryIPV6Addresses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InspectAdditionalNetwork) validateSecondaryIPAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryIPAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.SecondaryIPAddresses); i++ {
		if swag.IsZero(m.SecondaryIPAddresses[i]) { // not required
			continue
		}

		if m.SecondaryIPAddresses[i] != nil {
			if err := m.SecondaryIPAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SecondaryIPAddresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SecondaryIPAddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InspectAdditionalNetwork) validateSecondaryIPV6Addresses(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryIPV6Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.SecondaryIPV6Addresses); i++ {
		if swag.IsZero(m.SecondaryIPV6Addresses[i]) { // not required
			continue
		}

		if m.SecondaryIPV6Addresses[i] != nil {
			if err := m.SecondaryIPV6Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SecondaryIPv6Addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SecondaryIPv6Addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this inspect additional network based on the context it is used
func (m *InspectAdditionalNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSecondaryIPAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryIPV6Addresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InspectAdditionalNetwork) contextValidateSecondaryIPAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecondaryIPAddresses); i++ {

		if m.SecondaryIPAddresses[i] != nil {
			if err := m.SecondaryIPAddresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SecondaryIPAddresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SecondaryIPAddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InspectAdditionalNetwork) contextValidateSecondaryIPV6Addresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecondaryIPV6Addresses); i++ {

		if m.SecondaryIPV6Addresses[i] != nil {
			if err := m.SecondaryIPV6Addresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SecondaryIPv6Addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SecondaryIPv6Addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InspectAdditionalNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InspectAdditionalNetwork) UnmarshalBinary(b []byte) error {
	var res InspectAdditionalNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}