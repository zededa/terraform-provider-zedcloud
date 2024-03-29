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

// PatchEnvelopeState - PatchEnvelopeStateErrored: there is an error with config or during download or verification failed
//   - PatchEnvelopeStateReceived: Configuration received but no downloads started
//   - PatchEnvelopeStateDownloading: Artifact/Volume download started
//
// One or more of the artifacts are being downloaded
//   - PatchEnvelopeStateDownloaded: all downloads finished, verified and added to content tree
//   - PatchEnvelopeStateReady: patch envelope ready for application instances
//
// application instances will still not be allowed
// to fetch the patch envelope contents
//   - PatchEnvelopeStateActive: application instances are now allowed to fetch contents
//
// swagger:model PatchEnvelopeState
type PatchEnvelopeState string

func NewPatchEnvelopeState(value PatchEnvelopeState) *PatchEnvelopeState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PatchEnvelopeState.
func (m PatchEnvelopeState) Pointer() *PatchEnvelopeState {
	return &m
}

const (

	// PatchEnvelopeStatePatchEnvelopeStateUnknown captures enum value "PatchEnvelopeStateUnknown"
	PatchEnvelopeStatePatchEnvelopeStateUnknown PatchEnvelopeState = "PatchEnvelopeStateUnknown"

	// PatchEnvelopeStatePatchEnvelopeStateErrored captures enum value "PatchEnvelopeStateErrored"
	PatchEnvelopeStatePatchEnvelopeStateErrored PatchEnvelopeState = "PatchEnvelopeStateErrored"

	// PatchEnvelopeStatePatchEnvelopeStateReceived captures enum value "PatchEnvelopeStateReceived"
	PatchEnvelopeStatePatchEnvelopeStateReceived PatchEnvelopeState = "PatchEnvelopeStateReceived"

	// PatchEnvelopeStatePatchEnvelopeStateDownloading captures enum value "PatchEnvelopeStateDownloading"
	PatchEnvelopeStatePatchEnvelopeStateDownloading PatchEnvelopeState = "PatchEnvelopeStateDownloading"

	// PatchEnvelopeStatePatchEnvelopeStateDownloaded captures enum value "PatchEnvelopeStateDownloaded"
	PatchEnvelopeStatePatchEnvelopeStateDownloaded PatchEnvelopeState = "PatchEnvelopeStateDownloaded"

	// PatchEnvelopeStatePatchEnvelopeStateReady captures enum value "PatchEnvelopeStateReady"
	PatchEnvelopeStatePatchEnvelopeStateReady PatchEnvelopeState = "PatchEnvelopeStateReady"

	// PatchEnvelopeStatePatchEnvelopeStateActive captures enum value "PatchEnvelopeStateActive"
	PatchEnvelopeStatePatchEnvelopeStateActive PatchEnvelopeState = "PatchEnvelopeStateActive"
)

// for schema
var patchEnvelopeStateEnum []interface{}

func init() {
	var res []PatchEnvelopeState
	if err := json.Unmarshal([]byte(`["PatchEnvelopeStateUnknown","PatchEnvelopeStateErrored","PatchEnvelopeStateReceived","PatchEnvelopeStateDownloading","PatchEnvelopeStateDownloaded","PatchEnvelopeStateReady","PatchEnvelopeStateActive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchEnvelopeStateEnum = append(patchEnvelopeStateEnum, v)
	}
}

func (m PatchEnvelopeState) validatePatchEnvelopeStateEnum(path, location string, value PatchEnvelopeState) error {
	if err := validate.EnumCase(path, location, value, patchEnvelopeStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this patch envelope state
func (m PatchEnvelopeState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePatchEnvelopeStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this patch envelope state based on context it is used
func (m PatchEnvelopeState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
