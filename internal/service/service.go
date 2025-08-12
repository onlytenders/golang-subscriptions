package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/onlytenders/golang-subscriptions/internal/models"
)

type SubscriptionService interface {
	Create(ctx context.Context, subscription *models.Subscription) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Subscription, error)
	Update(ctx context.Context, subscription *models.Subscription) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]*models.Subscription, error)
	TotalCost(ctx context.Context, userID uuid.UUID, serviceName string, year int, month int) (int, error)
}
