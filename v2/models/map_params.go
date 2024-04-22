// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MapParams map params
//
// swagger:model MapParams
type MapParams struct {

	// Application Port value
	AppPort int64 `json:"appPort,omitempty"`
}

// Validate validates this map params
func (m *MapParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this map params based on context it is used
func (m *MapParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MapParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MapParams) UnmarshalBinary(b []byte) error {
	var res MapParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}