-- +migrate Up
CREATE TABLE IF NOT EXISTS wayt_settings (
    id         BIGSERIAL    PRIMARY KEY,
    key        VARCHAR(100) NOT NULL UNIQUE,
    value      TEXT         NOT NULL DEFAULT '',
    label      VARCHAR(200) NOT NULL DEFAULT '',
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

INSERT INTO wayt_settings (key, value, label) VALUES
('hero_title',       'Fill Your Tables Faster, Increase Your Revenue',            'Hero: Judul Utama'),
('hero_subtitle',    'Smart reservation & queue management system to help you maximize table utilization and boost your business efficiency.', 'Hero: Subtitle'),
('hero_cta',         'Get Started',                                                'Hero: Teks Tombol CTA'),
('hero_phone',       '+62 811-1234-5678',                                          'Kontak: Nomor Telepon'),
('cta_title',        'Ready to Boost Your Restaurant''s Revenue?',                 'CTA Banner: Judul'),
('cta_subtitle',     'Try Wayt free for 14 days and see the difference!',          'CTA Banner: Subtitle'),
('cta_button',       'Get Started Now',                                            'CTA Banner: Teks Tombol'),
('footer_secure',    '100% Secure & Reliable',                                     'Footer: Label Secure'),
('footer_support',   'Friendly Customer Support',                                  'Footer: Label Support'),
('footer_whatsapp',  'Easy WhatsApp Integration',                                  'Footer: Label WhatsApp');

-- +migrate Down
DROP TABLE IF EXISTS wayt_settings;
