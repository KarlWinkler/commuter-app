-- name: ListTrips :many
SELECT * FROM trip;

-- name: CreateTrip :one
INSERT INTO trip (name, distance, mode)
VALUES ($1, $2, $3)
RETURNING *;