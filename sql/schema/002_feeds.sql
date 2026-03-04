-- +goose up

CREATE TABLE feeds(
    name text NOT NULL,
    url text NOT NULL UNIQUE,
    user_id uuid REFERENCES users (id) ON DELETE CASCADE
);

-- +goose down

DROP TABLE feeds;