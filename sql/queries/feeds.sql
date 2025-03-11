-- name: CreateFeeds :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
    )
RETURNING *;


-- name: MarkFeedFetched :exec
UPDATE feeds set 
updated_at = $1, last_fetched_at= $2
where 
id = $3;

-- name: GetNextFeedToFetch :one

Select *
from feeds
order by last_fetched_at ASC
NULLS FIRST
limit 1;
