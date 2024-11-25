-- name: CreateUser :one
-- Insert a new user into the users table
INSERT INTO users (
    username,
    email,
    password_hash,
    first_name,
    last_name,
    bio,
    date_of_birth,
    phone_number,
    profile_picture_url,
    address
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetUserByID :one
-- Fetch a single user by their ID
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
-- Fetch a single user by their username
SELECT * FROM users
WHERE username = $1;

-- name: UpdateUser :one
-- Update a user's information
UPDATE users
SET 
    first_name = COALESCE($2, first_name),
    last_name = COALESCE($3, last_name),
    bio = COALESCE($4, bio),
    date_of_birth = COALESCE($5, date_of_birth),
    phone_number = COALESCE($6, phone_number),
    profile_picture_url = COALESCE($7, profile_picture_url),
    address = COALESCE($8, address),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
-- Delete a user by their ID
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
-- Fetch a paginated list of users from the table
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateLastLogin :exec
-- Update the last login timestamp for a user
UPDATE users
SET last_login = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: SetUserActiveStatus :exec
-- Update the active status of a user
UPDATE users
SET is_active = $2
WHERE id = $1;

