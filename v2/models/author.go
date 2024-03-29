// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Author author
//
// swagger:model Author
type Author struct {

	// Company of the developer
	//
	// UI map: AppEditPage:IdentityPane:Category_Field, AppDetailsPage:IdentityPane:Category_Field
	Company string `json:"company,omitempty"`

	// Contact email of the developer company
	//
	// UI map: AppEditPage:DeveloperPane:Email_Field, AppDetailsPage:DeveloperPane:Email_Field
	Email string `json:"email,omitempty"`

	// Not used by ZedUI
	Group string `json:"group,omitempty"`

	// Name of the developer
	//
	// UI map: AppEditPage:DeveloperPane:Name_Field, AppDetailsPage:DeveloperPane:Name_Field
	User string `json:"user,omitempty"`

	// Website of the developer company
	//
	// UI map: AppEditPage:DeveloperPane:Website_Field, AppDetailsPage:DeveloperPane:Website_Field
	Website string `json:"website,omitempty"`
}

// Validate validates this author
func (m *Author) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this author based on context it is used
func (m *Author) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Author) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Author) UnmarshalBinary(b []byte) error {
	var res Author
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
