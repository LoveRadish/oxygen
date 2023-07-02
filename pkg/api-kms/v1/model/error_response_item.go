// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorResponseItem error response item
//
// swagger:model ErrorResponseItem
type ErrorResponseItem struct {

	// Error field
	// Example: username
	Field string `json:"field,omitempty"`

	// Error Message
	// Example: You are unauthenticated
	Message string `json:"message,omitempty"`
}

// Validate validates this error response item
func (m *ErrorResponseItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error response item based on context it is used
func (m *ErrorResponseItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponseItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponseItem) UnmarshalBinary(b []byte) error {
	var res ErrorResponseItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}