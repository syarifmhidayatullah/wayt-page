-- +migrate Up
CREATE TABLE IF NOT EXISTS page_leads (
    id              BIGSERIAL    PRIMARY KEY,
    full_name       VARCHAR(150) NOT NULL,
    restaurant_name VARCHAR(150) NOT NULL,
    email           VARCHAR(200) NOT NULL,
    phone           VARCHAR(50)  NOT NULL DEFAULT '',
    created_at      TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_page_leads_created ON page_leads (created_at DESC);

-- +migrate Down
DROP TABLE IF EXISTS page_leads;
