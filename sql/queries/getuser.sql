-- name: GetUser :one
SELECT * FROM users where name = $1 LIMIT 1 ;

-- INSERT INTO users (id, created_at, updated_at, name)
-- VALUES (
--     $1,
--     $2,
--     $3,
--     $4
-- )
-- RETURNING *;
