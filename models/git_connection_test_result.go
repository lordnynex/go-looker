// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GitConnectionTestResult git connection test result
// swagger:model GitConnectionTestResult
type GitConnectionTestResult struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// A short string, uniquely naming this test
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Additional data from the test
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Either "pass" or "fail"
	// Read Only: true
	Status string `json:"status,omitempty"`
}

// Validate validates this git connection test result
func (m *GitConnectionTestResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitConnectionTestResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitConnectionTestResult) UnmarshalBinary(b []byte) error {
	var res GitConnectionTestResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}