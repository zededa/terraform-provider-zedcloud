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

// DevDataSecAtRest dev data sec at rest
//
// swagger:model DevDataSecAtRest
type DevDataSecAtRest struct {

	// err info
	ErrInfo *DeviceError `json:"errInfo,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// status
	Status *DeviceDataSecurityAtRestStatus `json:"status,omitempty"`
}

// Validate validates this dev data sec at rest
func (m *DevDataSecAtRest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevDataSecAtRest) validateErrInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrInfo) { // not required
		return nil
	}

	if m.ErrInfo != nil {
		if err := m.ErrInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errInfo")
			}
			return err
		}
	}

	return nil
}

func (m *DevDataSecAtRest) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this dev data sec at rest based on the context it is used
func (m *DevDataSecAtRest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevDataSecAtRest) contextValidateErrInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrInfo != nil {
		if err := m.ErrInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errInfo")
			}
			return err
		}
	}

	return nil
}

func (m *DevDataSecAtRest) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DevDataSecAtRest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevDataSecAtRest) UnmarshalBinary(b []byte) error {
	var res DevDataSecAtRest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
