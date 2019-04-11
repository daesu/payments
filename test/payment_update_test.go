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

		var paymentPurpose = "new purpose"
		var reference = "new reference"

		var payment models.Payment
		attributes := models.PaymentAttribute{}
		beneficiary := models.CustomerAccount{}
		debtor := models.CustomerAccount{}

		attributes.BeneficiaryParty = &beneficiary
		attributes.BeneficiaryParty = &debtor
		payment.Attributes = &attributes

		Convey("When user requests payments/<payment-id> endpoint with valid POST data", func() {

			var paymentID = "61b0c143-f1f9-457d-a889-80570b820348"

			url := fmt.Sprintf("http://localhost:8080/v1/payments/%s", paymentID)

			json := fmt.Sprintf(`
			{
				"payment_purpose": "%s",
				"reference": "%s" 
			}`, paymentPurpose, reference)

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

			Convey("It will get the specified payment values", func() {
				So(payment.Attributes.PaymentPurpose, ShouldEqual, paymentPurpose)
				So(payment.Attributes.Reference, ShouldEqual, reference)
			})

		})

	})
}
