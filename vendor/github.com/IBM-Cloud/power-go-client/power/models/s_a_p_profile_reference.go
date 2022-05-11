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

// SAPProfileReference s a p profile reference
//
// swagger:model SAPProfileReference
type SAPProfileReference struct {

	// Link to SAP profile resource
	// Required: true
	Href *string `json:"href"`

	// SAP Profile ID
	// Required: true
	ProfileID *string `json:"profileID"`
}

// Validate validates this s a p profile reference
func (m *SAPProfileReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SAPProfileReference) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *SAPProfileReference) validateProfileID(formats strfmt.Registry) error {

	if err := validate.Required("profileID", "body", m.ProfileID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this s a p profile reference based on context it is used
func (m *SAPProfileReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SAPProfileReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SAPProfileReference) UnmarshalBinary(b []byte) error {
	var res SAPProfileReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}