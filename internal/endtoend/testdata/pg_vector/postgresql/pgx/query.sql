-- name: InsertVector :exec
INSERT INTO items (embedding, half_embedding, sparse_embedding)
VALUES ($1, $2, $3);

-- name: NearestNeighbor :many
SELECT *
FROM items
ORDER BY embedding <-> $1
LIMIT 5;

-- name: NearestNeighborHalfvec :many
SELECT *
FROM items
ORDER BY half_embedding <-> $1
LIMIT 5;

-- name: NearestNeighborSparsevec :many
SELECT *
FROM items
ORDER BY sparse_embedding <-> $1
LIMIT 5;
