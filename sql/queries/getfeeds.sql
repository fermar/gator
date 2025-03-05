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

-- name: GetFeedsFollowsForUser :many
SELECT feed_follows.*, users.name as user_name , feeds.name as feed_name
from feed_follows
inner join users on users.ID = feed_follows.user_id
inner join feeds on feeds.ID = feed_follows.feed_id
    where 
    users.id = $1;
