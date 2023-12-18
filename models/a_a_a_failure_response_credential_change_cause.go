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

// AAAFailureResponseCredentialChangeCause a a a failure response credential change cause
//
// swagger:model AAAFailureResponseCredentialChangeCause
type AAAFailureResponseCredentialChangeCause string

func NewAAAFailureResponseCredentialChangeCause(value AAAFailureResponseCredentialChangeCause) *AAAFailureResponseCredentialChangeCause {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAAFailureResponseCredentialChangeCause.
func (m AAAFailureResponseCredentialChangeCause) Pointer() *AAAFailureResponseCredentialChangeCause {
	return &m
}

const (

	// AAAFailureResponseCredentialChangeCauseINVALID captures enum value "__INVALID__"
	AAAFailureResponseCredentialChangeCauseINVALID AAAFailureResponseCredentialChangeCause = "__INVALID__"

	// AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCauseUnknown captures enum value "AAAFailureResponseCredentialChangeCauseUnknown"
	AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCauseUnknown AAAFailureResponseCredentialChangeCause = "AAAFailureResponseCredentialChangeCauseUnknown"

	// AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCauseAuthenticationFailure captures enum value "AAAFailureResponseCredentialChangeCauseAuthenticationFailure"
	AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCauseAuthenticationFailure AAAFailureResponseCredentialChangeCause = "AAAFailureResponseCredentialChangeCauseAuthenticationFailure"

	// AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCausePermissionDenied captures enum value "AAAFailureResponseCredentialChangeCausePermissionDenied"
	AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCausePermissionDenied AAAFailureResponseCredentialChangeCause = "AAAFailureResponseCredentialChangeCausePermissionDenied"

	// AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCauseInternalServerError captures enum value "AAAFailureResponseCredentialChangeCauseInternalServerError"
	AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCauseInternalServerError AAAFailureResponseCredentialChangeCause = "AAAFailureResponseCredentialChangeCauseInternalServerError"

	// AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCauseAlreadyExists captures enum value "AAAFailureResponseCredentialChangeCauseAlreadyExists"
	AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCauseAlreadyExists AAAFailureResponseCredentialChangeCause = "AAAFailureResponseCredentialChangeCauseAlreadyExists"

	// AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCausePasswordTooSimple captures enum value "AAAFailureResponseCredentialChangeCausePasswordTooSimple"
	AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCausePasswordTooSimple AAAFailureResponseCredentialChangeCause = "AAAFailureResponseCredentialChangeCausePasswordTooSimple"

	// AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCauseNotFound captures enum value "AAAFailureResponseCredentialChangeCauseNotFound"
	AAAFailureResponseCredentialChangeCauseAAAFailureResponseCredentialChangeCauseNotFound AAAFailureResponseCredentialChangeCause = "AAAFailureResponseCredentialChangeCauseNotFound"
)

// for schema
var aAAFailureResponseCredentialChangeCauseEnum []interface{}

func init() {
	var res []AAAFailureResponseCredentialChangeCause
	if err := json.Unmarshal([]byte(`["__INVALID__","AAAFailureResponseCredentialChangeCauseUnknown","AAAFailureResponseCredentialChangeCauseAuthenticationFailure","AAAFailureResponseCredentialChangeCausePermissionDenied","AAAFailureResponseCredentialChangeCauseInternalServerError","AAAFailureResponseCredentialChangeCauseAlreadyExists","AAAFailureResponseCredentialChangeCausePasswordTooSimple","AAAFailureResponseCredentialChangeCauseNotFound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAAFailureResponseCredentialChangeCauseEnum = append(aAAFailureResponseCredentialChangeCauseEnum, v)
	}
}

func (m AAAFailureResponseCredentialChangeCause) validateAAAFailureResponseCredentialChangeCauseEnum(path, location string, value AAAFailureResponseCredentialChangeCause) error {
	if err := validate.EnumCase(path, location, value, aAAFailureResponseCredentialChangeCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a failure response credential change cause
func (m AAAFailureResponseCredentialChangeCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAAFailureResponseCredentialChangeCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a failure response credential change cause based on context it is used
func (m AAAFailureResponseCredentialChangeCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
