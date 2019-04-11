package payment

import "github.com/go-openapi/strfmt"

// PaymentConstruct ...
type PaymentConstruct struct {
	SchemePaymentTypeID string      `json:"scheme_payment_type_id"`
	OrganisationID      string      `json:"organisation_id"`
	Type                string      `json:"type"`
	BeneficaryID        string      `json:"beneficary_id"`
	DebtorID            string      `json:"debtor_id"`
	CurrencyID          string      `json:"currency_id"`
	Amount              float64     `json:"amount"`
	EndToEndReference   string      `json:"end_to_end_reference,omitempty"`
	NumericReference    float64     `json:"numeric_reference,omitempty"`
	Reference           string      `json:"reference,omitempty"`
	PaymentPurpose      string      `json:"payment_purpose,omitempty"`
	ProcessingDate      strfmt.Date `json:"processing_date,omitempty"`
}
