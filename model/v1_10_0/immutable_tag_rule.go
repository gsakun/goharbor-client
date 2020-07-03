// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImmutableTagRule immutable tag rule
//
// swagger:model ImmutableTagRule
type ImmutableTagRule struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// project id
	ProjectID int64 `json:"project_id,omitempty"`

	// tag filter
	TagFilter string `json:"tag_filter,omitempty"`
}

// Validate validates this immutable tag rule
func (m *ImmutableTagRule) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImmutableTagRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImmutableTagRule) UnmarshalBinary(b []byte) error {
	var res ImmutableTagRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
