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

// ValidationErrorDetail validation error detail
// swagger:model ValidationErrorDetail
type ValidationErrorDetail struct {

	// Error code
	// Read Only: true
	Code string `json:"code,omitempty"`

	// Documentation link
	// Required: true
	// Read Only: true
	// Format: uri
	DocumentationURL strfmt.URI `json:"documentation_url"`

	// Field with error
	// Read Only: true
	Field string `json:"field,omitempty"`

	// Error info message
	// Read Only: true
	Message string `json:"message,omitempty"`
}

// Validate validates this validation error detail
func (m *ValidationErrorDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentationURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValidationErrorDetail) validateDocumentationURL(formats strfmt.Registry) error {

	if err := validate.Required("documentation_url", "body", strfmt.URI(m.DocumentationURL)); err != nil {
		return err
	}

	if err := validate.FormatOf("documentation_url", "body", "uri", m.DocumentationURL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ValidationErrorDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValidationErrorDetail) UnmarshalBinary(b []byte) error {
	var res ValidationErrorDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}