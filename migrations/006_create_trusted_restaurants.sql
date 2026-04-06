-- +migrate Up
CREATE TABLE IF NOT EXISTS page_trusted_restaurants (
    id         BIGSERIAL    PRIMARY KEY,
    name       VARCHAR(150) NOT NULL,
    emoji      VARCHAR(20)  NOT NULL DEFAULT '🍽️',
    rating     VARCHAR(10)  NOT NULL DEFAULT '5.0',
    is_active  BOOLEAN      NOT NULL DEFAULT TRUE,
    sort_order INT          NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

INSERT INTO page_trusted_restaurants (name, emoji, rating, sort_order) VALUES
('Sushi House',    '🍣', '4.9', 1),
('Grill & Chill',  '🔥', '4.8', 2),
('Ocean View Cafe','🌊', '4.9', 3),
('Urban Dine',     '🏙️', '4.7', 4);

-- +migrate Down
DROP TABLE IF EXISTS page_trusted_restaurants;
