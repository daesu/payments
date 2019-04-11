// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ListMetadata List Metadata
// swagger:model list-metadata
type ListMetadata struct {

	// offset
	Offset int64 `json:"Offset"`

	// page size
	PageSize int64 `json:"PageSize"`

	// total size
	TotalSize int64 `json:"TotalSize"`
}

// Validate validates this list metadata
func (m *ListMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListMetadata) UnmarshalBinary(b []byte) error {
	var res ListMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
