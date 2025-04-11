CREATE TABLE novelku.authors (
    id    UUID PRIMARY KEY NOT NULL,
    slug       VARCHAR(50) NOT NULL UNIQUE,
    name  VARCHAR(50) NOT NULL UNIQUE
);

CREATE INDEX idx_authors_name_trgm ON novelku.authors USING GIN (name gin_trgm_ops);