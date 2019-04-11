// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/daesu/payments/gen/restapi/operations"
	"github.com/daesu/payments/gen/restapi/operations/health"
	"github.com/daesu/payments/gen/restapi/operations/payment"
)

//go:generate swagger generate server --target ../../gen --name Payments --spec ../../swagger/payment.yaml

func configureFlags(api *operations.PaymentsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PaymentsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.PaymentCreatePaymentHandler = payment.CreatePaymentHandlerFunc(func(params payment.CreatePaymentParams) middleware.Responder {
		return middleware.NotImplemented("operation payment.CreatePayment has not yet been implemented")
	})
	api.PaymentDeletePaymentHandler = payment.DeletePaymentHandlerFunc(func(params payment.DeletePaymentParams) middleware.Responder {
		return middleware.NotImplemented("operation payment.DeletePayment has not yet been implemented")
	})
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(func(params health.GetHealthParams) middleware.Responder {
		return middleware.NotImplemented("operation health.GetHealth has not yet been implemented")
	})
	api.PaymentGetPaymentHandler = payment.GetPaymentHandlerFunc(func(params payment.GetPaymentParams) middleware.Responder {
		return middleware.NotImplemented("operation payment.GetPayment has not yet been implemented")
	})
	api.PaymentListPaymentsHandler = payment.ListPaymentsHandlerFunc(func(params payment.ListPaymentsParams) middleware.Responder {
		return middleware.NotImplemented("operation payment.ListPayments has not yet been implemented")
	})
	api.PaymentUpdatePaymentHandler = payment.UpdatePaymentHandlerFunc(func(params payment.UpdatePaymentParams) middleware.Responder {
		return middleware.NotImplemented("operation payment.UpdatePayment has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
