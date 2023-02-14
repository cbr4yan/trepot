CREATE TABLE IF NOT EXISTS company(
    id text PRIMARY KEY,
    name text NOT NULL,
    email text NOT NULL,
    currency text NOT NULL,
    country text NOT NULL,
    active boolean NOT NULL DEFAULT true,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
