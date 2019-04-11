package test

import (
	"fmt"
	"testing"

	"github.com/daesu/payments/gen/models"
	"github.com/imroc/req"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreatePaymentEndpoint(t *testing.T) {
	Convey("Given API is running", t, func() {

		var amount = 102.12
		var currency = "EUR"
		var beneficiaryName = "Bernard Brown"
		var debtorName = "Jon Snow"

		var payment models.Payment
		attributes := models.PaymentAttribute{}
		beneficiary := models.CustomerAccount{}
		debtor := models.CustomerAccount{}

		attributes.BeneficiaryParty = &beneficiary
		attributes.BeneficiaryParty = &debtor
		payment.Attributes = &attributes

		Convey("When user requests payments/<payment-id> endpoint with valid POST data", func() {

			url := "http://localhost:8080/v1/payments"

			json := fmt.Sprintf(`{
				"amount": %f,
				"beneficiary": {
					"account_name": "Brown B",
					"account_number": "1234567",
					"account_number_code": "BBAN",
					"account_type": "0",
					"address": "Oakwood, Andor",
					"bank_id": 10090,
					"bank_id_code": "IRONB",
					"name": "%s"
				},
				"currency": "%s",
				"debtor": {
					"account_name": "J Snow",
					"account_number": "87654323",
					"account_number_code": "CDAN",
					"account_type": "0",
					"address": "123 Wall st, London",
					"bank_id": 10042,
					"bank_id_code": "IRONB",
					"name": "%s"
				}
			}`, amount, beneficiaryName, currency, debtorName)

			header := req.Header{
				"Content-Type": "application/json",
			}

			resp, err := req.Post(url, header, json)

			if err != nil {
				t.Error("Response: ", resp.String())
				t.Fail()
			}

			Convey("It will get 201 status", func() {
				So(resp.Response().StatusCode, ShouldEqual, 201)
			})

			err = resp.ToJSON(&payment)
			if err != nil {
				t.Error("Error: ", err.Error())
			}

			Convey("It will get the specified payment values", func() {
				So(payment.Attributes.Amount, ShouldEqual, amount)
				So(payment.Attributes.Currency, ShouldEqual, currency)
			})

			Convey("It will get the payment beneficiary", func() {
				name := payment.Attributes.BeneficiaryParty.Name
				So(name, ShouldEqual, beneficiaryName)
			})

			Convey("It will get the payment debtor", func() {
				name := payment.Attributes.DebtorParty.Name
				So(name, ShouldEqual, debtorName)
			})

		})

		Convey("When user requests payments/<payment-id> endpoint with missing required field", func() {

			// amount should not be string
			var amount = 102.12
			var currency = "EUR"
			var beneficiaryName = "Bernard Brown"

			url := "http://localhost:8080/v1/payments"

			json := fmt.Sprintf(`{
				"amount": %f,
				"beneficiary": {
					"account_name": "Brown B",
					"account_number": "1234567",
					"account_number_code": "BBAN",
					"account_type": "0",
					"address": "Oakwood, Andor",
					"bank_id": 10090,
					"bank_id_code": "IRONB",
					"name": "%s"
				},
				"currency": "%s"
			}`, amount, beneficiaryName, currency)

			header := req.Header{
				"Content-Type": "application/json",
			}

			resp, err := req.Post(url, header, json)

			if err != nil {
				t.Error("Response: ", resp.String())
				t.Fail()
			}

			Convey("It will get 422 status", func() {
				So(resp.Response().StatusCode, ShouldEqual, 422)
			})

		})

		Convey("When user requests payments/<payment-id> endpoint with invalid data", func() {

			url := "http://localhost:8080/v1/payments"

			json := `{ invalid post data }`

			header := req.Header{
				"Content-Type": "application/json",
			}

			resp, err := req.Post(url, header, json)

			if err != nil {
				t.Error("Response: ", resp.String())
				t.Fail()
			}

			Convey("It will get 400 status", func() {
				So(resp.Response().StatusCode, ShouldEqual, 400)
			})

		})

	})
}
