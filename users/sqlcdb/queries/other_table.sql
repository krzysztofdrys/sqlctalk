-- name: GetBookForNewestUsers :many

WITH newest_users AS (
    SELECT * FROM users
    WHERE created_at >= $1
) SELECT
      newest_users.id as user_id,
      books.id as book_id
      FROM newest_users JOIN books ON books.owner_user_id = newest_users.id;
