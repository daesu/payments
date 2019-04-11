package test

import (
	"fmt"
	"testing"

	"github.com/daesu/payments/gen/models"
	"github.com/imroc/req"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPaymentEndpoint(t *testing.T) {
	Convey("Given API is running", t, func() {

		// Payment with ID 61b0c143-f1f9-457d-a889-80570b820348 is seeded.
		var paymentID = "61b0c143-f1f9-457d-a889-80570b820348"
		var version = 0
		var organisationID = "c551e221-5fbc-4aa7-84fc-85a6a8a26720"
		var paymentType = "Payment"
		var amount = 102.12
		var currency = "GBP"
		var beneficiaryName = "Lews Therin"
		var debtorName = "Jon Snow"

		var payment models.Payment
		attributes := models.PaymentAttribute{}
		beneficiary := models.CustomerAccount{}
		debtor := models.CustomerAccount{}

		attributes.BeneficiaryParty = &beneficiary
		attributes.BeneficiaryParty = &debtor
		payment.Attributes = &attributes

		Convey("When user requests payments/<payment-id> endpoint with an invalid payment ID", func() {

			var invalidPaymentID = "61b0c143-f1f9-not-real-id"

			url := fmt.Sprintf("http://localhost:8080/v1/payments/%s", invalidPaymentID)

			resp, err := req.Get(url)
			if err != nil {
				t.Error("Response: ", resp.String())
				t.Fail()
			}

			Convey("It will get 404 status", func() {
				So(resp.Response().StatusCode, ShouldEqual, 404)
			})

		})

		Convey("When user requests payments/<payment-id> endpoint with a valid payment ID", func() {

			url := fmt.Sprintf("http://localhost:8080/v1/payments/%s", paymentID)

			resp, err := req.Get(url)
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

			Convey("It will get the referenced payment values", func() {
				So(payment.ID, ShouldEqual, paymentID)
				So(payment.Version, ShouldEqual, version)
				So(payment.OrganisationID, ShouldEqual, organisationID)
				So(payment.Type, ShouldEqual, paymentType)

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

	})
}
