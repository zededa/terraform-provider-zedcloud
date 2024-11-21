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

// OAUTHProfile o a u t h profile
//
// swagger:model OAUTHProfile
type OAUTHProfile struct {

	// OIDC endpoint for oauth validation
	OIDCEndPoint string `json:"OIDCEndPoint,omitempty"`

	// pass additional url parameters during the exchange and authorization process
	AdditionalParameters string `json:"additionalParameters,omitempty"`

	// OAUTH client ID
	ClientID string `json:"clientID,omitempty"`

	// OAUTH client secret
	ClientSecret string `json:"clientSecret,omitempty"`

	// crypto key
	CryptoKey string `json:"cryptoKey,omitempty"`

	// encrypted secrets
	EncryptedSecrets map[string]string `json:"encryptedSecrets,omitempty"`

	// id for Vmware IDP
	IdpID string `json:"idpId,omitempty"`

	// Config for JWT based authentication, jwks_uri is derived from OIDC Well Known Endpoints
	JwtAuthProfile *OAUTHProfileJWTAuthProfile `json:"jwtAuthProfile,omitempty"`

	// OIDC scope to fetch application role
	RoleScope string `json:"roleScope,omitempty"`
}

// Validate validates this o a u t h profile
func (m *OAUTHProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJwtAuthProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OAUTHProfile) validateJwtAuthProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.JwtAuthProfile) { // not required
		return nil
	}

	if m.JwtAuthProfile != nil {
		if err := m.JwtAuthProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jwtAuthProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jwtAuthProfile")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this o a u t h profile based on the context it is used
func (m *OAUTHProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJwtAuthProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OAUTHProfile) contextValidateJwtAuthProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.JwtAuthProfile != nil {

		if swag.IsZero(m.JwtAuthProfile) { // not required
			return nil
		}

		if err := m.JwtAuthProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jwtAuthProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jwtAuthProfile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OAUTHProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAUTHProfile) UnmarshalBinary(b []byte) error {
	var res OAUTHProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}