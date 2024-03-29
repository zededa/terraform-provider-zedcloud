// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceRawMetrics device raw metrics
//
// swagger:model DeviceRawMetrics
type DeviceRawMetrics struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// raw metrics
	RawMetrics string `json:"rawMetrics,omitempty"`
}

// Validate validates this device raw metrics
func (m *DeviceRawMetrics) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this device raw metrics based on context it is used
func (m *DeviceRawMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceRawMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceRawMetrics) UnmarshalBinary(b []byte) error {
	var res DeviceRawMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
