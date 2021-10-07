// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReplicationFilter replication filter
//
// swagger:model ReplicationFilter
type ReplicationFilter struct {

	// The replication policy filter type.
	Type string `json:"type,omitempty"`

	// The value of replication policy filter.
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this replication filter
func (m *ReplicationFilter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationFilter) UnmarshalBinary(b []byte) error {
	var res ReplicationFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
