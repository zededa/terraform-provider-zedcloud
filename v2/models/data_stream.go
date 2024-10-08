// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataStream data stream
//
// swagger:model DataStream
type DataStream struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// plugin Id
	PluginID string `json:"pluginId,omitempty"`
}

// Validate validates this data stream
func (m *DataStream) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this data stream based on context it is used
func (m *DataStream) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataStream) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataStream) UnmarshalBinary(b []byte) error {
	var res DataStream
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
