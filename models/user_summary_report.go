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

// UserSummaryReport user summary report
//
// swagger:model UserSummaryReport
type UserSummaryReport struct {

	// active users count
	ActiveUsersCount int64 `json:"activeUsersCount,omitempty"`

	// last loggedin user info
	LastLoggedinUserInfo *LastLoggedinUserInfo `json:"lastLoggedinUserInfo,omitempty"`

	// users
	Users int64 `json:"users,omitempty"`
}

// Validate validates this user summary report
func (m *UserSummaryReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastLoggedinUserInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSummaryReport) validateLastLoggedinUserInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.LastLoggedinUserInfo) { // not required
		return nil
	}

	if m.LastLoggedinUserInfo != nil {
		if err := m.LastLoggedinUserInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastLoggedinUserInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastLoggedinUserInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user summary report based on the context it is used
func (m *UserSummaryReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastLoggedinUserInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSummaryReport) contextValidateLastLoggedinUserInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.LastLoggedinUserInfo != nil {

		if swag.IsZero(m.LastLoggedinUserInfo) { // not required
			return nil
		}

		if err := m.LastLoggedinUserInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastLoggedinUserInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastLoggedinUserInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserSummaryReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSummaryReport) UnmarshalBinary(b []byte) error {
	var res UserSummaryReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
