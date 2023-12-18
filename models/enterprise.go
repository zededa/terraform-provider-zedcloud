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
	"github.com/go-openapi/validate"
)

// Enterprise Enterprise detail
//
// # Enterprise meta data
//
// swagger:model Enterprise
type Enterprise struct {

	// hubspot Id
	HubspotID string `json:"HubspotId,omitempty"`

	// sfdc Id
	SfdcID string `json:"SfdcId,omitempty"`

	// enterprise level key-value pairs
	Attributes map[string]string `json:"attributes,omitempty"`

	// Azure subscription ID tied to this enterprise
	AzureSubID string `json:"azureSubId,omitempty"`

	// List of all child enterprises
	ChildEnterprises []*EnterpriseSummary `json:"childEnterprises"`

	// zedcontrol host
	ControllerHostURL string `json:"controllerHostURL,omitempty"`

	// Detailed description of the enterprise
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// Unique system defined enterprise ID
	// Read Only: true
	// Pattern: [0-9A-Za-z_=-]{28}
	ID string `json:"id,omitempty"`

	// Perform authorization using parent enterprise
	InheritAuthFromParent bool `json:"inheritAuthFromParent,omitempty"`

	// User defined name of the enterprise. Once enterprise is created, name can’t be changed
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// string id = 100;
	// string name = 101;
	//
	// Parent enterprise ID
	// Pattern: [0-9A-Za-z_=-]{28}
	ParentEntpID string `json:"parentEntpId,omitempty"`

	// Policy version list
	PolicyList *PolicyVersionList `json:"policyList,omitempty"`

	// List of realms associated with the enterprise
	Realms []string `json:"realms"`

	// system defined info
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// Enterprise state
	State *EnterpriseState `json:"state,omitempty"`

	// Enable / Disable streaming of events to 3rd party end point
	StreamEvents *DataStream `json:"streamEvents,omitempty"`

	// User defined title for the enterprise. Title can be changed any time
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title *string `json:"title"`

	// Enterprise settings to enforce TOTP
	TotpSettings *TOTPSettings `json:"totpSettings,omitempty"`

	// Enterprise type
	Type *EnterpriseType `json:"type,omitempty"`
}

// Validate validates this enterprise
func (m *Enterprise) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildEnterprises(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentEntpID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreamEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotpSettings(formats); err != nil {
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

func (m *Enterprise) validateChildEnterprises(formats strfmt.Registry) error {
	if swag.IsZero(m.ChildEnterprises) { // not required
		return nil
	}

	for i := 0; i < len(m.ChildEnterprises); i++ {
		if swag.IsZero(m.ChildEnterprises[i]) { // not required
			continue
		}

		if m.ChildEnterprises[i] != nil {
			if err := m.ChildEnterprises[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childEnterprises" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("childEnterprises" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Enterprise) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *Enterprise) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[0-9A-Za-z_=-]{28}`); err != nil {
		return err
	}

	return nil
}

func (m *Enterprise) validateName(formats strfmt.Registry) error {

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

func (m *Enterprise) validateParentEntpID(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentEntpID) { // not required
		return nil
	}

	if err := validate.Pattern("parentEntpId", "body", m.ParentEntpID, `[0-9A-Za-z_=-]{28}`); err != nil {
		return err
	}

	return nil
}

func (m *Enterprise) validatePolicyList(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyList) { // not required
		return nil
	}

	if m.PolicyList != nil {
		if err := m.PolicyList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policyList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policyList")
			}
			return err
		}
	}

	return nil
}

func (m *Enterprise) validateRevision(formats strfmt.Registry) error {
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

func (m *Enterprise) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *Enterprise) validateStreamEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.StreamEvents) { // not required
		return nil
	}

	if m.StreamEvents != nil {
		if err := m.StreamEvents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streamEvents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streamEvents")
			}
			return err
		}
	}

	return nil
}

func (m *Enterprise) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", *m.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

func (m *Enterprise) validateTotpSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.TotpSettings) { // not required
		return nil
	}

	if m.TotpSettings != nil {
		if err := m.TotpSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("totpSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("totpSettings")
			}
			return err
		}
	}

	return nil
}

func (m *Enterprise) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
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

// ContextValidate validate this enterprise based on the context it is used
func (m *Enterprise) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChildEnterprises(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStreamEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotpSettings(ctx, formats); err != nil {
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

func (m *Enterprise) contextValidateChildEnterprises(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChildEnterprises); i++ {

		if m.ChildEnterprises[i] != nil {

			if swag.IsZero(m.ChildEnterprises[i]) { // not required
				return nil
			}

			if err := m.ChildEnterprises[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childEnterprises" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("childEnterprises" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Enterprise) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Enterprise) contextValidatePolicyList(ctx context.Context, formats strfmt.Registry) error {

	if m.PolicyList != nil {

		if swag.IsZero(m.PolicyList) { // not required
			return nil
		}

		if err := m.PolicyList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policyList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policyList")
			}
			return err
		}
	}

	return nil
}

func (m *Enterprise) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {

		if swag.IsZero(m.Revision) { // not required
			return nil
		}

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

func (m *Enterprise) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {

		if swag.IsZero(m.State) { // not required
			return nil
		}

		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *Enterprise) contextValidateStreamEvents(ctx context.Context, formats strfmt.Registry) error {

	if m.StreamEvents != nil {

		if swag.IsZero(m.StreamEvents) { // not required
			return nil
		}

		if err := m.StreamEvents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streamEvents")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streamEvents")
			}
			return err
		}
	}

	return nil
}

func (m *Enterprise) contextValidateTotpSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.TotpSettings != nil {

		if swag.IsZero(m.TotpSettings) { // not required
			return nil
		}

		if err := m.TotpSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("totpSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("totpSettings")
			}
			return err
		}
	}

	return nil
}

func (m *Enterprise) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

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
func (m *Enterprise) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Enterprise) UnmarshalBinary(b []byte) error {
	var res Enterprise
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
