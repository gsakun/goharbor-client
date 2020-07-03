// Code generated by go-swagger; DO NOT EDIT.

package v1_10_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebhookLastTrigger The webhook policy and last trigger time group by event type.
//
// swagger:model WebhookLastTrigger
type WebhookLastTrigger struct {

	// The creation time of webhook policy.
	CreationTime string `json:"creation_time,omitempty"`

	// Whether or not the webhook policy enabled.
	Enabled bool `json:"enabled,omitempty"`

	// The webhook event type.
	EventType string `json:"event_type,omitempty"`

	// The last trigger time of webhook policy.
	LastTriggerTime string `json:"last_trigger_time,omitempty"`
}

// Validate validates this webhook last trigger
func (m *WebhookLastTrigger) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookLastTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookLastTrigger) UnmarshalBinary(b []byte) error {
	var res WebhookLastTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
