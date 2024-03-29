// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkCounters NetworkCounter is used to store the Network Stats and Counters
//
// # NetworkCounter is used to store the Network Stats and Counters
//
// swagger:model NetworkCounters
type NetworkCounters struct {

	// ifName
	IfName string `json:"ifName,omitempty"`

	// Rx ACL Rate Drops
	RxACLDrops uint64 `json:"rxAclDrops,omitempty"`

	// Rx ACL Rate Limit Drops
	RxACLRateLimitDrops uint64 `json:"rxAclRateLimitDrops,omitempty"`

	// Rx Bytes
	RxBytes uint64 `json:"rxBytes,omitempty"`

	// Rx Drops
	RxDrops uint64 `json:"rxDrops,omitempty"`

	// Rx Errors
	RxErrors uint64 `json:"rxErrors,omitempty"`

	// Rx packets
	RxPkts uint64 `json:"rxPkts,omitempty"`

	// Tx ACL Rate Drops
	TxACLDrops uint64 `json:"txAclDrops,omitempty"`

	// Tx ACL Rate Limit Drops
	TxACLRateLimitDrops uint64 `json:"txAclRateLimitDrops,omitempty"`

	// Tx Bytes
	TxBytes uint64 `json:"txBytes,omitempty"`

	// Tx Drops
	TxDrops uint64 `json:"txDrops,omitempty"`

	// Tx Errors
	TxErrors uint64 `json:"txErrors,omitempty"`

	// Tx Packets
	TxPkts uint64 `json:"txPkts,omitempty"`
}

// Validate validates this network counters
func (m *NetworkCounters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network counters based on context it is used
func (m *NetworkCounters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCounters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCounters) UnmarshalBinary(b []byte) error {
	var res NetworkCounters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
