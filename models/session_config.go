// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SessionConfig session config
// swagger:model SessionConfig
type SessionConfig struct {

	// Allow users to have persistent sessions when they login
	AllowPersistentSessions bool `json:"allow_persistent_sessions,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Number of minutes for user sessions.  Must be between 5 and 43200
	SessionMinutes int64 `json:"session_minutes,omitempty"`

	// Allow users to have an unbounded number of concurrent sessions (otherwise, users will be limited to only one session at a time).
	UnlimitedSessionsPerUser bool `json:"unlimited_sessions_per_user,omitempty"`

	// Enforce session logout for sessions that are inactive for 15 minutes.
	UseInactivityBasedLogout bool `json:"use_inactivity_based_logout,omitempty"`
}

// Validate validates this session config
func (m *SessionConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SessionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SessionConfig) UnmarshalBinary(b []byte) error {
	var res SessionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
