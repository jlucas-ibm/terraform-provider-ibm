// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RouteNextHopIP The IP address of the next hop to which to route packets
// swagger:model RouteNextHopIP
type RouteNextHopIP struct {
	IP
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RouteNextHopIP) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 IP
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.IP = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RouteNextHopIP) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.IP)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route next hop IP
func (m *RouteNextHopIP) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RouteNextHopIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RouteNextHopIP) UnmarshalBinary(b []byte) error {
	var res RouteNextHopIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
