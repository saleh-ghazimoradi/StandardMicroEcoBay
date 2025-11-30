CREATE TABLE IF NOT EXISTS addresses (
    id bigserial PRIMARY KEY,
    address_line1 TEXT NOT NULL,
    address_line2 TEXT NOT NULL,
    city VARCHAR(25) NOT NULL,
    post_code VARCHAR(25) NOT NULL,
    country VARCHAR(25) NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);