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

// AppACEAction App ACE Action
//
// # App ACE actions
//
// swagger:model appACEAction
type AppACEAction struct {

	// ACE drop flag
	// Required: true
	Drop bool `json:"drop"`

	// ACE limit flag
	// Required: true
	Limit bool `json:"limit"`

	// ACE limit burst
	// Required: true
	Limitburst *int64 `json:"limitburst"`

	// ACE limit rate
	// Required: true
	Limitrate *int64 `json:"limitrate"`

	// ACE limit unit
	// Required: true
	Limitunit *string `json:"limitunit"`

	// Application map params
	// Required: true
	Mapparams *AppMapParams `json:"mapparams"`

	// application port map flag
	// Required: true
	Portmap bool `json:"portmap"`
}

// Validate validates this app a c e action
func (m *AppACEAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDrop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitburst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitrate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimitunit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapparams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortmap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppACEAction) validateDrop(formats strfmt.Registry) error {

	if err := validate.Required("drop", "body", bool(m.Drop)); err != nil {
		return err
	}

	return nil
}

func (m *AppACEAction) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", bool(m.Limit)); err != nil {
		return err
	}

	return nil
}

func (m *AppACEAction) validateLimitburst(formats strfmt.Registry) error {

	if err := validate.Required("limitburst", "body", m.Limitburst); err != nil {
		return err
	}

	return nil
}

func (m *AppACEAction) validateLimitrate(formats strfmt.Registry) error {

	if err := validate.Required("limitrate", "body", m.Limitrate); err != nil {
		return err
	}

	return nil
}

func (m *AppACEAction) validateLimitunit(formats strfmt.Registry) error {

	if err := validate.Required("limitunit", "body", m.Limitunit); err != nil {
		return err
	}

	return nil
}

func (m *AppACEAction) validateMapparams(formats strfmt.Registry) error {

	if err := validate.Required("mapparams", "body", m.Mapparams); err != nil {
		return err
	}

	if m.Mapparams != nil {
		if err := m.Mapparams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mapparams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mapparams")
			}
			return err
		}
	}

	return nil
}

func (m *AppACEAction) validatePortmap(formats strfmt.Registry) error {

	if err := validate.Required("portmap", "body", bool(m.Portmap)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this app a c e action based on the context it is used
func (m *AppACEAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMapparams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppACEAction) contextValidateMapparams(ctx context.Context, formats strfmt.Registry) error {

	if m.Mapparams != nil {
		if err := m.Mapparams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mapparams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mapparams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppACEAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppACEAction) UnmarshalBinary(b []byte) error {
	var res AppACEAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
