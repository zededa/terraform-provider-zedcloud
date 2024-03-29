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

// PhysicalStorage physical storage
//
// swagger:model PhysicalStorage
type PhysicalStorage struct {

	// compression ratio
	CompressionRatio float64 `json:"compressionRatio,omitempty"`

	// count zvols
	CountZvols int64 `json:"countZvols,omitempty"`

	// current raid
	CurrentRaid *PhysicalStorageRaidType `json:"currentRaid,omitempty"`

	// disks
	Disks []*PhysicalStorageDiskState `json:"disks"`

	// pool name
	PoolName string `json:"poolName,omitempty"`

	// Zpool Status message sent by EVE
	PoolStatusMsg string `json:"poolStatusMsg,omitempty"`

	// storage state
	StorageState *PhysicalStorageStatus `json:"storageState,omitempty"`

	// storage type
	StorageType *PhysicalStorageTypeInfo `json:"storageType,omitempty"`

	// zfs version
	ZfsVersion string `json:"zfsVersion,omitempty"`

	// zpool size
	ZpoolSize string `json:"zpoolSize,omitempty"`
}

// Validate validates this physical storage
func (m *PhysicalStorage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentRaid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhysicalStorage) validateCurrentRaid(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentRaid) { // not required
		return nil
	}

	if m.CurrentRaid != nil {
		if err := m.CurrentRaid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentRaid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currentRaid")
			}
			return err
		}
	}

	return nil
}

func (m *PhysicalStorage) validateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PhysicalStorage) validateStorageState(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageState) { // not required
		return nil
	}

	if m.StorageState != nil {
		if err := m.StorageState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageState")
			}
			return err
		}
	}

	return nil
}

func (m *PhysicalStorage) validateStorageType(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageType) { // not required
		return nil
	}

	if m.StorageType != nil {
		if err := m.StorageType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this physical storage based on the context it is used
func (m *PhysicalStorage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrentRaid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PhysicalStorage) contextValidateCurrentRaid(ctx context.Context, formats strfmt.Registry) error {

	if m.CurrentRaid != nil {
		if err := m.CurrentRaid.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentRaid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currentRaid")
			}
			return err
		}
	}

	return nil
}

func (m *PhysicalStorage) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PhysicalStorage) contextValidateStorageState(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageState != nil {
		if err := m.StorageState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageState")
			}
			return err
		}
	}

	return nil
}

func (m *PhysicalStorage) contextValidateStorageType(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageType != nil {
		if err := m.StorageType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PhysicalStorage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhysicalStorage) UnmarshalBinary(b []byte) error {
	var res PhysicalStorage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
