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

// UserState General well being of the user
//
// swagger:model UserState
type UserState string

func NewUserState(value UserState) *UserState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated UserState.
func (m UserState) Pointer() *UserState {
	return &m
}

const (

	// UserStateUSERSTATEUNSPECIFIED captures enum value "USER_STATE_UNSPECIFIED"
	UserStateUSERSTATEUNSPECIFIED UserState = "USER_STATE_UNSPECIFIED"

	// UserStateUSERSTATECREATED captures enum value "USER_STATE_CREATED"
	UserStateUSERSTATECREATED UserState = "USER_STATE_CREATED"

	// UserStateUSERSTATEACTIVE captures enum value "USER_STATE_ACTIVE"
	UserStateUSERSTATEACTIVE UserState = "USER_STATE_ACTIVE"

	// UserStateUSERSTATEINACTIVE captures enum value "USER_STATE_INACTIVE"
	UserStateUSERSTATEINACTIVE UserState = "USER_STATE_INACTIVE"

	// UserStateUSERSTATESIGNEDUP captures enum value "USER_STATE_SIGNEDUP"
	UserStateUSERSTATESIGNEDUP UserState = "USER_STATE_SIGNEDUP"

	// UserStateUSERSTATESUSPENDED captures enum value "USER_STATE_SUSPENDED"
	UserStateUSERSTATESUSPENDED UserState = "USER_STATE_SUSPENDED"
)

// for schema
var userStateEnum []interface{}

func init() {
	var res []UserState
	if err := json.Unmarshal([]byte(`["USER_STATE_UNSPECIFIED","USER_STATE_CREATED","USER_STATE_ACTIVE","USER_STATE_INACTIVE","USER_STATE_SIGNEDUP","USER_STATE_SUSPENDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userStateEnum = append(userStateEnum, v)
	}
}

func (m UserState) validateUserStateEnum(path, location string, value UserState) error {
	if err := validate.EnumCase(path, location, value, userStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this user state
func (m UserState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this user state based on context it is used
func (m UserState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
