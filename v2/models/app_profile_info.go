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

// AppProfileInfo app profile info
//
// swagger:model AppProfileInfo
type AppProfileInfo struct {

	// Unique ID of the asset group.
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	AppProfileID string `json:"appProfileId,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this app profile info
func (m *AppProfileInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppProfileID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppProfileInfo) validateAppProfileID(formats strfmt.Registry) error {
	if swag.IsZero(m.AppProfileID) { // not required
		return nil
	}

	if err := validate.Pattern("appProfileId", "body", m.AppProfileID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this app profile info based on the context it is used
func (m *AppProfileInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppProfileID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppProfileInfo) contextValidateAppProfileID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "appProfileId", "body", string(m.AppProfileID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppProfileInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppProfileInfo) UnmarshalBinary(b []byte) error {
	var res AppProfileInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
