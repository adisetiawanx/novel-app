CREATE TABLE novelku.genres (
    id    UUID PRIMARY KEY NOT NULL,
    slug  VARCHAR(50) NOT NULL UNIQUE,
    name  VARCHAR(50) NOT NULL UNIQUE
);

CREATE INDEX idx_genres_name_trgm ON novelku.genres USING GIN (name gin_trgm_ops);