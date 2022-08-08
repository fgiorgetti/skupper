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

// CreateContainerConfig CreateContainerConfig used when compatible endpoint creates a container
//
// swagger:model CreateContainerConfig
type CreateContainerConfig struct {

	// args escaped
	ArgsEscaped bool `json:"ArgsEscaped,omitempty"`

	// attach stderr
	AttachStderr bool `json:"AttachStderr,omitempty"`

	// attach stdin
	AttachStdin bool `json:"AttachStdin,omitempty"`

	// attach stdout
	AttachStdout bool `json:"AttachStdout,omitempty"`

	// cmd
	Cmd StrSlice `json:"Cmd,omitempty"`

	// domainname
	Domainname string `json:"Domainname,omitempty"`

	// entrypoint
	Entrypoint StrSlice `json:"Entrypoint,omitempty"`

	// env
	Env []string `json:"Env"`

	// exposed ports
	ExposedPorts PortSet `json:"ExposedPorts,omitempty"`

	// healthcheck
	Healthcheck *HealthConfig `json:"Healthcheck,omitempty"`

	// host config
	HostConfig *HostConfig `json:"HostConfig,omitempty"`

	// hostname
	Hostname string `json:"Hostname,omitempty"`

	// image
	Image string `json:"Image,omitempty"`

	// labels
	Labels map[string]string `json:"Labels,omitempty"`

	// mac address
	MacAddress string `json:"MacAddress,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// network disabled
	NetworkDisabled bool `json:"NetworkDisabled,omitempty"`

	// networking config
	NetworkingConfig *NetworkingConfig `json:"NetworkingConfig,omitempty"`

	// on build
	OnBuild []string `json:"OnBuild"`

	// open stdin
	OpenStdin bool `json:"OpenStdin,omitempty"`

	// shell
	Shell StrSlice `json:"Shell,omitempty"`

	// stdin once
	StdinOnce bool `json:"StdinOnce,omitempty"`

	// stop signal
	StopSignal string `json:"StopSignal,omitempty"`

	// stop timeout
	StopTimeout int64 `json:"StopTimeout,omitempty"`

	// tty
	Tty bool `json:"Tty,omitempty"`

	// unset env
	UnsetEnv []string `json:"UnsetEnv"`

	// unset env all
	UnsetEnvAll bool `json:"UnsetEnvAll,omitempty"`

	// user
	User string `json:"User,omitempty"`

	// volumes
	Volumes map[string]interface{} `json:"Volumes,omitempty"`

	// working dir
	WorkingDir string `json:"WorkingDir,omitempty"`
}

// Validate validates this create container config
func (m *CreateContainerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCmd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntrypoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExposedPorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthcheck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkingConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShell(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateContainerConfig) validateCmd(formats strfmt.Registry) error {
	if swag.IsZero(m.Cmd) { // not required
		return nil
	}

	if err := m.Cmd.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Cmd")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Cmd")
		}
		return err
	}

	return nil
}

func (m *CreateContainerConfig) validateEntrypoint(formats strfmt.Registry) error {
	if swag.IsZero(m.Entrypoint) { // not required
		return nil
	}

	if err := m.Entrypoint.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Entrypoint")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Entrypoint")
		}
		return err
	}

	return nil
}

func (m *CreateContainerConfig) validateExposedPorts(formats strfmt.Registry) error {
	if swag.IsZero(m.ExposedPorts) { // not required
		return nil
	}

	if m.ExposedPorts != nil {
		if err := m.ExposedPorts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ExposedPorts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ExposedPorts")
			}
			return err
		}
	}

	return nil
}

func (m *CreateContainerConfig) validateHealthcheck(formats strfmt.Registry) error {
	if swag.IsZero(m.Healthcheck) { // not required
		return nil
	}

	if m.Healthcheck != nil {
		if err := m.Healthcheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Healthcheck")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Healthcheck")
			}
			return err
		}
	}

	return nil
}

func (m *CreateContainerConfig) validateHostConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.HostConfig) { // not required
		return nil
	}

	if m.HostConfig != nil {
		if err := m.HostConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HostConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("HostConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CreateContainerConfig) validateNetworkingConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkingConfig) { // not required
		return nil
	}

	if m.NetworkingConfig != nil {
		if err := m.NetworkingConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkingConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NetworkingConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CreateContainerConfig) validateShell(formats strfmt.Registry) error {
	if swag.IsZero(m.Shell) { // not required
		return nil
	}

	if err := m.Shell.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Shell")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Shell")
		}
		return err
	}

	return nil
}

// ContextValidate validate this create container config based on the context it is used
func (m *CreateContainerConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCmd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntrypoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExposedPorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealthcheck(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHostConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkingConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShell(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateContainerConfig) contextValidateCmd(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Cmd.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Cmd")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Cmd")
		}
		return err
	}

	return nil
}

func (m *CreateContainerConfig) contextValidateEntrypoint(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Entrypoint.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Entrypoint")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Entrypoint")
		}
		return err
	}

	return nil
}

func (m *CreateContainerConfig) contextValidateExposedPorts(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ExposedPorts.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ExposedPorts")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("ExposedPorts")
		}
		return err
	}

	return nil
}

func (m *CreateContainerConfig) contextValidateHealthcheck(ctx context.Context, formats strfmt.Registry) error {

	if m.Healthcheck != nil {
		if err := m.Healthcheck.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Healthcheck")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Healthcheck")
			}
			return err
		}
	}

	return nil
}

func (m *CreateContainerConfig) contextValidateHostConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.HostConfig != nil {
		if err := m.HostConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HostConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("HostConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CreateContainerConfig) contextValidateNetworkingConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkingConfig != nil {
		if err := m.NetworkingConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkingConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NetworkingConfig")
			}
			return err
		}
	}

	return nil
}

func (m *CreateContainerConfig) contextValidateShell(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Shell.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Shell")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Shell")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateContainerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateContainerConfig) UnmarshalBinary(b []byte) error {
	var res CreateContainerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
