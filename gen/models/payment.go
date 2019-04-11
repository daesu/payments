// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Payment Payment
//
// properties for payment
// swagger:model payment
type Payment struct {

	// attributes
	Attributes *PaymentAttribute `json:"attributes,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// organisation id
	OrganisationID string `json:"organisation_id,omitempty"`

	// type
	// Enum: [Payment DirectDebit Payment, DirectDebit' Mandate]
	Type string `json:"type,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this payment
func (m *Payment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Payment) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

var paymentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Payment","DirectDebit","Payment, DirectDebit'","Mandate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentTypeTypePropEnum = append(paymentTypeTypePropEnum, v)
	}
}

const (

	// PaymentTypePayment captures enum value "Payment"
	PaymentTypePayment string = "Payment"

	// PaymentTypeDirectDebit captures enum value "DirectDebit"
	PaymentTypeDirectDebit string = "DirectDebit"

	// PaymentTypePaymentDirectDebit captures enum value "Payment, DirectDebit'"
	PaymentTypePaymentDirectDebit string = "Payment, DirectDebit'"

	// PaymentTypeMandate captures enum value "Mandate"
	PaymentTypeMandate string = "Mandate"
)

// prop value enum
func (m *Payment) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Payment) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Payment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Payment) UnmarshalBinary(b []byte) error {
	var res Payment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}