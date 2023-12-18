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

// AAAFrontendRefreshResponse a a a frontend refresh response
//
// swagger:model AAA_Frontend_RefreshResponse
type AAAFrontendRefreshResponse struct {

	// cause
	Cause *AAAFrontendRefreshResponseCause `json:"cause,omitempty"`

	// token
	Token *Token64 `json:"token,omitempty"`

	// BEGIN: Only valid when Cause == OK
	UserID string `json:"userId,omitempty"`
}

// Validate validates this a a a frontend refresh response
func (m *AAAFrontendRefreshResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAAFrontendRefreshResponse) validateCause(formats strfmt.Registry) error {
	if swag.IsZero(m.Cause) { // not required
		return nil
	}

	if m.Cause != nil {
		if err := m.Cause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

func (m *AAAFrontendRefreshResponse) validateToken(formats strfmt.Registry) error {
	if swag.IsZero(m.Token) { // not required
		return nil
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a a a frontend refresh response based on the context it is used
func (m *AAAFrontendRefreshResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCause(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAAFrontendRefreshResponse) contextValidateCause(ctx context.Context, formats strfmt.Registry) error {

	if m.Cause != nil {

		if swag.IsZero(m.Cause) { // not required
			return nil
		}

		if err := m.Cause.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

func (m *AAAFrontendRefreshResponse) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if m.Token != nil {

		if swag.IsZero(m.Token) { // not required
			return nil
		}

		if err := m.Token.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AAAFrontendRefreshResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAAFrontendRefreshResponse) UnmarshalBinary(b []byte) error {
	var res AAAFrontendRefreshResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
