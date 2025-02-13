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

// Policy Policy detailed configuration for a resource group (project)
//
// Policy configuration defines set of policies to be applied on a resource grup (project). Policy can be one of the types specified in PolicyType. Multiple policies can be applied to a single resource group (project).
//
// swagger:model Policy
type Policy struct {

	// app policy, which is used in auto app instance deployment
	AppPolicy *AppPolicy `json:"appPolicy,omitempty"`

	// attestation policy to enforce on all devices in this project
	AttestationPolicy *AttestationPolicy `json:"attestationPolicy,omitempty"`

	// Mapping of policy  variable keys and policy variable values
	Attr map[string]string `json:"attr,omitempty"`

	// azure policy, which is used in configuring azure iot-edge.
	AzurePolicy *AzurePolicy `json:"azurePolicy,omitempty"`

	// cluster policy to bring up cluster on devices in this project
	ClusterPolicy *ClusterPolicy `json:"clusterPolicy,omitempty"`

	// configuration lock policy to enforce on all devices in this project
	ConfigurationLockPolicy *ConfigurationLockPolicy `json:"configurationLockPolicy,omitempty"`

	// Detailed description of the policy
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// edgeview policy on devices of this project
	EdgeviewPolicy *EdgeviewPolicy `json:"edgeviewPolicy,omitempty"`

	// System defined universally unique Id of the policy request
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// local operator console policy to enforce on all devices in this project
	LocalOperatorConsolePolicy *LocalOperatorConsolePolicy `json:"localOperatorConsolePolicy,omitempty"`

	// module policy, which is used in auto module deployment
	ModulePolicy *ModulePolicy `json:"modulePolicy,omitempty"`

	// User defined name of the policy request, unique across the enterprise. Once object is created, name can’t be changed
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// network policy to enforce on all devices in this project
	NetworkPolicy *NetworkPolicy `json:"networkPolicy,omitempty"`

	// system defined info
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// status of the policy
	// Read Only: true
	Status *PolicyStatus `json:"status,omitempty"`

	// Detailed status message of the policy
	// Max Length: 256
	StatusMessage string `json:"statusMessage,omitempty"`

	// User defined title of the policy. Title can be changed at any time
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: ^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$
	Title *string `json:"title"`

	// type of policy
	// Required: true
	Type *PolicyType `json:"type"`
}

// Validate validates this policy config
func (m *Policy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttestationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzurePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigurationLockPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeviewPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalOperatorConsolePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModulePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Policy) validateAppPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AppPolicy) { // not required
		return nil
	}

	if m.AppPolicy != nil {
		if err := m.AppPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateAttestationPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AttestationPolicy) { // not required
		return nil
	}

	if m.AttestationPolicy != nil {
		if err := m.AttestationPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attestationPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attestationPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateAzurePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.AzurePolicy) { // not required
		return nil
	}

	if m.AzurePolicy != nil {
		if err := m.AzurePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azurePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azurePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateClusterPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterPolicy) { // not required
		return nil
	}

	if m.ClusterPolicy != nil {
		if err := m.ClusterPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateConfigurationLockPolicy(formats strfmt.Registry) error {
	if m.ConfigurationLockPolicy != nil {
		if err := m.ConfigurationLockPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configurationLockPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configurationLockPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateEdgeviewPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.EdgeviewPolicy) { // not required
		return nil
	}

	if m.EdgeviewPolicy != nil {
		if err := m.EdgeviewPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeviewPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeviewPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateLocalOperatorConsolePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalOperatorConsolePolicy) { // not required
		return nil
	}

	if m.LocalOperatorConsolePolicy != nil {
		if err := m.LocalOperatorConsolePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localOperatorConsolePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localOperatorConsolePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateModulePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ModulePolicy) { // not required
		return nil
	}

	if m.ModulePolicy != nil {
		if err := m.ModulePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modulePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modulePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateNetworkPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkPolicy) { // not required
		return nil
	}

	if m.NetworkPolicy != nil {
		if err := m.NetworkPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateStatusMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusMessage) { // not required
		return nil
	}

	if err := validate.MaxLength("statusMessage", "body", m.StatusMessage, 256); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", *m.Title, `^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this policy config based on the context it is used
func (m *Policy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttestationPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzurePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfigurationLockPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeviewPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalOperatorConsolePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModulePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Policy) contextValidateAppPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AppPolicy != nil {
		if err := m.AppPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateAttestationPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AttestationPolicy != nil {
		if err := m.AttestationPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attestationPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attestationPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateAzurePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.AzurePolicy != nil {
		if err := m.AzurePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azurePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azurePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateClusterPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterPolicy != nil {
		if err := m.ClusterPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateConfigurationLockPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigurationLockPolicy != nil {
		if err := m.ConfigurationLockPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configurationLockPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configurationLockPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateEdgeviewPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.EdgeviewPolicy != nil {
		if err := m.EdgeviewPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeviewPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeviewPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Policy) contextValidateLocalOperatorConsolePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.LocalOperatorConsolePolicy != nil {
		if err := m.LocalOperatorConsolePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("localOperatorConsolePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("localOperatorConsolePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateModulePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.ModulePolicy != nil {
		if err := m.ModulePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modulePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modulePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateNetworkPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkPolicy != nil {
		if err := m.NetworkPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("networkPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {
		if err := m.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Policy) UnmarshalBinary(b []byte) error {
	var res Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
