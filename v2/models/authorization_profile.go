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

// AuthorizationProfile Authorization profile detail
//
// # Authorization profile  meta data
//
// swagger:model AuthorizationProfile
type AuthorizationProfile struct {

	// Mark this profile as active. Only one profile can be active in a given enterprise
	Active bool `json:"active,omitempty"`

	// Default Role ID to associate with the profile
	// Required: true
	// Pattern: [0-9A-Za-z_=-]{28}
	DefaultRoleID *string `json:"defaultRoleId"`

	// Detailed description of the profile
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// Do not automatically create new users if this is set
	DisableAutoUserCreate bool `json:"disableAutoUserCreate,omitempty"`

	// Parent enterprise ID of the authorization profile
	// Pattern: [0-9A-Za-z_=-]{28}
	EnterpriseID string `json:"enterpriseId,omitempty"`

	// Unique system defined profile ID
	// Read Only: true
	// Pattern: [0-9A-Za-z_=-]{28}
	ID string `json:"id,omitempty"`

	// User defined name of the profile. Profile name is unique within an enterprise. Name can't be changed once created
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// Oauth profile configuration details
	OauthProfile *OAUTHProfile `json:"oauthProfile,omitempty"`

	// password profile
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`

	// Authorization profile type
	ProfileType *AuthProfileType `json:"profileType,omitempty"`

	// system defined info
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// test only
	TestOnly bool `json:"testOnly,omitempty"`

	// User defined title for the profile. Title can be changed anytime
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title *string `json:"title"`

	// Type of the profile
	Type *AuthType `json:"type,omitempty"`
}

// Validate validates this authorization profile
func (m *AuthorizationProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultRoleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnterpriseID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOauthProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizationProfile) validateDefaultRoleID(formats strfmt.Registry) error {

	if err := validate.Required("defaultRoleId", "body", m.DefaultRoleID); err != nil {
		return err
	}

	if err := validate.Pattern("defaultRoleId", "body", *m.DefaultRoleID, `[0-9A-Za-z_=-]{28}`); err != nil {
		return err
	}

	return nil
}

func (m *AuthorizationProfile) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *AuthorizationProfile) validateEnterpriseID(formats strfmt.Registry) error {
	if swag.IsZero(m.EnterpriseID) { // not required
		return nil
	}

	if err := validate.Pattern("enterpriseId", "body", m.EnterpriseID, `[0-9A-Za-z_=-]{28}`); err != nil {
		return err
	}

	return nil
}

func (m *AuthorizationProfile) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[0-9A-Za-z_=-]{28}`); err != nil {
		return err
	}

	return nil
}

func (m *AuthorizationProfile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *AuthorizationProfile) validateOauthProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.OauthProfile) { // not required
		return nil
	}

	if m.OauthProfile != nil {
		if err := m.OauthProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauthProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauthProfile")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationProfile) validatePasswordProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.PasswordProfile) { // not required
		return nil
	}

	if m.PasswordProfile != nil {
		if err := m.PasswordProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("passwordProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("passwordProfile")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationProfile) validateProfileType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProfileType) { // not required
		return nil
	}

	if m.ProfileType != nil {
		if err := m.ProfileType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profileType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profileType")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationProfile) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationProfile) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", *m.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

func (m *AuthorizationProfile) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this authorization profile based on the context it is used
func (m *AuthorizationProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOauthProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePasswordProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProfileType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizationProfile) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AuthorizationProfile) contextValidateOauthProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.OauthProfile != nil {

		if swag.IsZero(m.OauthProfile) { // not required
			return nil
		}

		if err := m.OauthProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauthProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauthProfile")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationProfile) contextValidatePasswordProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.PasswordProfile != nil {

		if swag.IsZero(m.PasswordProfile) { // not required
			return nil
		}

		if err := m.PasswordProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("passwordProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("passwordProfile")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationProfile) contextValidateProfileType(ctx context.Context, formats strfmt.Registry) error {

	if m.ProfileType != nil {

		if swag.IsZero(m.ProfileType) { // not required
			return nil
		}

		if err := m.ProfileType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profileType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("profileType")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationProfile) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {

		if swag.IsZero(m.Revision) { // not required
			return nil
		}

		if err := m.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationProfile) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthorizationProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizationProfile) UnmarshalBinary(b []byte) error {
	var res AuthorizationProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
