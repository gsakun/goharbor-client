// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChartInfoEntry The object contains basic chart information
//
// swagger:model ChartInfoEntry
type ChartInfoEntry struct {

	// The created time of chart
	// Required: true
	Created *string `json:"created"`

	// Flag to indicate if the chart is deprecated
	Deprecated bool `json:"deprecated,omitempty"`

	// The home website of chart
	Home string `json:"home,omitempty"`

	// The icon path of chart
	Icon string `json:"icon,omitempty"`

	// latest version of chart
	LatestVersion string `json:"latest_version,omitempty"`

	// Name of chart
	// Required: true
	Name *string `json:"name"`

	// Total count of chart versions
	// Required: true
	TotalVersions *int64 `json:"total_versions"`

	// The created time of chart
	Updated string `json:"updated,omitempty"`
}

// Validate validates this chart info entry
func (m *ChartInfoEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChartInfoEntry) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *ChartInfoEntry) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ChartInfoEntry) validateTotalVersions(formats strfmt.Registry) error {

	if err := validate.Required("total_versions", "body", m.TotalVersions); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChartInfoEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChartInfoEntry) UnmarshalBinary(b []byte) error {
	var res ChartInfoEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
