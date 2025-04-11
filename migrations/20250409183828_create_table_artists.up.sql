CREATE TABLE novelku.artists (
    id    UUID PRIMARY KEY NOT NULL,
    slug       VARCHAR(50) NOT NULL UNIQUE,
    name  VARCHAR(50) NOT NULL UNIQUE
);

CREATE INDEX idx_artists_name_trgm ON novelku.artists USING GIN (name gin_trgm_ops);