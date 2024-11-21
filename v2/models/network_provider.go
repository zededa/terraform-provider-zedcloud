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

// NetworkProvider NetworkProvider describes a cellular Network Provider
//
// # NetworkProvider describes a cellular Network Provider
//
// swagger:model NetworkProvider
type NetworkProvider struct {

	// Time when was the connection established
	// Format: date-time
	ConnectedAt strfmt.DateTime `json:"connectedAt,omitempty"`

	// Human-friendly name
	Name string `json:"name,omitempty"`

	// PLMN code
	Plmn string `json:"plmn,omitempty"`

	// Radio Access Technology
	Rat *RadioAccessTechnology `json:"rat,omitempty"`

	// Is Roaming ON
	Roaming bool `json:"roaming,omitempty"`
}

// Validate validates this network provider
func (m *NetworkProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkProvider) validateConnectedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("connectedAt", "body", "date-time", m.ConnectedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NetworkProvider) validateRat(formats strfmt.Registry) error {
	if swag.IsZero(m.Rat) { // not required
		return nil
	}

	if m.Rat != nil {
		if err := m.Rat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rat")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network provider based on the context it is used
func (m *NetworkProvider) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkProvider) contextValidateRat(ctx context.Context, formats strfmt.Registry) error {

	if m.Rat != nil {

		if swag.IsZero(m.Rat) { // not required
			return nil
		}

		if err := m.Rat.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rat")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkProvider) UnmarshalBinary(b []byte) error {
	var res NetworkProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}