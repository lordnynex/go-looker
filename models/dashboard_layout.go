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

// DashboardLayout dashboard layout
// swagger:model DashboardLayout
type DashboardLayout struct {

	// Is Active
	Active bool `json:"active,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Column Width
	ColumnWidth int64 `json:"column_width,omitempty"`

	// Id of Dashboard
	DashboardID string `json:"dashboard_id,omitempty"`

	// Components
	// Read Only: true
	DashboardLayoutComponents []*DashboardLayoutComponent `json:"dashboard_layout_components"`

	// Title extracted from the dashboard this layout represents.
	// Read Only: true
	DashboardTitle string `json:"dashboard_title,omitempty"`

	// Whether or not the dashboard layout is deleted.
	// Read Only: true
	Deleted *bool `json:"deleted,omitempty"`

	// Unique Id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Type
	Type string `json:"type,omitempty"`

	// Width
	Width int64 `json:"width,omitempty"`
}

// Validate validates this dashboard layout
func (m *DashboardLayout) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboardLayoutComponents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardLayout) validateDashboardLayoutComponents(formats strfmt.Registry) error {

	if swag.IsZero(m.DashboardLayoutComponents) { // not required
		return nil
	}

	for i := 0; i < len(m.DashboardLayoutComponents); i++ {
		if swag.IsZero(m.DashboardLayoutComponents[i]) { // not required
			continue
		}

		if m.DashboardLayoutComponents[i] != nil {
			if err := m.DashboardLayoutComponents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboard_layout_components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardLayout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardLayout) UnmarshalBinary(b []byte) error {
	var res DashboardLayout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
