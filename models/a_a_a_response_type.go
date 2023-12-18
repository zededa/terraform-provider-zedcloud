// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AAAResponseType a a a response type
//
// swagger:model AAAResponseType
type AAAResponseType string

func NewAAAResponseType(value AAAResponseType) *AAAResponseType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAAResponseType.
func (m AAAResponseType) Pointer() *AAAResponseType {
	return &m
}

const (

	// AAAResponseTypeINVALID captures enum value "__INVALID__"
	AAAResponseTypeINVALID AAAResponseType = "__INVALID__"

	// AAAResponseTypeAAAResponseTypeSuccess captures enum value "AAAResponseTypeSuccess"
	AAAResponseTypeAAAResponseTypeSuccess AAAResponseType = "AAAResponseTypeSuccess"

	// AAAResponseTypeAAAResponseTypeFailure captures enum value "AAAResponseTypeFailure"
	AAAResponseTypeAAAResponseTypeFailure AAAResponseType = "AAAResponseTypeFailure"

	// AAAResponseTypeAAAResponseTypeNotify captures enum value "AAAResponseTypeNotify"
	AAAResponseTypeAAAResponseTypeNotify AAAResponseType = "AAAResponseTypeNotify"

	// AAAResponseTypeAAAResponseTypeEnterpriseSignup captures enum value "AAAResponseTypeEnterpriseSignup"
	AAAResponseTypeAAAResponseTypeEnterpriseSignup AAAResponseType = "AAAResponseTypeEnterpriseSignup"

	// AAAResponseTypeAAAResponseTypeAdminUserSignup captures enum value "AAAResponseTypeAdminUserSignup"
	AAAResponseTypeAAAResponseTypeAdminUserSignup AAAResponseType = "AAAResponseTypeAdminUserSignup"

	// AAAResponseTypeAAAResponseTypeRedirect captures enum value "AAAResponseTypeRedirect"
	AAAResponseTypeAAAResponseTypeRedirect AAAResponseType = "AAAResponseTypeRedirect"

	// AAAResponseTypeAAAResponseTypeLoginMode captures enum value "AAAResponseTypeLoginMode"
	AAAResponseTypeAAAResponseTypeLoginMode AAAResponseType = "AAAResponseTypeLoginMode"

	// AAAResponseTypeTOTPEnrolment captures enum value "TOTPEnrolment"
	AAAResponseTypeTOTPEnrolment AAAResponseType = "TOTPEnrolment"

	// AAAResponseTypeEnrolTOTP captures enum value "EnrolTOTP"
	AAAResponseTypeEnrolTOTP AAAResponseType = "EnrolTOTP"
)

// for schema
var aAAResponseTypeEnum []interface{}

func init() {
	var res []AAAResponseType
	if err := json.Unmarshal([]byte(`["__INVALID__","AAAResponseTypeSuccess","AAAResponseTypeFailure","AAAResponseTypeNotify","AAAResponseTypeEnterpriseSignup","AAAResponseTypeAdminUserSignup","AAAResponseTypeRedirect","AAAResponseTypeLoginMode","TOTPEnrolment","EnrolTOTP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAAResponseTypeEnum = append(aAAResponseTypeEnum, v)
	}
}

func (m AAAResponseType) validateAAAResponseTypeEnum(path, location string, value AAAResponseType) error {
	if err := validate.EnumCase(path, location, value, aAAResponseTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a response type
func (m AAAResponseType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAAResponseTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a response type based on context it is used
func (m AAAResponseType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
