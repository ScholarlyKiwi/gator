-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (
        id,
        created_at,
        updated_at,
        user_id,
        feed_id)
    SELECT $1,
           $2,
           $3,
           $4,
           $5
    RETURNING *
)
SELECT inserted_feed_follow.*,
       feeds.name AS feedname,
       users.name AS username
  FROM inserted_feed_follow
 INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
 INNER JOIN users ON users.id = inserted_feed_follow.user_id;

-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows
 WHERE user_id = $1
   AND feed_id = $2;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*,
       feeds.name AS feedname,
       users.name AS username
  FROM feed_follows
 INNER JOIN feeds ON feeds.id = feed_follows.feed_id
 INNER JOIN users ON users.id = feed_follows.user_id
 WHERE users.name = $1;

-- name: GetFeedFollowsForUserFeed :one
SELECT feed_follows.*,
       feeds.name AS feedname,
       users.name AS username
  FROM feed_follows
 INNER JOIN feeds ON feeds.id = feed_follows.feed_id
 INNER JOIN users ON users.id = feed_follows.user_id
 WHERE users.name = $1
   and feeds.url = $2;