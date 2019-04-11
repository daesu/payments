package test

import (
	"fmt"
	"testing"

	"github.com/imroc/req"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDeletePaymentEndpoint(t *testing.T) {
	Convey("Given API is running", t, func() {

		Convey("When user requests DELETE payments/<payment-id> endpoint", func() {

			var paymentID = "fd54832d-d872-428b-a10d-17ddf782b4df"

			url := fmt.Sprintf("http://localhost:8080/v1/payments/%s", paymentID)

			header := req.Header{
				"Content-Type": "application/json",
			}

			resp, err := req.Delete(url, header)

			if err != nil {
				t.Error("Response: ", resp.String())
				t.Fail()
			}

			Convey("It will get 200 status", func() {
				So(resp.Response().StatusCode, ShouldEqual, 200)
			})

		})

		Convey("When user requests DELETE payments/<payment-id> endpoint for an already deleted payment", func() {

			var paymentID = "fd54832d-d872-428b-a10d-17ddf782b4df"

			url := fmt.Sprintf("http://localhost:8080/v1/payments/%s", paymentID)

			header := req.Header{
				"Content-Type": "application/json",
			}

			resp, err := req.Delete(url, header)

			if err != nil {
				t.Error("Response: ", resp.String())
				t.Fail()
			}

			Convey("It will get 404 status", func() {
				So(resp.Response().StatusCode, ShouldEqual, 404)
			})

		})

	})
}
