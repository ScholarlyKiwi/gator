-- +goose up

CREATE TABLE feeds(
    id              uuid NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    name            text NOT NULL,
    url             text NOT NULL UNIQUE,
    user_id         uuid NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (id)
);

-- +goose down

DROP TABLE feeds;