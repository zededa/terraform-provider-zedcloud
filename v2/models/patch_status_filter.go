// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchStatusFilter patch status filter
//
// swagger:model PatchStatusFilter
type PatchStatusFilter struct {

	// device name
	DeviceName string `json:"deviceName,omitempty"`
}

// Validate validates this patch status filter
func (m *PatchStatusFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch status filter based on context it is used
func (m *PatchStatusFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchStatusFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchStatusFilter) UnmarshalBinary(b []byte) error {
	var res PatchStatusFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
