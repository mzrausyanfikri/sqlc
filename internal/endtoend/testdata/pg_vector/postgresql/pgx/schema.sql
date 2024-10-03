CREATE EXTENSION IF NOT EXISTS "vector";

CREATE TABLE items (
    id bigserial PRIMARY KEY,
    embedding vector(3),
    half_embedding halfvec(3),
    sparse_embedding sparsevec(5)
);
