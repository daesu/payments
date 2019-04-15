package payment

import (
	"context"

	log "github.com/daesu/payments/logging"

	"github.com/daesu/payments/gen/models"
	"github.com/daesu/payments/gen/restapi/operations/payment"
)

type Service interface {
	ListPayments(ctx context.Context, params *payment.ListPaymentsParams) (*models.PaymentList, error)
	CreatePayment(ctx context.Context, params *payment.CreatePaymentParams) (*models.Payment, error)
	UpdatePayment(ctx context.Context, params *payment.UpdatePaymentParams) (*models.Payment, error)
	GetPayment(ctx context.Context, params *payment.GetPaymentParams) (*models.Payment, error)
	DeletePayment(ctx context.Context, params *payment.DeletePaymentParams) error
}

type service struct {
	repo Repository
}

// NewService create a service instance
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// ListPayments calls on repository to get a list of payments
func (s *service) ListPayments(ctx context.Context, params *payment.ListPaymentsParams) (*models.PaymentList, error) {
	log.Info("entered service ListPayments")

	payments, err := s.repo.ListPayments(ctx, params)
	if err != nil {
		log.Error(log.Trace(), err)
		return nil, err
	}

	meta := models.ListMetadata{}
	meta.TotalSize = int64(len(payments))
	meta.PageSize = int64(len(payments))

	ul := models.PaymentList{}
	ul.Data = payments
	ul.Metadata = &meta

	return &ul, nil
}

// CreatePayment calls on repository to create a new payment
// and then to return the payment response expected.
func (s *service) CreatePayment(ctx context.Context, params *payment.CreatePaymentParams) (*models.Payment, error) {
	log.Info("entered service CreatePayment")

	payment, err := s.repo.CreatePayment(ctx, params.Payment)
	if err != nil {
		log.Error(log.Trace(), err)
		return nil, err
	}

	paymentDetails, err := s.repo.GetPayment(ctx, payment.ID)
	if err != nil {
		log.Error(log.Trace(), err)
		return nil, err
	}

	return paymentDetails, err

}

// GetPayment returns the expected payment response.
func (s *service) GetPayment(ctx context.Context, params *payment.GetPaymentParams) (*models.Payment, error) {
	log.Info("entered service GetPayment")

	payment, err := s.repo.GetPayment(ctx, params.PaymentID)
	if err != nil {
		log.Error(log.Trace(), err)
		return nil, err
	}

	return payment, err
}

// DeletePayment deletes specified payment if it exists
func (s *service) DeletePayment(ctx context.Context, params *payment.DeletePaymentParams) error {
	log.Info("entered service DeletePayment")

	err := s.repo.DeletePayment(ctx, params.PaymentID)
	if err != nil {
		log.Error(log.Trace(), err)
		return err
	}

	return nil
}

// UpdatePayment replaces the specified payment with the details specified
// except for the EnDToEndReference field.
func (s *service) UpdatePayment(ctx context.Context, params *payment.UpdatePaymentParams) (*models.Payment, error) {
	log.Info("entered service UpdatePayment")

	payment, err := s.repo.UpdatePayment(ctx, params.UpdatePayment, params.PaymentID)
	if err != nil {
		log.Error(log.Trace(), err)
		return nil, err
	}

	paymentDetails, err := s.repo.GetPayment(ctx, payment.ID)
	if err != nil {
		log.Error(log.Trace(), err)
		return nil, err
	}

	return paymentDetails, err
}
