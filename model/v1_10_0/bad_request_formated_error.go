// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BadRequestFormatedError Bad request
//
// swagger:model BadRequestFormatedError
type BadRequestFormatedError struct {
	ChartAPIError
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BadRequestFormatedError) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ChartAPIError
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ChartAPIError = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BadRequestFormatedError) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ChartAPIError)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this bad request formated error
func (m *BadRequestFormatedError) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ChartAPIError
	if err := m.ChartAPIError.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *BadRequestFormatedError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BadRequestFormatedError) UnmarshalBinary(b []byte) error {
	var res BadRequestFormatedError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
