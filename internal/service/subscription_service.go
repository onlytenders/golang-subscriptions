package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/onlytenders/golang-subscriptions/internal/models"
	"github.com/onlytenders/golang-subscriptions/internal/repository"
	"go.uber.org/zap"
)

type subscriptionService struct {
	repo   repository.SubscriptionRepository
	logger *zap.Logger
}

func NewSubscriptionService(repo repository.SubscriptionRepository, logger *zap.Logger) SubscriptionService {
	return &subscriptionService{repo: repo, logger: logger}
}

func (s *subscriptionService) Create(ctx context.Context, sub *models.Subscription) (uuid.UUID, error) {
	s.logger.Info("Creating subscription", zap.Any("subscription", sub))
	return s.repo.Create(ctx, sub)
}

func (s *subscriptionService) GetByID(ctx context.Context, id uuid.UUID) (*models.Subscription, error) {
	s.logger.Info("Getting subscription by ID", zap.String("id", id.String()))
	return s.repo.GetByID(ctx, id)
}

func (s *subscriptionService) Update(ctx context.Context, sub *models.Subscription) error {
	s.logger.Info("Updating subscription", zap.Any("subscription", sub))
	return s.repo.Update(ctx, sub)
}

func (s *subscriptionService) Delete(ctx context.Context, id uuid.UUID) error {
	s.logger.Info("Deleting subscription", zap.String("id", id.String()))
	return s.repo.Delete(ctx, id)
}

func (s *subscriptionService) List(ctx context.Context) ([]*models.Subscription, error) {
	s.logger.Info("Listing subscriptions")
	return s.repo.List(ctx)
}

func (s *subscriptionService) TotalCost(ctx context.Context, userID uuid.UUID, serviceName string, year int, month int) (int, error) {
	s.logger.Info("Calculating total cost",
		zap.String("user_id", userID.String()),
		zap.String("service_name", serviceName),
		zap.Int("year", year),
		zap.Int("month", month))

	subs, err := s.repo.List(ctx)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, sub := range subs {
		if sub.UserID == userID &&
			(serviceName == "" || sub.ServiceName == serviceName) &&
			sub.StartDate.Year() <= year && (sub.StartDate.Month() <= time.Month(month)) &&
			(sub.EndDate == nil || (sub.EndDate.Year() >= year && sub.EndDate.Month() >= time.Month(month))) {
			total += sub.Price
		}
	}
	return total, nil
}
