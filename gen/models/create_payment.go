// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreatePayment createPayment
//
// properties to create a payment
// swagger:model create-payment
type CreatePayment struct {

	// amount
	// Required: true
	Amount *float64 `json:"amount"`

	// beneficiary
	// Required: true
	Beneficiary *CreateCustomerAccount `json:"beneficiary"`

	// currency
	// Required: true
	Currency *string `json:"currency"`

	// debtor
	// Required: true
	Debtor *CreateCustomerAccount `json:"debtor"`

	// end to end reference
	EndToEndReference *string `json:"end_to_end_reference,omitempty"`

	// payment purpose
	PaymentPurpose string `json:"payment_purpose,omitempty"`

	// reference
	Reference string `json:"reference,omitempty"`
}

// Validate validates this create payment
func (m *CreatePayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePayment) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *CreatePayment) validateBeneficiary(formats strfmt.Registry) error {

	if err := validate.Required("beneficiary", "body", m.Beneficiary); err != nil {
		return err
	}

	if m.Beneficiary != nil {
		if err := m.Beneficiary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beneficiary")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePayment) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *CreatePayment) validateDebtor(formats strfmt.Registry) error {

	if err := validate.Required("debtor", "body", m.Debtor); err != nil {
		return err
	}

	if m.Debtor != nil {
		if err := m.Debtor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debtor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreatePayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePayment) UnmarshalBinary(b []byte) error {
	var res CreatePayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
