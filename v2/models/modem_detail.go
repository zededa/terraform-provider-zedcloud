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

// ModemDetail ModemDetail describes a cellular Modem
//
// # ModemDetail describes a cellular Modem
//
// swagger:model ModemDetail
type ModemDetail struct {

	// Modem control protocol
	ControlProtocol *ZCellularControlProtocol `json:"controlProtocol,omitempty"`

	// Modem firmware version
	FirmwareVersion string `json:"firmwareVersion,omitempty"`

	// Modem IMEI
	Imei string `json:"imei,omitempty"`

	// Modem manufacturer
	Manufacturer string `json:"manufacturer,omitempty"`

	// Modem model
	Model string `json:"model,omitempty"`

	// Modem operating state
	OperatingState *ZCellularOperatingState `json:"operatingState,omitempty"`
}

// Validate validates this modem detail
func (m *ModemDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateControlProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModemDetail) validateControlProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.ControlProtocol) { // not required
		return nil
	}

	if m.ControlProtocol != nil {
		if err := m.ControlProtocol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controlProtocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controlProtocol")
			}
			return err
		}
	}

	return nil
}

func (m *ModemDetail) validateOperatingState(formats strfmt.Registry) error {
	if swag.IsZero(m.OperatingState) { // not required
		return nil
	}

	if m.OperatingState != nil {
		if err := m.OperatingState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operatingState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operatingState")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this modem detail based on the context it is used
func (m *ModemDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateControlProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperatingState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModemDetail) contextValidateControlProtocol(ctx context.Context, formats strfmt.Registry) error {

	if m.ControlProtocol != nil {

		if swag.IsZero(m.ControlProtocol) { // not required
			return nil
		}

		if err := m.ControlProtocol.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("controlProtocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("controlProtocol")
			}
			return err
		}
	}

	return nil
}

func (m *ModemDetail) contextValidateOperatingState(ctx context.Context, formats strfmt.Registry) error {

	if m.OperatingState != nil {

		if swag.IsZero(m.OperatingState) { // not required
			return nil
		}

		if err := m.OperatingState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operatingState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operatingState")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModemDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModemDetail) UnmarshalBinary(b []byte) error {
	var res ModemDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
