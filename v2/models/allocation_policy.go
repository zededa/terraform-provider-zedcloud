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

// AllocationPolicy Iot hub allocation policy.
//
// swagger:model AllocationPolicy
type AllocationPolicy string

func NewAllocationPolicy(value AllocationPolicy) *AllocationPolicy {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AllocationPolicy.
func (m AllocationPolicy) Pointer() *AllocationPolicy {
	return &m
}

const (

	// AllocationPolicyALLOCATIONPOLICYUNSPECIFIED captures enum value "ALLOCATION_POLICY_UNSPECIFIED"
	AllocationPolicyALLOCATIONPOLICYUNSPECIFIED AllocationPolicy = "ALLOCATION_POLICY_UNSPECIFIED"

	// AllocationPolicyALLOCATIONPOLICYHASHED captures enum value "ALLOCATION_POLICY_HASHED"
	AllocationPolicyALLOCATIONPOLICYHASHED AllocationPolicy = "ALLOCATION_POLICY_HASHED"

	// AllocationPolicyALLOCATIONPOLICYGEOLATENCY captures enum value "ALLOCATION_POLICY_GEOLATENCY"
	AllocationPolicyALLOCATIONPOLICYGEOLATENCY AllocationPolicy = "ALLOCATION_POLICY_GEOLATENCY"

	// AllocationPolicyALLOCATIONPOLICYSTATIC captures enum value "ALLOCATION_POLICY_STATIC"
	AllocationPolicyALLOCATIONPOLICYSTATIC AllocationPolicy = "ALLOCATION_POLICY_STATIC"

	// AllocationPolicyALLOCATIONPOLICYCUSTOM captures enum value "ALLOCATION_POLICY_CUSTOM"
	AllocationPolicyALLOCATIONPOLICYCUSTOM AllocationPolicy = "ALLOCATION_POLICY_CUSTOM"
)

// for schema
var allocationPolicyEnum []interface{}

func init() {
	var res []AllocationPolicy
	if err := json.Unmarshal([]byte(`["ALLOCATION_POLICY_UNSPECIFIED","ALLOCATION_POLICY_HASHED","ALLOCATION_POLICY_GEOLATENCY","ALLOCATION_POLICY_STATIC","ALLOCATION_POLICY_CUSTOM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		allocationPolicyEnum = append(allocationPolicyEnum, v)
	}
}

func (m AllocationPolicy) validateAllocationPolicyEnum(path, location string, value AllocationPolicy) error {
	if err := validate.EnumCase(path, location, value, allocationPolicyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this allocation policy
func (m AllocationPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAllocationPolicyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this allocation policy based on context it is used
func (m AllocationPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
