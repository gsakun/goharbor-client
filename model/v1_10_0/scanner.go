// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Scanner scanner
//
// swagger:model Scanner
type Scanner struct {

	// Name of the scanner
	Name string `json:"name,omitempty"`

	// Name of the scanner provider
	Vendor string `json:"vendor,omitempty"`

	// Version of the scanner adapter
	Version string `json:"version,omitempty"`
}

// Validate validates this scanner
func (m *Scanner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Scanner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Scanner) UnmarshalBinary(b []byte) error {
	var res Scanner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
