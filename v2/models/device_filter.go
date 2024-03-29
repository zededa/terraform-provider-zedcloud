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

// DeviceFilter DeviceFilter payload detail
//
// # DeviceFilter request paylod
//
// swagger:model DeviceFilter
type DeviceFilter struct {

	// admin state of the device
	AdminState *AdminState `json:"adminState,omitempty"`

	// name pattern
	// Required: true
	NamePattern *string `json:"namePattern"`

	// project
	// Required: true
	Project *string `json:"project"`

	// project name pattern
	ProjectNamePattern string `json:"projectNamePattern,omitempty"`
}

// Validate validates this device filter
func (m *DeviceFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamePattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceFilter) validateAdminState(formats strfmt.Registry) error {
	if swag.IsZero(m.AdminState) { // not required
		return nil
	}

	if m.AdminState != nil {
		if err := m.AdminState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminState")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceFilter) validateNamePattern(formats strfmt.Registry) error {

	if err := validate.Required("namePattern", "body", m.NamePattern); err != nil {
		return err
	}

	return nil
}

func (m *DeviceFilter) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this device filter based on the context it is used
func (m *DeviceFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdminState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceFilter) contextValidateAdminState(ctx context.Context, formats strfmt.Registry) error {

	if m.AdminState != nil {
		if err := m.AdminState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adminState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adminState")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceFilter) UnmarshalBinary(b []byte) error {
	var res DeviceFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
