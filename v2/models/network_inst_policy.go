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

// NetworkInstPolicy network inst policy
//
// swagger:model NetworkInstPolicy
type NetworkInstPolicy struct {

	// all the required metadata for a policy like id, name, different types of tags
	MetaData *PolicyCommon `json:"metaData,omitempty"`

	// network instance config details
	NetInstConfig *NetworkInstConfig `json:"netInstConfig,omitempty"`
}

// Validate validates this network inst policy
func (m *NetworkInstPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetaData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetInstConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInstPolicy) validateMetaData(formats strfmt.Registry) error {
	if swag.IsZero(m.MetaData) { // not required
		return nil
	}

	if m.MetaData != nil {
		if err := m.MetaData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metaData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metaData")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstPolicy) validateNetInstConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NetInstConfig) { // not required
		return nil
	}

	if m.NetInstConfig != nil {
		if err := m.NetInstConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netInstConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netInstConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network inst policy based on the context it is used
func (m *NetworkInstPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetaData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetInstConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInstPolicy) contextValidateMetaData(ctx context.Context, formats strfmt.Registry) error {

	if m.MetaData != nil {
		if err := m.MetaData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metaData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metaData")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstPolicy) contextValidateNetInstConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NetInstConfig != nil {
		if err := m.NetInstConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netInstConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netInstConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInstPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInstPolicy) UnmarshalBinary(b []byte) error {
	var res NetworkInstPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
