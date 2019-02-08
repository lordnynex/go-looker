// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User user
// swagger:model User
type User struct {

	// Model access filters.
	// Read Only: true
	AccessFilters []*AccessFilter `json:"access_filters"`

	// URL for the avatar image (may be generic)
	// Read Only: true
	// Format: uri
	AvatarURL strfmt.URI `json:"avatar_url,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// API user credentials. NO LONGER SUPPORTED.
	// Read Only: true
	CredentialsAPI *CredentialsAPI `json:"credentials_api,omitempty"`

	// API 3 credentials
	// Read Only: true
	CredentialsApi3 []*CredentialsApi3 `json:"credentials_api3"`

	// Email/Password login credentials
	// Read Only: true
	CredentialsEmail *CredentialsEmail `json:"credentials_email,omitempty"`

	// Embed credentials
	// Read Only: true
	CredentialsEmbed []*CredentialsEmbed `json:"credentials_embed"`

	// Google auth credentials
	// Read Only: true
	CredentialsGoogle *CredentialsGoogle `json:"credentials_google,omitempty"`

	// LDAP credentials
	// Read Only: true
	CredentialsLdap *CredentialsLDAP `json:"credentials_ldap,omitempty"`

	// LookerOpenID credentials. Used for login by Looker Analysts
	// Read Only: true
	CredentialsLookerOpenid *CredentialsLookerOpenid `json:"credentials_looker_openid,omitempty"`

	// OpenID Connect auth credentials
	// Read Only: true
	CredentialsOidc *CredentialsOIDC `json:"credentials_oidc,omitempty"`

	// Saml auth credentials
	// Read Only: true
	CredentialsSaml *CredentialsSaml `json:"credentials_saml,omitempty"`

	// Two-factor credentials
	// Read Only: true
	CredentialsTotp *CredentialsTotp `json:"credentials_totp,omitempty"`

	// Full name for display (available only if both first_name and last_name are set)
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// EMail address
	// Read Only: true
	Email string `json:"email,omitempty"`

	// (Embed only) ID of user's group space based on the external_group_id optionally specified during embed user login
	// Read Only: true
	EmbedGroupSpaceID int64 `json:"embed_group_space_id,omitempty"`

	// First name
	FirstName string `json:"first_name,omitempty"`

	// Array of ids of the groups for this user
	// Read Only: true
	GroupIds []int64 `json:"group_ids"`

	// ID string for user's home space
	HomeSpaceID string `json:"home_space_id,omitempty"`

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Account has been disabled
	IsDisabled bool `json:"is_disabled,omitempty"`

	// Last name
	LastName string `json:"last_name,omitempty"`

	// User's preferred locale. User locale takes precedence over Looker's system-wide default locale. Locale determines language of display strings and date and numeric formatting in API responses. Locale string must be a 2 letter language code or a combination of language code and region code: 'en' or 'en-US', for example.
	Locale string `json:"locale,omitempty"`

	// Array of strings representing the Looker versions that this user has used (this only goes back as far as '3.54.0')
	// Read Only: true
	LookerVersions []string `json:"looker_versions"`

	// ID of user's personal space
	// Read Only: true
	PersonalSpaceID int64 `json:"personal_space_id,omitempty"`

	// User is identified as an employee of Looker
	// Read Only: true
	PresumedLookerEmployee *bool `json:"presumed_looker_employee,omitempty"`

	// Array of ids of the roles for this user
	// Read Only: true
	RoleIds []int64 `json:"role_ids"`

	// Active sessions
	// Read Only: true
	Sessions []*Session `json:"sessions"`

	// Per user dictionary of undocumented state information owned by the Looker UI.
	UIState map[string]string `json:"ui_state,omitempty"`

	// Link to get this item
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// User is identified as an employee of Looker who has been verified via Looker corporate authentication
	// Read Only: true
	VerifiedLookerEmployee *bool `json:"verified_looker_employee,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvatarURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsApi3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsEmbed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsGoogle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsLdap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsLookerOpenid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsOidc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsSaml(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialsTotp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateAccessFilters(formats strfmt.Registry) error {

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

func (m *User) validateAvatarURL(formats strfmt.Registry) error {

	if swag.IsZero(m.AvatarURL) { // not required
		return nil
	}

	if err := validate.FormatOf("avatar_url", "body", "uri", m.AvatarURL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateCredentialsAPI(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsAPI) { // not required
		return nil
	}

	if m.CredentialsAPI != nil {
		if err := m.CredentialsAPI.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_api")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsApi3(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsApi3) { // not required
		return nil
	}

	for i := 0; i < len(m.CredentialsApi3); i++ {
		if swag.IsZero(m.CredentialsApi3[i]) { // not required
			continue
		}

		if m.CredentialsApi3[i] != nil {
			if err := m.CredentialsApi3[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credentials_api3" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateCredentialsEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsEmail) { // not required
		return nil
	}

	if m.CredentialsEmail != nil {
		if err := m.CredentialsEmail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_email")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsEmbed(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsEmbed) { // not required
		return nil
	}

	for i := 0; i < len(m.CredentialsEmbed); i++ {
		if swag.IsZero(m.CredentialsEmbed[i]) { // not required
			continue
		}

		if m.CredentialsEmbed[i] != nil {
			if err := m.CredentialsEmbed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("credentials_embed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateCredentialsGoogle(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsGoogle) { // not required
		return nil
	}

	if m.CredentialsGoogle != nil {
		if err := m.CredentialsGoogle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_google")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsLdap(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsLdap) { // not required
		return nil
	}

	if m.CredentialsLdap != nil {
		if err := m.CredentialsLdap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_ldap")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsLookerOpenid(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsLookerOpenid) { // not required
		return nil
	}

	if m.CredentialsLookerOpenid != nil {
		if err := m.CredentialsLookerOpenid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_looker_openid")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsOidc(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsOidc) { // not required
		return nil
	}

	if m.CredentialsOidc != nil {
		if err := m.CredentialsOidc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_oidc")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsSaml(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsSaml) { // not required
		return nil
	}

	if m.CredentialsSaml != nil {
		if err := m.CredentialsSaml.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_saml")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateCredentialsTotp(formats strfmt.Registry) error {

	if swag.IsZero(m.CredentialsTotp) { // not required
		return nil
	}

	if m.CredentialsTotp != nil {
		if err := m.CredentialsTotp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials_totp")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateSessions(formats strfmt.Registry) error {

	if swag.IsZero(m.Sessions) { // not required
		return nil
	}

	for i := 0; i < len(m.Sessions); i++ {
		if swag.IsZero(m.Sessions[i]) { // not required
			continue
		}

		if m.Sessions[i] != nil {
			if err := m.Sessions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sessions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
