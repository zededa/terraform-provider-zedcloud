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

// NetworkInstance Network instance detailed configuration
//
// Network instance provides Edge applications a variety of connectivity choices for all types of networks. This enables logical secure connectivity between Edge applications within a single Edge node and within a logical group of Edge nodes. This provides detailed configuration of a Network instance.
// Example: {"id":"d85a545f-6510-4327-b03d-c02eef119e99","name":"sample-app"}
//
// swagger:model NetworkInstance
type NetworkInstance struct {

	// id of the Cluster in which the network instance is configured
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	ClusterID string `json:"clusterID,omitempty"`

	// Detailed description of the network instance
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// flag to indicate if this is default network
	//
	// flag to indicate if this is the default network instance for the device
	DeviceDefault bool `json:"deviceDefault,omitempty"`

	// device on which this network is running
	//
	// id of the device on which network instance is created
	// Required: true
	DeviceID *string `json:"deviceId"`

	// dhcp - DEPRECATED
	//
	// Deprecated
	Dhcp bool `json:"dhcp,omitempty"`

	// List of Static DNS entries
	DNSList []*StaticDNSList `json:"dnsList"`

	// System defined universally unique Id of the network instance
	// Read Only: true
	// Pattern: [0-9A-Za-z-]+
	ID string `json:"id,omitempty"`

	// Dhcp Server Configuration
	IP *DhcpServerConfig `json:"ip,omitempty"`

	// Kind of Network Instance ( Local, Switch etc )
	// Required: true
	Kind *NetworkInstanceKind `json:"kind"`

	// Lisp Config : read only for now. Deprecated.
	Lisp *LispConfig `json:"lisp,omitempty"`

	// Maximum transmission unit (MTU) to set for the network instance and all application interfaces connected to it
	Mtu int64 `json:"mtu,omitempty"`

	// User defined name of the network instance, unique across the enterprise. Once object is created, name can’t be changed
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// id of the network policy to be attached to this network instance
	NetworkPolicyID string `json:"networkPolicyId,omitempty"`

	// Deprecated
	Oconfig string `json:"oconfig,omitempty"`

	// Service specific Config
	Opaque *NetInstOpaqueConfig `json:"opaque,omitempty"`

	// name of port mapping in the model
	//
	// name of port mapping in the model
	// Required: true
	Port *string `json:"port"`

	// Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.
	PortTags map[string]string `json:"portTags,omitempty"`

	// id of the project in which network instance is created
	ProjectID string `json:"projectId,omitempty"`

	// Automatically propagate connected routes
	PropagateConnectedRoutes bool `json:"propagateConnectedRoutes,omitempty"`

	// system defined info for the object
	Revision *ObjectRevision `json:"revision,omitempty"`

	// List of Static IP routes
	StaticRoutes []*StaticIPRoute `json:"staticRoutes"`

	// Tags are name/value pairs that enable you to categorize resources. Tag names are case insensitive with max_length 512 and min_length 3. Tag values are case sensitive with max_length 256 and min_length 3.
	Tags map[string]string `json:"tags,omitempty"`

	// User defined title of the network instance. Title can be changed at any time
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: ^[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+$
	Title *string `json:"title"`

	// Type of DHCP for this Network Instance
	Type *NetworkInstanceDhcpType `json:"type,omitempty"`
}

// Validate validates this net inst config
func (m *NetworkInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLisp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpaque(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticRoutes(formats); err != nil {
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

func (m *NetworkInstance) validateClusterID(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterID) { // not required
		return nil
	}

	if err := validate.MinLength("clusterID", "body", m.ClusterID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("clusterID", "body", m.ClusterID, 256); err != nil {
		return err
	}

	if err := validate.Pattern("clusterID", "body", m.ClusterID, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *NetworkInstance) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *NetworkInstance) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("deviceId", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *NetworkInstance) validateDNSList(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSList) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSList); i++ {
		if swag.IsZero(m.DNSList[i]) { // not required
			continue
		}

		if m.DNSList[i] != nil {
			if err := m.DNSList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInstance) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[0-9A-Za-z-]+`); err != nil {
		return err
	}

	return nil
}

func (m *NetworkInstance) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if m.IP != nil {
		if err := m.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstance) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	if m.Kind != nil {
		if err := m.Kind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstance) validateLisp(formats strfmt.Registry) error {
	if swag.IsZero(m.Lisp) { // not required
		return nil
	}

	if m.Lisp != nil {
		if err := m.Lisp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lisp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lisp")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstance) validateName(formats strfmt.Registry) error {

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

func (m *NetworkInstance) validateOpaque(formats strfmt.Registry) error {
	if swag.IsZero(m.Opaque) { // not required
		return nil
	}

	if m.Opaque != nil {
		if err := m.Opaque.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opaque")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opaque")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstance) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *NetworkInstance) validateRevision(formats strfmt.Registry) error {
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

func (m *NetworkInstance) validateStaticRoutes(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticRoutes) { // not required
		return nil
	}

	for i := 0; i < len(m.StaticRoutes); i++ {
		if swag.IsZero(m.StaticRoutes[i]) { // not required
			continue
		}

		if m.StaticRoutes[i] != nil {
			if err := m.StaticRoutes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("staticRoutes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("staticRoutes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInstance) validateTitle(formats strfmt.Registry) error {

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

func (m *NetworkInstance) validateType(formats strfmt.Registry) error {
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

// ContextValidate validate this net inst config based on the context it is used
func (m *NetworkInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLisp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpaque(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticRoutes(ctx, formats); err != nil {
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

func (m *NetworkInstance) contextValidateDNSList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSList); i++ {

		if m.DNSList[i] != nil {
			if err := m.DNSList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInstance) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NetworkInstance) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if m.IP != nil {
		if err := m.IP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstance) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if m.Kind != nil {
		if err := m.Kind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kind")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstance) contextValidateLisp(ctx context.Context, formats strfmt.Registry) error {

	if m.Lisp != nil {
		if err := m.Lisp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lisp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lisp")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstance) contextValidateOpaque(ctx context.Context, formats strfmt.Registry) error {

	if m.Opaque != nil {
		if err := m.Opaque.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opaque")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opaque")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkInstance) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NetworkInstance) contextValidateStaticRoutes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StaticRoutes); i++ {

		if m.StaticRoutes[i] != nil {

			if swag.IsZero(m.StaticRoutes[i]) { // not required
				return nil
			}

			if err := m.StaticRoutes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("staticRoutes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("staticRoutes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInstance) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NetworkInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInstance) UnmarshalBinary(b []byte) error {
	var res NetworkInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
