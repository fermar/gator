-- name: GetFeeds :many
SELECT feeds.name , feeds.url , users.name 
FROM feeds 
inner join users on users.id = feeds.user_id;

-- INSERT INTO users (id, created_at, updated_at, name)
-- VALUES (
--     $1,
--     $2,
--     $3,
--     $4
-- )
-- RETURNING *;
