CREATE TABLE novelku.translators (
    id    UUID PRIMARY KEY NOT NULL,
    slug       VARCHAR(50) NOT NULL UNIQUE,
    name  VARCHAR(50) NOT NULL UNIQUE
);

CREATE INDEX idx_translator_name_trgm ON novelku.translators USING GIN (name gin_trgm_ops);