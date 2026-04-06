-- +migrate Up
CREATE TABLE IF NOT EXISTS page_pricing_plans (
    id             BIGSERIAL    PRIMARY KEY,
    name           VARCHAR(100) NOT NULL,
    price_monthly  INT          NOT NULL DEFAULT 0,
    price_yearly   INT          NOT NULL DEFAULT 0,
    features       JSONB        NOT NULL DEFAULT '[]',
    is_popular     BOOLEAN      NOT NULL DEFAULT FALSE,
    is_active      BOOLEAN      NOT NULL DEFAULT TRUE,
    sort_order     INT          NOT NULL DEFAULT 0,
    created_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

INSERT INTO page_pricing_plans (name, price_monthly, price_yearly, features, is_popular, sort_order) VALUES
('Basic',   199000, 159000, '["Reservation Management","Basic Waitlist Features","Automated Reminders","Queue Management"]', false, 1),
('Pro',     399000, 319000, '["Basic features plus:","Advanced Waitlist Management","Revenue Reports & Analytics"]', true, 2),
('Premium', 799000, 689000, '["All Pro features plus:","Custom Integration","Priority Support","Dedicated Account Manager"]', false, 3);

-- +migrate Down
DROP TABLE IF EXISTS page_pricing_plans;
