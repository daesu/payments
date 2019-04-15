package payment

import (
	"fmt"

	"github.com/daesu/payments/gen/restapi/operations"
	"github.com/daesu/payments/gen/restapi/operations/payment"
	log "github.com/daesu/payments/logging"
	"github.com/daesu/payments/swagger"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// Configure setups handlers on api with Service
func Configure(api *operations.PaymentsAPI, service Service) {

	api.PaymentListPaymentsHandler = payment.ListPaymentsHandlerFunc(func(params payment.ListPaymentsParams) middleware.Responder {
		log.Info("entering ListPaymentsHandler")

		log.WithFields(logrus.Fields{
			"Offset":   *params.Offset,
			"PageSize": *params.PageSize,
		}).Info("ListPaymentsHandler")

		log.Info(fmt.Sprintf("{URL: %#v}",
			*params.HTTPRequest.URL))

		result, err := service.ListPayments(params.HTTPRequest.Context(), &params)
		if err != nil {
			return swagger.PaymentErrorHandler("ListPayments", err)
		}
		return payment.NewListPaymentsOK().WithPayload(result)
	})

	api.PaymentCreatePaymentHandler = payment.CreatePaymentHandlerFunc(func(params payment.CreatePaymentParams) middleware.Responder {
		log.Info("entering PaymentCreatePaymentHandler")

		log.WithFields(logrus.Fields{
			"HttpRequest":   fmt.Sprintf("%#v", *params.HTTPRequest),
			"createPayment": fmt.Sprintf("%#v", *params.Payment),
		}).Info("CreatePaymentHandlerFunc")

		result, err := service.CreatePayment(params.HTTPRequest.Context(), &params)
		if err != nil {
			return swagger.PaymentErrorHandler("CreatePayment", err)
		}
		return payment.NewCreatePaymentCreated().WithPayload(result)
	})

	api.PaymentGetPaymentHandler = payment.GetPaymentHandlerFunc(func(params payment.GetPaymentParams) middleware.Responder {
		log.Info("entering PaymentGetPaymentHandler")

		log.WithFields(logrus.Fields{
			"HttpRequest": fmt.Sprintf("%#v", *params.HTTPRequest),
			"PaymentID":   params.PaymentID,
		}).Info("GetPaymentHandlerFunc")

		result, err := service.GetPayment(params.HTTPRequest.Context(), &params)
		if err != nil {
			return swagger.PaymentErrorHandler("GetPayment", err)
		}
		return payment.NewGetPaymentOK().WithPayload(result)
	})

	api.PaymentUpdatePaymentHandler = payment.UpdatePaymentHandlerFunc(func(params payment.UpdatePaymentParams) middleware.Responder {
		log.Info("entering PaymentUpdatePaymentHandler")

		log.WithFields(logrus.Fields{
			"HttpRequest":     fmt.Sprintf("%#v", *params.HTTPRequest),
			"UpdatePayment":   fmt.Sprintf("%#v", *params.UpdatePayment),
			"UpdatePaymentID": fmt.Sprintf("%#v", params.PaymentID),
		}).Info("UpdatePaymentHandlerFunc")

		result, err := service.UpdatePayment(params.HTTPRequest.Context(), &params)
		if err != nil {
			return swagger.PaymentErrorHandler("UpdatePayment", err)
		}
		return payment.NewUpdatePaymentOK().WithPayload(result)
	})

	api.PaymentDeletePaymentHandler = payment.DeletePaymentHandlerFunc(func(params payment.DeletePaymentParams) middleware.Responder {
		log.Info("entering PaymentDeletePaymentHandler")

		log.WithFields(logrus.Fields{
			"HttpRequest": fmt.Sprintf("%#v", *params.HTTPRequest),
			"PaymentID":   params.PaymentID,
		}).Info("DeletePaymentHandlerFunc")

		err := service.DeletePayment(params.HTTPRequest.Context(), &params)
		if err != nil {
			return swagger.PaymentErrorHandler("DeletePayment", err)
		}
		return payment.NewDeletePaymentOK()
	})

}
