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

// LookmlModelExploreField lookml model explore field
// swagger:model LookmlModelExploreField
type LookmlModelExploreField struct {

	// The appropriate horizontal text alignment the values of this field shoud be displayed in. Valid values are: "left", "right".
	// Read Only: true
	Align string `json:"align,omitempty"`

	// Whether it's possible to filter on this field.
	// Read Only: true
	CanFilter *bool `json:"can_filter,omitempty"`

	// Whether this field can be time filtered.
	// Read Only: true
	CanTimeFilter *bool `json:"can_time_filter,omitempty"`

	// Field category Valid values are: "parameter", "filter", "measure", "dimension".
	// Read Only: true
	Category string `json:"category,omitempty"`

	// The default value that this field uses when filtering. Null if there is no default value.
	// Read Only: true
	DefaultFilterValue string `json:"default_filter_value,omitempty"`

	// Description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Whether this field was specified in "dynamic_fields" and is not part of the model.
	// Read Only: true
	Dynamic *bool `json:"dynamic,omitempty"`

	// An array enumerating all the possible values that this field can contain. When null, there is no limit to the set of possible values this field can contain.
	// Read Only: true
	Enumerations []*LookmlModelExploreFieldEnumeration `json:"enumerations"`

	// An error message indicating a problem with the definition of this field. If there are no errors, this will be null.
	// Read Only: true
	Error string `json:"error,omitempty"`

	// A label creating a grouping of fields. All fields with this label should be presented together when displayed in a UI.
	// Read Only: true
	FieldGroupLabel string `json:"field_group_label,omitempty"`

	// When presented in a field group via field_group_label, a shorter name of the field to be displayed in that context.
	// Read Only: true
	FieldGroupVariant string `json:"field_group_variant,omitempty"`

	// The style of dimension fill that is possible for this field. Null if no dimension fill is possible. Valid values are: "enumeration", "range".
	// Read Only: true
	FillStyle string `json:"fill_style,omitempty"`

	// An offset (in months) from the calendar start month to the fiscal start month defined in the LookML model this field belongs to.
	// Read Only: true
	FiscalMonthOffset int64 `json:"fiscal_month_offset,omitempty"`

	// Whether this field has a set of allowed_values specified in LookML.
	// Read Only: true
	HasAllowedValues *bool `json:"has_allowed_values,omitempty"`

	// Whether this field should be hidden from the user interface.
	// Read Only: true
	Hidden *bool `json:"hidden,omitempty"`

	// Whether this field is a filter.
	// Read Only: true
	IsFilter *bool `json:"is_filter,omitempty"`

	// Whether this field represents a fiscal time value.
	// Read Only: true
	IsFiscal *bool `json:"is_fiscal,omitempty"`

	// Whether this field is of a type that represents a numeric value.
	// Read Only: true
	IsNumeric *bool `json:"is_numeric,omitempty"`

	// Whether this field is of a type that represents a time value.
	// Read Only: true
	IsTimeframe *bool `json:"is_timeframe,omitempty"`

	// Fully-qualified human-readable label of the field.
	// Read Only: true
	Label string `json:"label,omitempty"`

	// The name of the parameter that will provide a parameterized label for this field, if available in the current context.
	// Read Only: true
	LabelFromParameter string `json:"label_from_parameter,omitempty"`

	// The human-readable label of the field, without the view label.
	// Read Only: true
	LabelShort string `json:"label_short,omitempty"`

	// A URL linking to the definition of this field in the LookML IDE.
	// Read Only: true
	LookmlLink string `json:"lookml_link,omitempty"`

	// If applicable, a map layer this field is associated with.
	// Read Only: true
	MapLayer *LookmlModelExploreFieldMapLayer `json:"map_layer,omitempty"`

	// Whether this field is a measure.
	// Read Only: true
	Measure *bool `json:"measure,omitempty"`

	// Fully-qualified name of the field.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Whether this field is a parameter.
	// Read Only: true
	Parameter *bool `json:"parameter,omitempty"`

	// Whether this field can be removed from a query.
	// Read Only: true
	Permanent *bool `json:"permanent,omitempty"`

	// Whether or not the field represents a primary key.
	// Read Only: true
	PrimaryKey *bool `json:"primary_key,omitempty"`

	// The name of the project this field is defined in.
	// Read Only: true
	ProjectName string `json:"project_name,omitempty"`

	// When true, it's not possible to re-sort this field's values without re-running the SQL query, due to database logic that affects the sort.
	// Read Only: true
	RequiresRefreshOnSort *bool `json:"requires_refresh_on_sort,omitempty"`

	// The LookML scope this field belongs to. The scope is typically the field's view.
	// Read Only: true
	Scope string `json:"scope,omitempty"`

	// Whether this field can be sorted.
	// Read Only: true
	Sortable *bool `json:"sortable,omitempty"`

	// The path portion of source_file_path.
	// Read Only: true
	SourceFile string `json:"source_file,omitempty"`

	// The fully-qualified path of the project file this field is defined in.
	// Read Only: true
	SourceFilePath string `json:"source_file_path,omitempty"`

	// SQL expression as defined in the LookML model. The SQL syntax shown here is a representation intended for auditability, and is not neccessarily an exact match for what will ultimately be run in the database. It may contain special LookML syntax or annotations that are not valid SQL. This will be null if the current user does not have the see_lookml permission for the field's model.
	// Read Only: true
	SQL string `json:"sql,omitempty"`

	// An array of conditions and values that make up a SQL Case expression, as defined in the LookML model. The SQL syntax shown here is a representation intended for auditability, and is not neccessarily an exact match for what will ultimately be run in the database. It may contain special LookML syntax or annotations that are not valid SQL. This will be null if the current user does not have the see_lookml permission for the field's model.
	// Read Only: true
	SQLCase []*LookmlModelExploreFieldSQLCase `json:"sql_case"`

	// The name of the dimension to base suggest queries from.
	// Read Only: true
	SuggestDimension string `json:"suggest_dimension,omitempty"`

	// The name of the explore to base suggest queries from.
	// Read Only: true
	SuggestExplore string `json:"suggest_explore,omitempty"`

	// Whether or not suggestions are possible for this field.
	// Read Only: true
	Suggestable *bool `json:"suggestable,omitempty"`

	// If available, a list of suggestions for this field. For most fields, a suggest query is a more appropriate way to get an up-to-date list of suggestions. Or use enumerations to list all the possible values.
	// Read Only: true
	Suggestions []string `json:"suggestions"`

	// An array of arbitrary string tags provided in the model for this field.
	// Read Only: true
	Tags []string `json:"tags"`

	// Details on the time interval this field represents, if it is_timeframe.
	// Read Only: true
	TimeInterval *LookmlModelExploreFieldTimeInterval `json:"time_interval,omitempty"`

	// The LookML type of the field.
	// Read Only: true
	Type string `json:"type,omitempty"`

	// An array of user attribute types that are allowed to be used in filters on this field. Valid values are: "advanced_filter_string", "advanced_filter_number", "advanced_filter_datetime", "string", "number", "datetime", "yesno", "zipcode".
	// Read Only: true
	UserAttributeFilterTypes []string `json:"user_attribute_filter_types"`

	// If specified, the LookML value format string for formatting values of this field.
	// Read Only: true
	ValueFormat string `json:"value_format,omitempty"`

	// The name of the view this field belongs to.
	// Read Only: true
	View string `json:"view,omitempty"`

	// The human-readable label of the view the field belongs to.
	// Read Only: true
	ViewLabel string `json:"view_label,omitempty"`

	// The name of the starting day of the week. Valid values are: "monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday".
	// Read Only: true
	WeekStartDay string `json:"week_start_day,omitempty"`
}

// Validate validates this lookml model explore field
func (m *LookmlModelExploreField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnumerations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapLayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSQLCase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LookmlModelExploreField) validateEnumerations(formats strfmt.Registry) error {

	if swag.IsZero(m.Enumerations) { // not required
		return nil
	}

	for i := 0; i < len(m.Enumerations); i++ {
		if swag.IsZero(m.Enumerations[i]) { // not required
			continue
		}

		if m.Enumerations[i] != nil {
			if err := m.Enumerations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("enumerations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LookmlModelExploreField) validateMapLayer(formats strfmt.Registry) error {

	if swag.IsZero(m.MapLayer) { // not required
		return nil
	}

	if m.MapLayer != nil {
		if err := m.MapLayer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("map_layer")
			}
			return err
		}
	}

	return nil
}

func (m *LookmlModelExploreField) validateSQLCase(formats strfmt.Registry) error {

	if swag.IsZero(m.SQLCase) { // not required
		return nil
	}

	for i := 0; i < len(m.SQLCase); i++ {
		if swag.IsZero(m.SQLCase[i]) { // not required
			continue
		}

		if m.SQLCase[i] != nil {
			if err := m.SQLCase[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sql_case" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LookmlModelExploreField) validateTimeInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeInterval) { // not required
		return nil
	}

	if m.TimeInterval != nil {
		if err := m.TimeInterval.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time_interval")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LookmlModelExploreField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookmlModelExploreField) UnmarshalBinary(b []byte) error {
	var res LookmlModelExploreField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
