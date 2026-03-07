-- +goose up

CREATE TABLE posts(
    id              uuid NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    title           TEXT NOT NULL,
    url             TEXT NOT NULL,
    description     TEXT NOT NULL,
    published_at    TIMESTAMP NOT NULL,
    feed_id         uuid NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    PRIMARY KEY(id),
    CONSTRAINT posts__url UNIQUE (url)
);

-- +goose down

DROP TABLE posts;