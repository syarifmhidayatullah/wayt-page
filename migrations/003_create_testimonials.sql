-- +migrate Up
CREATE TABLE IF NOT EXISTS page_testimonials (
    id         BIGSERIAL    PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    restaurant VARCHAR(150) NOT NULL,
    quote      TEXT         NOT NULL,
    rating     INT          NOT NULL DEFAULT 5,
    phone      VARCHAR(30)  NOT NULL DEFAULT '',
    is_active  BOOLEAN      NOT NULL DEFAULT TRUE,
    sort_order INT          NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

INSERT INTO page_testimonials (name, restaurant, quote, rating, phone, sort_order) VALUES
('Andi Susanto', 'Warung Sate Shinta',
 'Wayt has transformed our reservation management. More table turnovers and happier customers!',
 5, '+62 811-1234-5678', 1);

-- +migrate Down
DROP TABLE IF EXISTS page_testimonials;
