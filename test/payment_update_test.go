package test

import (
	"fmt"
	"testing"

	"github.com/daesu/payments/gen/models"
	"github.com/imroc/req"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdatePaymentEndpoint(t *testing.T) {
	Convey("Given API is running", t, func() {

		var amount = 43.24
		var currency = "GBP"
		var beneficiaryName = "Sherlock Holmes"
		var debtorName = "James Moriarty"

		var payment models.Payment
		attributes := models.PaymentAttribute{}
		beneficiary := models.CustomerAccount{}
		debtor := models.CustomerAccount{}

		attributes.BeneficiaryParty = &beneficiary
		attributes.BeneficiaryParty = &debtor
		payment.Attributes = &attributes

		Convey("When user requests payments/<payment-id> endpoint with valid PUT data", func() {

			var paymentID = "61b0c143-f1f9-457d-a889-80570b820348"
			url := fmt.Sprintf("http://localhost:8080/v1/payments/%s", paymentID)

			json := fmt.Sprintf(`{
				"amount": %f,
				"beneficiary": {
					"account_name": "Holmes S",
					"account_number": "0986542",
					"account_number_code": "BBAN",
					"account_type": "0",
					"address": "221B Baker Street",
					"bank_id": 10080,
					"bank_id_code": "IRONC",
					"name": "%s"
				},
				"currency": "%s",
				"debtor": {
					"account_name": "Moriarty J",
					"account_number": "43745282",
					"account_number_code": "CDAN",
					"account_type": "0",
					"address": "Canary Warf, London",
					"bank_id": 123477,
					"bank_id_code": "IROND",
					"name": "%s"
				}
			}`, amount, beneficiaryName, currency, debtorName)

			header := req.Header{
				"Content-Type": "application/json",
			}

			resp, err := req.Put(url, header, json)

			if err != nil {
				t.Error("Response: ", resp.String())
				t.Fail()
			}

			Convey("It will get 200 status", func() {
				So(resp.Response().StatusCode, ShouldEqual, 200)
			})

			err = resp.ToJSON(&payment)
			if err != nil {
				t.Error("Error: ", err.Error())
			}

			Convey("It will get the updated payment values", func() {
				So(payment.Attributes.Amount, ShouldEqual, amount)
				So(payment.Attributes.Currency, ShouldEqual, currency)
			})

			Convey("It will get the updated beneficiary", func() {
				name := payment.Attributes.BeneficiaryParty.Name
				So(name, ShouldEqual, beneficiaryName)
			})

			Convey("It will get the updated debtor", func() {
				name := payment.Attributes.DebtorParty.Name
				So(name, ShouldEqual, debtorName)
			})

		})

	})
}
