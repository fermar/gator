-- +goose Up

CREATE TABLE feed_follows (
  id uuid primary key,
  created_at timestamp not null,
  updated_at timestamp not null,
  user_id  uuid not null,
  feed_id  uuid not null,
  UNIQUE (user_id,feed_id),
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feed_follows;
