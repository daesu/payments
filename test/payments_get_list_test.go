package test

import (
	"testing"

	"github.com/daesu/payments/gen/models"
	"github.com/imroc/req"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {

}

func TestGetPaymentsEndpoint(t *testing.T) {
	Convey("Given API is running", t, func() {

		var payments models.PaymentList

		Convey("When user access list payments endpoint", func() {

			url := "http://localhost:8080/v1/payments?pageSize=2"

			resp, err := req.Get(url)
			if err != nil {
				t.Error("Response: ", resp.String())
				t.Fail()

			}

			Convey("It will get 200 status", func() {
				So(resp.Response().StatusCode, ShouldEqual, 200)
			})

			Convey("It will get array of payments with length equals 2", func() {
				err = resp.ToJSON(&payments)
				if err != nil {
					t.Error("Error: ", err.Error())
				}

				totalPayments := len(payments.Data)
				So(totalPayments, ShouldEqual, 2)

			})

		})

	})
}
