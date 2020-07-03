// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Role role
//
// swagger:model Role
type Role struct {

	// Description of permissions for the role.
	RoleCode string `json:"role_code,omitempty"`

	// ID in table.
	RoleID int32 `json:"role_id,omitempty"`

	// role mask
	RoleMask string `json:"role_mask,omitempty"`

	// Name the the role.
	RoleName string `json:"role_name,omitempty"`
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
