// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LdapFailedImportUsers ldap failed import users
//
// swagger:model LdapFailedImportUsers
type LdapFailedImportUsers struct {

	// fail reason.
	Error string `json:"error,omitempty"`

	// the uid can't add to system.
	LdapUID string `json:"ldap_uid,omitempty"`
}

// Validate validates this ldap failed import users
func (m *LdapFailedImportUsers) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapFailedImportUsers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapFailedImportUsers) UnmarshalBinary(b []byte) error {
	var res LdapFailedImportUsers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
