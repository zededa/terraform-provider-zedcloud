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

// IoType Input/Output Type
//
// - IO_TYPE_UNSPECIFIED: No operation/ Invalid peration - IO_TYPE_ETH: Ethernet - IO_TYPE_USB: USB Type - IO_TYPE_COM: Communication Port - IO_TYPE_AUDIO: Audio Port - IO_TYPE_WLAN: wireless LAN - IO_TYPE_WWAN: Wireless wide area network - IO_TYPE_HDMI: High-Definition Multimedia Interface - IO_TYPE_LTE: LTE Interfaces - IO_TYPE_OTHER: Other Io Types
//
// swagger:model IoType
type IoType string

func NewIoType(value IoType) *IoType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IoType.
func (m IoType) Pointer() *IoType {
	return &m
}

const (

	// IoTypeIOTYPEUNSPECIFIED captures enum value "IO_TYPE_UNSPECIFIED"
	IoTypeIOTYPEUNSPECIFIED IoType = "IO_TYPE_UNSPECIFIED"

	// IoTypeIOTYPEETH captures enum value "IO_TYPE_ETH"
	IoTypeIOTYPEETH IoType = "IO_TYPE_ETH"

	// IoTypeIOTYPEUSB captures enum value "IO_TYPE_USB"
	IoTypeIOTYPEUSB IoType = "IO_TYPE_USB"

	// IoTypeIOTYPECOM captures enum value "IO_TYPE_COM"
	IoTypeIOTYPECOM IoType = "IO_TYPE_COM"

	// IoTypeIOTYPEAUDIO captures enum value "IO_TYPE_AUDIO"
	IoTypeIOTYPEAUDIO IoType = "IO_TYPE_AUDIO"

	// IoTypeIOTYPEWLAN captures enum value "IO_TYPE_WLAN"
	IoTypeIOTYPEWLAN IoType = "IO_TYPE_WLAN"

	// IoTypeIOTYPEWWAN captures enum value "IO_TYPE_WWAN"
	IoTypeIOTYPEWWAN IoType = "IO_TYPE_WWAN"

	// IoTypeIOTYPEHDMI captures enum value "IO_TYPE_HDMI"
	IoTypeIOTYPEHDMI IoType = "IO_TYPE_HDMI"

	// IoTypeIOTYPELTE captures enum value "IO_TYPE_LTE"
	IoTypeIOTYPELTE IoType = "IO_TYPE_LTE"

	// IoTypeIOTYPEOTHER captures enum value "IO_TYPE_OTHER"
	IoTypeIOTYPEOTHER IoType = "IO_TYPE_OTHER"
)

// for schema
var ioTypeEnum []interface{}

func init() {
	var res []IoType
	if err := json.Unmarshal([]byte(`["IO_TYPE_UNSPECIFIED","IO_TYPE_ETH","IO_TYPE_USB","IO_TYPE_COM","IO_TYPE_AUDIO","IO_TYPE_WLAN","IO_TYPE_WWAN","IO_TYPE_HDMI","IO_TYPE_LTE","IO_TYPE_OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ioTypeEnum = append(ioTypeEnum, v)
	}
}

func (m IoType) validateIoTypeEnum(path, location string, value IoType) error {
	if err := validate.EnumCase(path, location, value, ioTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this io type
func (m IoType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIoTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this io type based on context it is used
func (m IoType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
