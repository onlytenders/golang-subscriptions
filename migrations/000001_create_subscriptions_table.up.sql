-- +migrate Up
CREATE TABLE IF NOT EXISTS subscriptions
(
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id       UUID         NOT NULL,
    service_name  VARCHAR(255) NOT NULL,
    price         INTEGER      NOT NULL,
    start_date    DATE         NOT NULL,
    end_date      DATE
);
