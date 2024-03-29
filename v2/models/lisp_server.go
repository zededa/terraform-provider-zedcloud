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

// LispServer LispServer payload detail
//
// # LispServer request paylod
//
// swagger:model LispServer
type LispServer struct {

	// lisp credential
	// Required: true
	Credential *string `json:"credential"`

	// name/IP
	// Required: true
	NameOrIP *string `json:"nameOrIp"`
}

// Validate validates this lisp server
func (m *LispServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameOrIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LispServer) validateCredential(formats strfmt.Registry) error {

	if err := validate.Required("credential", "body", m.Credential); err != nil {
		return err
	}

	return nil
}

func (m *LispServer) validateNameOrIP(formats strfmt.Registry) error {

	if err := validate.Required("nameOrIp", "body", m.NameOrIP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this lisp server based on context it is used
func (m *LispServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LispServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LispServer) UnmarshalBinary(b []byte) error {
	var res LispServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
