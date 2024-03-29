// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExternalOpaqueBinaryBlob external opaque binary blob
//
// swagger:model ExternalOpaqueBinaryBlob
type ExternalOpaqueBinaryBlob struct {

	// Eve shall use the image name when fileNameToUse is empty
	// optional - this can be image type or size etc encoded into a single string.
	// only the application instance agents will interpret this.
	BlobMetaData string `json:"blobMetaData,omitempty"`

	// optional - file name to be used for storing this data in EVE.
	// the same file name shall be advertised to application agent
	FileNameToUse string `json:"fileNameToUse,omitempty"`

	// image Id
	ImageID string `json:"imageId,omitempty"`

	// Name of the image uploaded into data store
	ImageName string `json:"imageName,omitempty"`
}

// Validate validates this external opaque binary blob
func (m *ExternalOpaqueBinaryBlob) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this external opaque binary blob based on context it is used
func (m *ExternalOpaqueBinaryBlob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalOpaqueBinaryBlob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalOpaqueBinaryBlob) UnmarshalBinary(b []byte) error {
	var res ExternalOpaqueBinaryBlob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
