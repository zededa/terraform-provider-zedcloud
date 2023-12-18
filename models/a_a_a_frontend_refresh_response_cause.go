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

// AAAFrontendRefreshResponseCause a a a frontend refresh response cause
//
// swagger:model AAA_Frontend_RefreshResponseCause
type AAAFrontendRefreshResponseCause string

func NewAAAFrontendRefreshResponseCause(value AAAFrontendRefreshResponseCause) *AAAFrontendRefreshResponseCause {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAAFrontendRefreshResponseCause.
func (m AAAFrontendRefreshResponseCause) Pointer() *AAAFrontendRefreshResponseCause {
	return &m
}

const (

	// AAAFrontendRefreshResponseCauseUNSPECIFIED captures enum value "UNSPECIFIED"
	AAAFrontendRefreshResponseCauseUNSPECIFIED AAAFrontendRefreshResponseCause = "UNSPECIFIED"

	// AAAFrontendRefreshResponseCauseOK captures enum value "OK"
	AAAFrontendRefreshResponseCauseOK AAAFrontendRefreshResponseCause = "OK"

	// AAAFrontendRefreshResponseCauseUNKNOWN captures enum value "UNKNOWN"
	AAAFrontendRefreshResponseCauseUNKNOWN AAAFrontendRefreshResponseCause = "UNKNOWN"

	// AAAFrontendRefreshResponseCauseCREDENTIALS captures enum value "CREDENTIALS"
	AAAFrontendRefreshResponseCauseCREDENTIALS AAAFrontendRefreshResponseCause = "CREDENTIALS"

	// AAAFrontendRefreshResponseCauseEXPIRED captures enum value "EXPIRED"
	AAAFrontendRefreshResponseCauseEXPIRED AAAFrontendRefreshResponseCause = "EXPIRED"
)

// for schema
var aAAFrontendRefreshResponseCauseEnum []interface{}

func init() {
	var res []AAAFrontendRefreshResponseCause
	if err := json.Unmarshal([]byte(`["UNSPECIFIED","OK","UNKNOWN","CREDENTIALS","EXPIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAAFrontendRefreshResponseCauseEnum = append(aAAFrontendRefreshResponseCauseEnum, v)
	}
}

func (m AAAFrontendRefreshResponseCause) validateAAAFrontendRefreshResponseCauseEnum(path, location string, value AAAFrontendRefreshResponseCause) error {
	if err := validate.EnumCase(path, location, value, aAAFrontendRefreshResponseCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a frontend refresh response cause
func (m AAAFrontendRefreshResponseCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAAFrontendRefreshResponseCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a frontend refresh response cause based on context it is used
func (m AAAFrontendRefreshResponseCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
