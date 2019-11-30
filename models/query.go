// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Query query
// swagger:model Query
type Query struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Client Id: used to generate shortened explore URLs. If set by client, must be a unique 22 character alphanumeric string. Otherwise one will be generated.
	ClientID string `json:"client_id,omitempty"`

	// Column Limit
	ColumnLimit string `json:"column_limit,omitempty"`

	// Dynamic Fields
	DynamicFields string `json:"dynamic_fields,omitempty"`

	// Expanded Share Url
	// Read Only: true
	ExpandedShareURL string `json:"expanded_share_url,omitempty"`

	// Fields
	Fields []string `json:"fields"`

	// Fill Fields
	FillFields []string `json:"fill_fields"`

	// The filter_config represents the state of the filter UI on the explore page for a given query. When running a query via the Looker UI, this parameter takes precedence over "filters". When creating a query or modifying an existing query, "filter_config" should be set to null. Setting it to any other value could cause unexpected filtering behavior. The format should be considered opaque.
	FilterConfig map[string]string `json:"filter_config,omitempty"`

	// Filter Expression
	FilterExpression string `json:"filter_expression,omitempty"`

	// Filters
	Filters map[string]string `json:"filters,omitempty"`

	// Has Table Calculations
	// Read Only: true
	HasTableCalculations *bool `json:"has_table_calculations,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Limit
	Limit string `json:"limit,omitempty"`

	// Model
	// Required: true
	Model *string `json:"model"`

	// Pivots
	Pivots []string `json:"pivots"`

	// Query Timezone
	QueryTimezone string `json:"query_timezone,omitempty"`

	// Raw Total
	RowTotal string `json:"row_total,omitempty"`

	// Runtime
	Runtime float64 `json:"runtime,omitempty"`

	// Share Url
	// Read Only: true
	ShareURL string `json:"share_url,omitempty"`

	// Slug
	// Read Only: true
	Slug string `json:"slug,omitempty"`

	// Sorting for the query results. Use the format `["view.field", ...]` to sort on fields in ascending order. Use the format `["view.field desc", ...]` to sort on fields in descending order. Use `["__UNSORTED__"]` (2 underscores before and after) to disable sorting entirely. Empty sorts `[]` will trigger a default sort.
	Sorts []string `json:"sorts"`

	// Fields on which to run subtotals
	Subtotals []string `json:"subtotals"`

	// Total
	Total bool `json:"total,omitempty"`

	// Expanded Url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// Explore Name
	// Required: true
	View *string `json:"view"`

	// Visualization configuration properties. These properties are typically opaque and differ based on the type of visualization used. There is no specified set of allowed keys. The values can be any type supported by JSON. A "type" key with a string value is often present, and is used by Looker to determine which visualization to present. Visualizations ignore unknown vis_config properties.
	VisConfig map[string]string `json:"vis_config,omitempty"`

	// Visible UI Sections
	VisibleUISections string `json:"visible_ui_sections,omitempty"`
}

// Validate validates this query
func (m *Query) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateView(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Query) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *Query) validateView(formats strfmt.Registry) error {

	if err := validate.Required("view", "body", m.View); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Query) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Query) UnmarshalBinary(b []byte) error {
	var res Query
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
