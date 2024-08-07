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

// CellularStatus CellularStatus contains status information for a single cellular network
//
// # CellularStatus contains status information for a single cellular network
//
// swagger:model CellularStatus
type CellularStatus struct {

	// Modem detail
	Modem *ModemDetail `json:"modem,omitempty"`

	// List of network providers
	Providers []*NetworkProvider `json:"providers"`

	// List of SIM Cards
	SimCards []*SIMCard `json:"simCards"`
}

// Validate validates this cellular status
func (m *CellularStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSimCards(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CellularStatus) validateModem(formats strfmt.Registry) error {
	if swag.IsZero(m.Modem) { // not required
		return nil
	}

	if m.Modem != nil {
		if err := m.Modem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modem")
			}
			return err
		}
	}

	return nil
}

func (m *CellularStatus) validateProviders(formats strfmt.Registry) error {
	if swag.IsZero(m.Providers) { // not required
		return nil
	}

	for i := 0; i < len(m.Providers); i++ {
		if swag.IsZero(m.Providers[i]) { // not required
			continue
		}

		if m.Providers[i] != nil {
			if err := m.Providers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("providers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CellularStatus) validateSimCards(formats strfmt.Registry) error {
	if swag.IsZero(m.SimCards) { // not required
		return nil
	}

	for i := 0; i < len(m.SimCards); i++ {
		if swag.IsZero(m.SimCards[i]) { // not required
			continue
		}

		if m.SimCards[i] != nil {
			if err := m.SimCards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("simCards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("simCards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cellular status based on the context it is used
func (m *CellularStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProviders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSimCards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CellularStatus) contextValidateModem(ctx context.Context, formats strfmt.Registry) error {

	if m.Modem != nil {

		if swag.IsZero(m.Modem) { // not required
			return nil
		}

		if err := m.Modem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modem")
			}
			return err
		}
	}

	return nil
}

func (m *CellularStatus) contextValidateProviders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Providers); i++ {

		if m.Providers[i] != nil {

			if swag.IsZero(m.Providers[i]) { // not required
				return nil
			}

			if err := m.Providers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("providers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CellularStatus) contextValidateSimCards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SimCards); i++ {

		if m.SimCards[i] != nil {

			if swag.IsZero(m.SimCards[i]) { // not required
				return nil
			}

			if err := m.SimCards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("simCards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("simCards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CellularStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CellularStatus) UnmarshalBinary(b []byte) error {
	var res CellularStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
