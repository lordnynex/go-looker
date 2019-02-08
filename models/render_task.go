// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RenderTask render task
// swagger:model RenderTask
type RenderTask struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Date/Time render task was created
	// Read Only: true
	CreatedAt string `json:"created_at,omitempty"`

	// Filter values to apply to the dashboard queries, in URL query format
	// Read Only: true
	DashboardFilters string `json:"dashboard_filters,omitempty"`

	// Id of dashboard to render
	// Read Only: true
	DashboardID int64 `json:"dashboard_id,omitempty"`

	// Dashboard layout style: single_column or tiled
	// Read Only: true
	DashboardStyle string `json:"dashboard_style,omitempty"`

	// Date/Time render task was completed
	// Read Only: true
	FinalizedAt string `json:"finalized_at,omitempty"`

	// Output height in pixels. Flowed layouts may ignore this value.
	// Read Only: true
	Height int64 `json:"height,omitempty"`

	// Id of this render task
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Id of look to render
	// Read Only: true
	LookID int64 `json:"look_id,omitempty"`

	// Id of lookml dashboard to render
	// Read Only: true
	LookmlDashboardID string `json:"lookml_dashboard_id,omitempty"`

	// Id of query to render
	// Read Only: true
	QueryID int64 `json:"query_id,omitempty"`

	// Number of seconds elapsed running queries
	// Read Only: true
	QueryRuntime float64 `json:"query_runtime,omitempty"`

	// Number of seconds elapsed rendering data
	// Read Only: true
	RenderRuntime float64 `json:"render_runtime,omitempty"`

	// Output format: pdf, png, or jpg
	// Read Only: true
	ResultFormat string `json:"result_format,omitempty"`

	// Total seconds elapsed for render task
	// Read Only: true
	Runtime float64 `json:"runtime,omitempty"`

	// Render task status: enqueued_for_query, querying, enqueued_for_render, rendering, success, failure
	// Read Only: true
	Status string `json:"status,omitempty"`

	// Additional information about the current status
	// Read Only: true
	StatusDetail string `json:"status_detail,omitempty"`

	// The user account permissions in which the render task will execute
	// Read Only: true
	UserID int64 `json:"user_id,omitempty"`

	// Output width in pixels
	// Read Only: true
	Width int64 `json:"width,omitempty"`
}

// Validate validates this render task
func (m *RenderTask) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RenderTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RenderTask) UnmarshalBinary(b []byte) error {
	var res RenderTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
