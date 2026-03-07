-- +goose up

CREATE TABLE feed_follows(
    id   UUID NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    user_id         uuid NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    feed_id         uuid NOT NULL REFERENCES feeds (id) ON DELETE CASCADE,
    PRIMARY KEY(id),
    CONSTRAINT feed_follows__user_id__feed_id UNIQUE (user_id, feed_id)
);

-- +goose down
DROP TABLE feed_follows;