// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LookmlModelExplore lookml model explore
// swagger:model LookmlModelExplore
type LookmlModelExplore struct {

	// (DEPRECATED) Array of access filter field names
	// Read Only: true
	AccessFilterFields []string `json:"access_filter_fields"`

	// Access filters
	// Read Only: true
	AccessFilters []*LookmlModelExploreAccessFilter `json:"access_filters"`

	// Aliases
	// Read Only: true
	Aliases []*LookmlModelExploreAlias `json:"aliases"`

	// Always filter
	// Read Only: true
	AlwaysFilter []*LookmlModelExploreAlwaysFilter `json:"always_filter"`

	// Can Explain
	// Read Only: true
	CanExplain *bool `json:"can_explain,omitempty"`

	// Can pivot in the DB
	// Read Only: true
	CanPivotInDb *bool `json:"can_pivot_in_db,omitempty"`

	// Can Save
	// Read Only: true
	CanSave *bool `json:"can_save,omitempty"`

	// Can use subtotals
	// Read Only: true
	CanSubtotal *bool `json:"can_subtotal,omitempty"`

	// Can Total
	// Read Only: true
	CanTotal *bool `json:"can_total,omitempty"`

	// Conditionally filter
	// Read Only: true
	ConditionallyFilter []*LookmlModelExploreConditionallyFilter `json:"conditionally_filter"`

	// Connection name
	// Read Only: true
	ConnectionName string `json:"connection_name,omitempty"`

	// Description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Errors
	// Read Only: true
	Errors []*LookmlModelExploreError `json:"errors"`

	// Fields
	// Read Only: true
	Fields *LookmlModelExploreFieldset `json:"fields,omitempty"`

	// List of model source files
	// Read Only: true
	Files []string `json:"files"`

	// Label used to group explores in the navigation menus
	// Read Only: true
	GroupLabel string `json:"group_label,omitempty"`

	// Has timezone support
	// Read Only: true
	HasTimezoneSupport *bool `json:"has_timezone_support,omitempty"`

	// Is hidden
	// Read Only: true
	Hidden *bool `json:"hidden,omitempty"`

	// Fully qualified name model plus explore name
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Array of index fields
	// Read Only: true
	IndexFields []string `json:"index_fields"`

	// Views joined into this explore
	// Read Only: true
	Joins []*LookmlModelExploreJoins `json:"joins"`

	// Label
	// Read Only: true
	Label string `json:"label,omitempty"`

	// Name of model
	// Read Only: true
	ModelName string `json:"model_name,omitempty"`

	// Explore name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// How nulls are sorted, possible values are "low", "high", "first" and "last"
	// Read Only: true
	NullSortTreatment string `json:"null_sort_treatment,omitempty"`

	// Name of project
	// Read Only: true
	ProjectName string `json:"project_name,omitempty"`

	// Scopes
	// Read Only: true
	Scopes []string `json:"scopes"`

	// Sets
	// Read Only: true
	Sets []*LookmlModelExploreSet `json:"sets"`

	// Primary source_file file
	// Read Only: true
	SourceFile string `json:"source_file,omitempty"`

	// A sql_table_name expression that defines what sql table the view/explore maps onto. Example: "prod_orders2 AS orders" in a view named orders.
	// Read Only: true
	SQLTableName string `json:"sql_table_name,omitempty"`

	// An array of items describing which custom measure types are supported for creating a custom measure 'baed_on' each possible dimension type.
	// Read Only: true
	SupportedMeasureTypes []*LookmlModelExploreSupportedMeasureType `json:"supported_measure_types"`

	// Cost estimates supported
	// Read Only: true
	SupportsCostEstimate *bool `json:"supports_cost_estimate,omitempty"`

	// An array of arbitrary string tags provided in the model for this explore.
	// Read Only: true
	Tags []string `json:"tags"`

	// Name of view
	// Read Only: true
	ViewName string `json:"view_name,omitempty"`
}

// Validate validates this lookml model explore
func (m *LookmlModelExplore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAliases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlwaysFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditionallyFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJoins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedMeasureTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LookmlModelExplore) validateAccessFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.AccessFilters); i++ {
		if swag.IsZero(m.AccessFilters[i]) { // not required
			continue
		}

		if m.AccessFilters[i] != nil {
			if err := m.AccessFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("access_filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LookmlModelExplore) validateAliases(formats strfmt.Registry) error {

	if swag.IsZero(m.Aliases) { // not required
		return nil
	}

	for i := 0; i < len(m.Aliases); i++ {
		if swag.IsZero(m.Aliases[i]) { // not required
			continue
		}

		if m.Aliases[i] != nil {
			if err := m.Aliases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LookmlModelExplore) validateAlwaysFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.AlwaysFilter) { // not required
		return nil
	}

	for i := 0; i < len(m.AlwaysFilter); i++ {
		if swag.IsZero(m.AlwaysFilter[i]) { // not required
			continue
		}

		if m.AlwaysFilter[i] != nil {
			if err := m.AlwaysFilter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("always_filter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LookmlModelExplore) validateConditionallyFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.ConditionallyFilter) { // not required
		return nil
	}

	for i := 0; i < len(m.ConditionallyFilter); i++ {
		if swag.IsZero(m.ConditionallyFilter[i]) { // not required
			continue
		}

		if m.ConditionallyFilter[i] != nil {
			if err := m.ConditionallyFilter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditionally_filter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LookmlModelExplore) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LookmlModelExplore) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	if m.Fields != nil {
		if err := m.Fields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fields")
			}
			return err
		}
	}

	return nil
}

func (m *LookmlModelExplore) validateJoins(formats strfmt.Registry) error {

	if swag.IsZero(m.Joins) { // not required
		return nil
	}

	for i := 0; i < len(m.Joins); i++ {
		if swag.IsZero(m.Joins[i]) { // not required
			continue
		}

		if m.Joins[i] != nil {
			if err := m.Joins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("joins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LookmlModelExplore) validateSets(formats strfmt.Registry) error {

	if swag.IsZero(m.Sets) { // not required
		return nil
	}

	for i := 0; i < len(m.Sets); i++ {
		if swag.IsZero(m.Sets[i]) { // not required
			continue
		}

		if m.Sets[i] != nil {
			if err := m.Sets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LookmlModelExplore) validateSupportedMeasureTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportedMeasureTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportedMeasureTypes); i++ {
		if swag.IsZero(m.SupportedMeasureTypes[i]) { // not required
			continue
		}

		if m.SupportedMeasureTypes[i] != nil {
			if err := m.SupportedMeasureTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supported_measure_types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LookmlModelExplore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookmlModelExplore) UnmarshalBinary(b []byte) error {
	var res LookmlModelExplore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
