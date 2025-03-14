-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
    )
RETURNING *;

-- name: GetPostsFromUser :many

Select posts.*
from posts
inner join feeds on feeds.id = posts.feed_id
where 
feeds.user_id = $1
order by posts.published_at desc
limit $2;
