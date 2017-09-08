
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ language 'plpgsql';
-- +goose StatementEnd

CREATE TYPE token_type AS ENUM (
  'RESET',
  'VALIDATION'
);

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username TEXT UNIQUE,
  password TEXT,
  email VARCHAR(170) UNIQUE,
  version INTEGER DEFAULT 10,
  is_deleted BOOLEAN DEFAULT false,
  is_validated BOOLEAN DEFAULT false,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NULL,
  deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TRIGGER updated_at_trigger BEFORE UPDATE
ON users FOR EACH ROW EXECUTE PROCEDURE
updated_at();

CREATE TABLE IF NOT EXISTS time_token (
  id UUID NOT NULL DEFAULT uuid_generate_v4(),
  user_id INTEGER,
  token_type token_type,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  used_at TIMESTAMPTZ DEFAULT NULL
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users CASCADE;
DROP TABLE time_token CASCADE;
DROP TYPE token_type CASCADE;
DROP EXTENSION "uuid-ossp";
DROP FUNCTION updated_at() CASCADE;
