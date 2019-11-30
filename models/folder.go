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

// Folder folder
// swagger:model Folder
type Folder struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Children Count
	// Read Only: true
	ChildCount int64 `json:"child_count,omitempty"`

	// Id of content metadata
	// Read Only: true
	ContentMetadataID int64 `json:"content_metadata_id,omitempty"`

	// User Id of Creator
	// Read Only: true
	CreatorID int64 `json:"creator_id,omitempty"`

	// Dashboards
	// Read Only: true
	Dashboards []*DashboardBase `json:"dashboards"`

	// Embedder's Id if this folder was autogenerated as an embedding shared folder via 'external_group_id' in an SSO embed login
	// Read Only: true
	ExternalID string `json:"external_id,omitempty"`

	// Unique Id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Folder is an embed folder
	// Read Only: true
	IsEmbed *bool `json:"is_embed,omitempty"`

	// Folder is the root embed shared folder
	// Read Only: true
	IsEmbedSharedRoot *bool `json:"is_embed_shared_root,omitempty"`

	// Folder is the root embed users folder
	// Read Only: true
	IsEmbedUsersRoot *bool `json:"is_embed_users_root,omitempty"`

	// Folder is a user's personal folder
	// Read Only: true
	IsPersonal *bool `json:"is_personal,omitempty"`

	// Folder is descendant of a user's personal folder
	// Read Only: true
	IsPersonalDescendant *bool `json:"is_personal_descendant,omitempty"`

	// Folder is the root shared folder
	// Read Only: true
	IsSharedRoot *bool `json:"is_shared_root,omitempty"`

	// Folder is the root user folder
	// Read Only: true
	IsUsersRoot *bool `json:"is_users_root,omitempty"`

	// Looks
	// Read Only: true
	Looks []*LookWithDashboards `json:"looks"`

	// Unique Name
	// Required: true
	Name *string `json:"name"`

	// Id of Parent
	// Required: true
	ParentID *string `json:"parent_id"`
}

// Validate validates this folder
func (m *Folder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDashboards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLooks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Folder) validateDashboards(formats strfmt.Registry) error {

	if swag.IsZero(m.Dashboards) { // not required
		return nil
	}

	for i := 0; i < len(m.Dashboards); i++ {
		if swag.IsZero(m.Dashboards[i]) { // not required
			continue
		}

		if m.Dashboards[i] != nil {
			if err := m.Dashboards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Folder) validateLooks(formats strfmt.Registry) error {

	if swag.IsZero(m.Looks) { // not required
		return nil
	}

	for i := 0; i < len(m.Looks); i++ {
		if swag.IsZero(m.Looks[i]) { // not required
			continue
		}

		if m.Looks[i] != nil {
			if err := m.Looks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("looks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Folder) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Folder) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parent_id", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Folder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Folder) UnmarshalBinary(b []byte) error {
	var res Folder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}