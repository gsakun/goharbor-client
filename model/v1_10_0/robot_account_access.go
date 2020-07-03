// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RobotAccountAccess robot account access
//
// swagger:model RobotAccountAccess
type RobotAccountAccess struct {

	// the action to resource that perdefined in harbor rbac
	Action string `json:"action,omitempty"`

	// the resource of harbor
	Resource string `json:"resource,omitempty"`
}

// Validate validates this robot account access
func (m *RobotAccountAccess) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RobotAccountAccess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RobotAccountAccess) UnmarshalBinary(b []byte) error {
	var res RobotAccountAccess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
