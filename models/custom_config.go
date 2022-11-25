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

// CustomConfig custom config
//
// swagger:model CustomConfig
type CustomConfig struct {

	// add
	Add bool `json:"add,omitempty"`

	// allow storage resize
	AllowStorageResize bool `json:"allowStorageResize,omitempty"`

	// field delimiter
	FieldDelimiter string `json:"fieldDelimiter,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// override
	Override bool `json:"override,omitempty"`

	// template
	Template string `json:"template,omitempty"`

	// variable groups
	VariableGroups []*CustomConfigVariableGroup `json:"variableGroups"`
}

// Validate validates this custom config
func (m *CustomConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVariableGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomConfig) validateVariableGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.VariableGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.VariableGroups); i++ {
		if swag.IsZero(m.VariableGroups[i]) { // not required
			continue
		}

		if m.VariableGroups[i] != nil {
			if err := m.VariableGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variableGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variableGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this custom config based on the context it is used
func (m *CustomConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVariableGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomConfig) contextValidateVariableGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VariableGroups); i++ {

		if m.VariableGroups[i] != nil {
			if err := m.VariableGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("variableGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("variableGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomConfig) UnmarshalBinary(b []byte) error {
	var res CustomConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
