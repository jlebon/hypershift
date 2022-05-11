// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlacementGroupCreate placement group create
//
// swagger:model PlacementGroupCreate
type PlacementGroupCreate struct {

	// The name of the Placement Group
	// Required: true
	Name *string `json:"name"`

	// The Placement Group Policy
	// Required: true
	// Enum: [affinity anti-affinity]
	Policy *string `json:"policy"`
}

// Validate validates this placement group create
func (m *PlacementGroupCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlacementGroupCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var placementGroupCreateTypePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["affinity","anti-affinity"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		placementGroupCreateTypePolicyPropEnum = append(placementGroupCreateTypePolicyPropEnum, v)
	}
}

const (

	// PlacementGroupCreatePolicyAffinity captures enum value "affinity"
	PlacementGroupCreatePolicyAffinity string = "affinity"

	// PlacementGroupCreatePolicyAntiDashAffinity captures enum value "anti-affinity"
	PlacementGroupCreatePolicyAntiDashAffinity string = "anti-affinity"
)

// prop value enum
func (m *PlacementGroupCreate) validatePolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, placementGroupCreateTypePolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlacementGroupCreate) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", m.Policy); err != nil {
		return err
	}

	// value enum
	if err := m.validatePolicyEnum("policy", "body", *m.Policy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this placement group create based on context it is used
func (m *PlacementGroupCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlacementGroupCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlacementGroupCreate) UnmarshalBinary(b []byte) error {
	var res PlacementGroupCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}