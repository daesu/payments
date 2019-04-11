package health

import (
	"context"
	"fmt"
	"time"

	log "github.com/daesu/payments/logging"

	"github.com/daesu/payments/gen/models"
)

// Service handles async log of audit event
type Service interface {
	GetHealth(ctx context.Context) (*models.Health, error)
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) GetHealth(ctx context.Context) (*models.Health, error) {
	log.Info("entered service GetHealth")

	t := time.Now()
	health := models.Health{
		DateTime: t.String(),
	}

	log.Debug(fmt.Sprintf("%#v", health))

	return &health, nil
}
