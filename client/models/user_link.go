// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserLink user link
//
// swagger:model userLink
type UserLink struct {

	// ID
	ID string `json:"ID,omitempty"`

	// group type
	GroupType string `json:"groupType,omitempty"`

	// personal data
	PersonalData bool `json:"personalData,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`

	// uri
	URI string `json:"uri,omitempty"`
}

// Validate validates this user link
func (m *UserLink) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user link based on context it is used
func (m *UserLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserLink) UnmarshalBinary(b []byte) error {
	var res UserLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}