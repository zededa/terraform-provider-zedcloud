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

// NetInstOpaqueConfig Network Instance Opaque config. This is service specific configuration.
//
// swagger:model NetInstOpaqueConfig
type NetInstOpaqueConfig struct {

	// lisp - Deprecated
	//
	// Deprecated - Lisp config
	Lisp *LispConfig `json:"lisp,omitempty"`

	// base64 encoded string of opaque config
	Oconfig string `json:"oconfig,omitempty"`

	// type of Opaque config
	Type *OpaqueConfigType `json:"type,omitempty"`
}

// Validate validates this net inst opaque config
func (m *NetInstOpaqueConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLisp(formats); err != nil {
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

func (m *NetInstOpaqueConfig) validateLisp(formats strfmt.Registry) error {
	if swag.IsZero(m.Lisp) { // not required
		return nil
	}

	if m.Lisp != nil {
		if err := m.Lisp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lisp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lisp")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstOpaqueConfig) validateType(formats strfmt.Registry) error {
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

// ContextValidate validate this net inst opaque config based on the context it is used
func (m *NetInstOpaqueConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLisp(ctx, formats); err != nil {
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

func (m *NetInstOpaqueConfig) contextValidateLisp(ctx context.Context, formats strfmt.Registry) error {

	if m.Lisp != nil {
		if err := m.Lisp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lisp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lisp")
			}
			return err
		}
	}

	return nil
}

func (m *NetInstOpaqueConfig) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NetInstOpaqueConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetInstOpaqueConfig) UnmarshalBinary(b []byte) error {
	var res NetInstOpaqueConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
