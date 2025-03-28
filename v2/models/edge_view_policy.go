// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EdgeViewPolicy Edgeview policy body detail
//
// # Policy for device edgeview operation
//
// swagger:model EdgeViewPolicy
type EdgeViewPolicy struct {

	// Allow inherit instance to change access policy
	AccessAllowChange bool `json:"accessAllowChange,omitempty"`

	// Allow device to enable Edgeview in this project
	// Required: true
	EdgeviewAllow *bool `json:"edgeviewAllow"`

	// Edgeview configuration and policies
	Edgeviewcfg *EdgeviewCfg `json:"edgeviewcfg,omitempty"`

	// Maximum seconds allowed for Edgeview session
	MaxExpireSec int64 `json:"maxExpireSec,omitempty"`

	// Maximum instances allowed for Edgeview
	MaxInst int64 `json:"maxInst,omitempty"`

	// all the required metadata for a policy like id, name, different types of tags
	MetaData *PolicyCommon `json:"metaData,omitempty"`
}

// Validate validates this edge view policy
func (m *EdgeViewPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdgeviewAllow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeviewcfg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeViewPolicy) validateEdgeviewAllow(formats strfmt.Registry) error {

	if err := validate.Required("edgeviewAllow", "body", m.EdgeviewAllow); err != nil {
		return err
	}

	return nil
}

func (m *EdgeViewPolicy) validateEdgeviewcfg(formats strfmt.Registry) error {
	if swag.IsZero(m.Edgeviewcfg) { // not required
		return nil
	}

	if m.Edgeviewcfg != nil {
		if err := m.Edgeviewcfg.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeviewcfg")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeviewcfg")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeViewPolicy) validateMetaData(formats strfmt.Registry) error {
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

// ContextValidate validate this edge view policy based on the context it is used
func (m *EdgeViewPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdgeviewcfg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetaData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeViewPolicy) contextValidateEdgeviewcfg(ctx context.Context, formats strfmt.Registry) error {

	if m.Edgeviewcfg != nil {

		if swag.IsZero(m.Edgeviewcfg) { // not required
			return nil
		}

		if err := m.Edgeviewcfg.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeviewcfg")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeviewcfg")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeViewPolicy) contextValidateMetaData(ctx context.Context, formats strfmt.Registry) error {

	if m.MetaData != nil {

		if swag.IsZero(m.MetaData) { // not required
			return nil
		}

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

// MarshalBinary interface implementation
func (m *EdgeViewPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeViewPolicy) UnmarshalBinary(b []byte) error {
	var res EdgeViewPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
