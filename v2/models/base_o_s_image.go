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

// BaseOSImage Base OS image payload detail
//
// # Base OS image request details
//
// swagger:model BaseOSImage
type BaseOSImage struct {

	// activation flag
	// Required: true
	Activate *bool `json:"activate"`

	// image name
	// Required: true
	ImageName *string `json:"imageName"`

	// immutable Volume for this base image
	ImvolID string `json:"imvolId,omitempty"`

	// system generated unique id for an image
	// Required: true
	UUID *string `json:"uuid"`

	// image version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this base o s image
func (m *BaseOSImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseOSImage) validateActivate(formats strfmt.Registry) error {

	if err := validate.Required("activate", "body", m.Activate); err != nil {
		return err
	}

	return nil
}

func (m *BaseOSImage) validateImageName(formats strfmt.Registry) error {

	if err := validate.Required("imageName", "body", m.ImageName); err != nil {
		return err
	}

	return nil
}

func (m *BaseOSImage) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

func (m *BaseOSImage) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this base o s image based on context it is used
func (m *BaseOSImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseOSImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseOSImage) UnmarshalBinary(b []byte) error {
	var res BaseOSImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
