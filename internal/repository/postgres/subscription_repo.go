package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/onlytenders/golang-subscriptions/internal/models"
)

type SubscriptionRepo struct {
	db *pgxpool.Pool
}

func NewSubscriptionRepo(db *pgxpool.Pool) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

func (r *SubscriptionRepo) Create(ctx context.Context, sub *models.Subscription) (uuid.UUID, error) {
	var id uuid.UUID
	err := r.db.QueryRow(ctx,
		`INSERT INTO subscriptions (user_id, service_name, price, start_date, end_date)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		sub.UserID, sub.ServiceName, sub.Price, sub.StartDate, sub.EndDate).Scan(&id)
	return id, err
}

func (r *SubscriptionRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Subscription, error) {
	var sub models.Subscription
	err := r.db.QueryRow(ctx,
		`SELECT id, user_id, service_name, price, start_date, end_date
		 FROM subscriptions WHERE id = $1`,
		id).Scan(&sub.ID, &sub.UserID, &sub.ServiceName, &sub.Price, &sub.StartDate, &sub.EndDate)
	return &sub, err
}

func (r *SubscriptionRepo) Update(ctx context.Context, sub *models.Subscription) error {
	_, err := r.db.Exec(ctx,
		`UPDATE subscriptions
		 SET user_id = $1, service_name = $2, price = $3, start_date = $4, end_date = $5
		 WHERE id = $6`,
		sub.UserID, sub.ServiceName, sub.Price, sub.StartDate, sub.EndDate, sub.ID)
	return err
}

func (r *SubscriptionRepo) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, `DELETE FROM subscriptions WHERE id = $1`, id)
	return err
}

func (r *SubscriptionRepo) List(ctx context.Context) ([]*models.Subscription, error) {
	rows, err := r.db.Query(ctx, `SELECT id, user_id, service_name, price, start_date, end_date FROM subscriptions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []*models.Subscription
	for rows.Next() {
		var sub models.Subscription
		if err := rows.Scan(&sub.ID, &sub.UserID, &sub.ServiceName, &sub.Price, &sub.StartDate, &sub.EndDate); err != nil {
			return nil, err
		}
		subs = append(subs, &sub)
	}
	return subs, nil
}
