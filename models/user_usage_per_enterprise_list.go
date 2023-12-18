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

// UserUsagePerEnterpriseList user usage per enterprise list
//
// swagger:model UserUsagePerEnterpriseList
type UserUsagePerEnterpriseList struct {

	// user usage per enterprise list
	UserUsagePerEntp []*UserUsagePerEnterprise `json:"userUsagePerEntp"`
}

// Validate validates this user usage per enterprise list
func (m *UserUsagePerEnterpriseList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserUsagePerEntp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserUsagePerEnterpriseList) validateUserUsagePerEntp(formats strfmt.Registry) error {
	if swag.IsZero(m.UserUsagePerEntp) { // not required
		return nil
	}

	for i := 0; i < len(m.UserUsagePerEntp); i++ {
		if swag.IsZero(m.UserUsagePerEntp[i]) { // not required
			continue
		}

		if m.UserUsagePerEntp[i] != nil {
			if err := m.UserUsagePerEntp[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userUsagePerEntp" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userUsagePerEntp" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this user usage per enterprise list based on the context it is used
func (m *UserUsagePerEnterpriseList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserUsagePerEntp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserUsagePerEnterpriseList) contextValidateUserUsagePerEntp(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserUsagePerEntp); i++ {

		if m.UserUsagePerEntp[i] != nil {

			if swag.IsZero(m.UserUsagePerEntp[i]) { // not required
				return nil
			}

			if err := m.UserUsagePerEntp[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userUsagePerEntp" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userUsagePerEntp" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserUsagePerEnterpriseList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUsagePerEnterpriseList) UnmarshalBinary(b []byte) error {
	var res UserUsagePerEnterpriseList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
