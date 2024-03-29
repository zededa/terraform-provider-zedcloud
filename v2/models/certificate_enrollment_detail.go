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

// CertificateEnrollmentDetail certificate enrollment detail
//
// swagger:model CertificateEnrollmentDetail
type CertificateEnrollmentDetail struct {

	// group certificate enrollment
	GroupCertificateEnrollment *GroupCertificateEnrollment `json:"groupCertificateEnrollment,omitempty"`

	// individual certificate enrollment
	IndividualCertificateEnrollment IndividualCertificateEnrollment `json:"individualCertificateEnrollment,omitempty"`

	// type
	Type *EnrollmentType `json:"type,omitempty"`
}

// Validate validates this certificate enrollment detail
func (m *CertificateEnrollmentDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupCertificateEnrollment(formats); err != nil {
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

func (m *CertificateEnrollmentDetail) validateGroupCertificateEnrollment(formats strfmt.Registry) error {
	if swag.IsZero(m.GroupCertificateEnrollment) { // not required
		return nil
	}

	if m.GroupCertificateEnrollment != nil {
		if err := m.GroupCertificateEnrollment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groupCertificateEnrollment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("groupCertificateEnrollment")
			}
			return err
		}
	}

	return nil
}

func (m *CertificateEnrollmentDetail) validateType(formats strfmt.Registry) error {
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

// ContextValidate validate this certificate enrollment detail based on the context it is used
func (m *CertificateEnrollmentDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroupCertificateEnrollment(ctx, formats); err != nil {
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

func (m *CertificateEnrollmentDetail) contextValidateGroupCertificateEnrollment(ctx context.Context, formats strfmt.Registry) error {

	if m.GroupCertificateEnrollment != nil {
		if err := m.GroupCertificateEnrollment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groupCertificateEnrollment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("groupCertificateEnrollment")
			}
			return err
		}
	}

	return nil
}

func (m *CertificateEnrollmentDetail) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
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
func (m *CertificateEnrollmentDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateEnrollmentDetail) UnmarshalBinary(b []byte) error {
	var res CertificateEnrollmentDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
