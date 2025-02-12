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

// NetworkInstConfig network inst config
//
// swagger:model NetworkInstConfig
type NetworkInstConfig struct {

	// flag to indicate if this is default network
	//
	// flag to indicate if this is the default network instance for the device
	DeviceDefault bool `json:"deviceDefault,omitempty"`

	// List of Static DNS entries
	DNSList []*StaticDNSList `json:"dnsList"`

	// Dhcp Server Configuration
	IP *DhcpServerConfig `json:"ip,omitempty"`

	// Kind of Network Instance ( Local, Switch etc )
	Kind *NetworkInstanceKind `json:"kind,omitempty"`

	// Service specific Config
	Opaque *NetInstOpaqueConfig `json:"opaque,omitempty"`

	// name of port mapping in the model
	//
	// name of port mapping in the model
	Port string `json:"port,omitempty"`

	// Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.
	PortTags map[string]string `json:"portTags,omitempty"`

	// flag to indicate automatically propagate connected routes
	//
	// Automatically propagate connected routes
	PropagateConnectedRoutes bool `json:"propagateConnectedRoutes,omitempty"`

	// list of static IP routes
	//
	// List of Static IP routes
	StaticRoutes []*StaticIPRoute `json:"staticRoutes"`

	// Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.
	Tags map[string]string `json:"tags,omitempty"`

	// Type of DHCP for this Network Instance
	Type *NetworkInstanceDhcpType `json:"type,omitempty"`
}

// Validate validates this network inst config
func (m *NetworkInstConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpaque(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInstConfig) validateDNSList(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSList) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSList); i++ {
		if swag.IsZero(m.DNSList[i]) { // not required
			continue
		}

		if m.DNSList[i] != nil {
			if err := m.DNSList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInstConfig) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if m.IP != nil {
		if err := m.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstConfig) validateKind(formats strfmt.Registry) error {
	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if m.Kind != nil {
		if err := m.Kind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstConfig) validateOpaque(formats strfmt.Registry) error {
	if swag.IsZero(m.Opaque) { // not required
		return nil
	}

	if m.Opaque != nil {
		if err := m.Opaque.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opaque")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opaque")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstConfig) validateStaticRoutes(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticRoutes) { // not required
		return nil
	}

	for i := 0; i < len(m.StaticRoutes); i++ {
		if swag.IsZero(m.StaticRoutes[i]) { // not required
			continue
		}

		if m.StaticRoutes[i] != nil {
			if err := m.StaticRoutes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("staticRoutes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("staticRoutes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInstConfig) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network inst config based on the context it is used
func (m *NetworkInstConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpaque(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticRoutes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInstConfig) contextValidateDNSList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSList); i++ {

		if m.DNSList[i] != nil {
			if err := m.DNSList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInstConfig) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if m.IP != nil {
		if err := m.IP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstConfig) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if m.Kind != nil {
		if err := m.Kind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstConfig) contextValidateOpaque(ctx context.Context, formats strfmt.Registry) error {

	if m.Opaque != nil {
		if err := m.Opaque.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opaque")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opaque")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstConfig) contextValidateStaticRoutes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StaticRoutes); i++ {

		if m.StaticRoutes[i] != nil {
			if err := m.StaticRoutes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("staticRoutes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("staticRoutes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInstConfig) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInstConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInstConfig) UnmarshalBinary(b []byte) error {
	var res NetworkInstConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
