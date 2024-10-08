// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PasswordProfile password profile
//
// swagger:model PasswordProfile
type PasswordProfile struct {

	// max length
	MaxLength int64 `json:"maxLength,omitempty"`

	// max password age
	MaxPasswordAge int64 `json:"maxPasswordAge,omitempty"`

	// min length
	MinLength int64 `json:"minLength,omitempty"`

	// min lowercase chars
	MinLowercaseChars int64 `json:"minLowercaseChars,omitempty"`

	// min numeric chars
	MinNumericChars int64 `json:"minNumericChars,omitempty"`

	// min password age
	MinPasswordAge int64 `json:"minPasswordAge,omitempty"`

	// min symbol chars
	MinSymbolChars int64 `json:"minSymbolChars,omitempty"`

	// min uppercase chars
	MinUppercaseChars int64 `json:"minUppercaseChars,omitempty"`

	// num prev password check
	NumPrevPasswordCheck int64 `json:"numPrevPasswordCheck,omitempty"`

	// password expiry notification period in seconds
	PasswordExpiryNotificationPeriodInSeconds int64 `json:"passwordExpiryNotificationPeriodInSeconds,omitempty"`
}

// Validate validates this password profile
func (m *PasswordProfile) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this password profile based on context it is used
func (m *PasswordProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PasswordProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PasswordProfile) UnmarshalBinary(b []byte) error {
	var res PasswordProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
