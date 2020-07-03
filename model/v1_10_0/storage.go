// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Storage storage
//
// swagger:model Storage
type Storage struct {

	// Free volume size.
	Free int64 `json:"free,omitempty"`

	// Total volume size.
	Total int64 `json:"total,omitempty"`
}

// Validate validates this storage
func (m *Storage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Storage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Storage) UnmarshalBinary(b []byte) error {
	var res Storage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
