// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VMManifestImage VM manifest image
//
// swagger:model VMManifestImage
type VMManifestImage struct {

	// UI map: AppEditPage:DrivesPane:Cleartext, AppDetailsPage:DrivesPane:ClearText_Field
	Cleartext bool `json:"cleartext,omitempty"`

	// enum: CDROM, HDD, NET
	//
	// UI map: AppEditPage:DrivesPane:Drive_Type_Field, AppDetailsPage:DrivesPane:Drive_Type_Field
	Drvtype string `json:"drvtype,omitempty"`

	// UI map: AppEditPage:DrivesPane:Ignorepurge, AppDetailsPage:DrivesPane:Ignorepurgee_Field
	Ignorepurge bool `json:"ignorepurge,omitempty"`

	// UI map: AppEditPage:DrivesPane:Image_Format_Field, AppDetailsPage:DrivesPane:Image_Format_Field
	Imageformat *ConfigFormat `json:"imageformat,omitempty"`

	// UI map: AppEditPage:DrivesPane:Image_ID_Field, AppDetailsPage:DrivesPane:Image_ID_Field
	Imageid string `json:"imageid,omitempty"`

	// UI map: AppEditPage:DrivesPane:Image_Name_Field, AppDetailsPage:DrivesPane:Image_Name_Field
	Imagename string `json:"imagename,omitempty"`

	// UI map: AppEditPage:DrivesPane:Max_Size_Field, AppDetailsPage:DrivesPane:Max_Size_Field
	Maxsize string `json:"maxsize,omitempty"`

	// UI map: AppEditPage:DrivesPane:Mountpath, AppDetailsPage:DrivesPane:Mountpath_Field
	Mountpath string `json:"mountpath,omitempty"`

	// Not used by ZedUI
	Params []*Param `json:"params"`

	// UI map: AppEditPage:DrivesPane:Preserve_Field, AppDetailsPage:DrivesPane:Preserve_Field
	Preserve bool `json:"preserve,omitempty"`

	// Not used by ZedUI
	Readonly bool `json:"readonly,omitempty"`

	// enum: Disk, Kernel, Initrd, RamDisk
	//
	// UI map: AppEditPage:DrivesPane:Target_Field, AppDetailsPage:DrivesPane:Target_Field
	Target string `json:"target,omitempty"`

	// UI map: AppEditPage:DrivesPane:Volume_Label, AppDetailsPage:DrivesPane:Volume_Label
	Volumelabel string `json:"volumelabel,omitempty"`
}

// Validate validates this VM manifest image
func (m *VMManifestImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageformat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMManifestImage) validateImageformat(formats strfmt.Registry) error {
	if swag.IsZero(m.Imageformat) { // not required
		return nil
	}

	if m.Imageformat != nil {
		if err := m.Imageformat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageformat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageformat")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifestImage) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(m.Params) { // not required
		return nil
	}

	for i := 0; i < len(m.Params); i++ {
		if swag.IsZero(m.Params[i]) { // not required
			continue
		}

		if m.Params[i] != nil {
			if err := m.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this VM manifest image based on the context it is used
func (m *VMManifestImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImageformat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMManifestImage) contextValidateImageformat(ctx context.Context, formats strfmt.Registry) error {

	if m.Imageformat != nil {
		if err := m.Imageformat.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageformat")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageformat")
			}
			return err
		}
	}

	return nil
}

func (m *VMManifestImage) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Params); i++ {

		if m.Params[i] != nil {
			if err := m.Params[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMManifestImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMManifestImage) UnmarshalBinary(b []byte) error {
	var res VMManifestImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
