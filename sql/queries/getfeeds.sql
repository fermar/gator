-- name: GetFeeds :many
SELECT feeds.name , feeds.url , users.name 
FROM feeds 
inner join users on users.id = feeds.user_id;


-- name: GetFeedsByName :one
SELECT feeds.*
FROM feeds 
-- inner join users on users.id = feeds.user_id
where feeds.name = $1
limit 1;


-- name: GetFeedsByUrl :one
SELECT feeds.*
FROM feeds 
where feeds.url = $1
limit 1;

