// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LookmlModelExploreSupportedMeasureType lookml model explore supported measure type
// swagger:model LookmlModelExploreSupportedMeasureType
type LookmlModelExploreSupportedMeasureType struct {

	// dimension type
	// Read Only: true
	DimensionType string `json:"dimension_type,omitempty"`

	// measure types
	// Read Only: true
	MeasureTypes []string `json:"measure_types"`
}

// Validate validates this lookml model explore supported measure type
func (m *LookmlModelExploreSupportedMeasureType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LookmlModelExploreSupportedMeasureType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookmlModelExploreSupportedMeasureType) UnmarshalBinary(b []byte) error {
	var res LookmlModelExploreSupportedMeasureType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
