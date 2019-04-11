package swagger

import (
	"github.com/daesu/payments/gen/restapi/operations/health"
	"github.com/daesu/payments/gen/restapi/operations/payment"
	"github.com/daesu/payments/utils"

	"github.com/sirupsen/logrus"

	"github.com/daesu/payments/gen/models"
	"github.com/go-openapi/runtime/middleware"
)

type codedResponse interface {
	Code() string
}

// ErrorResponse wraps the error in the api standard models.ErrorResponse object
func ErrorResponse(err error) *models.ErrorResponse {
	cd := ""
	if e, ok := err.(codedResponse); ok {
		cd = e.Code()
	}

	e := models.ErrorResponse{
		Code:    cd,
		Message: err.Error(),
	}
	return &e
}

func HealthErrorHandler(label string, err error) middleware.Responder {
	logrus.WithError(err).Error(label)

	return health.NewGetHealthBadRequest().WithPayload(ErrorResponse(err))

}

func PaymentErrorHandler(label string, err error) middleware.Responder {
	// logrus.WithError(err).Error(label)

	switch err.Error() {
	case utils.ErrNotFound.Error(), utils.ErrNotFoundOrganisation.Error():
		return payment.NewGetPaymentNotFound().WithPayload(ErrorResponse(err))
	case utils.ErrDuplicate.Error():
		return payment.NewCreatePaymentConflict().WithPayload(ErrorResponse(err))
	default:
		return payment.NewListPaymentsBadRequest().WithPayload(ErrorResponse(err))
	}
}
