// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Label label
//
// swagger:model Label
type Label struct {

	// The color of label.
	Color string `json:"color,omitempty"`

	// The creation time of label.
	CreationTime string `json:"creation_time,omitempty"`

	// The label is deleted or not.
	Deleted bool `json:"deleted,omitempty"`

	// The description of label.
	Description string `json:"description,omitempty"`

	// The ID of label.
	ID int64 `json:"id,omitempty"`

	// The name of label.
	Name string `json:"name,omitempty"`

	// The project ID if the label is a project label.
	ProjectID int64 `json:"project_id,omitempty"`

	// The scope of label, g for global labels and p for project labels.
	Scope string `json:"scope,omitempty"`

	// The update time of label.
	UpdateTime string `json:"update_time,omitempty"`
}

// Validate validates this label
func (m *Label) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Label) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Label) UnmarshalBinary(b []byte) error {
	var res Label
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
