// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RepoSignature repo signature
//
// swagger:model RepoSignature
type RepoSignature struct {

	// The JSON object of the hash of the image.
	Hashes interface{} `json:"hashes,omitempty"`

	// The tag of image.
	Tag string `json:"tag,omitempty"`
}

// Validate validates this repo signature
func (m *RepoSignature) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RepoSignature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepoSignature) UnmarshalBinary(b []byte) error {
	var res RepoSignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
