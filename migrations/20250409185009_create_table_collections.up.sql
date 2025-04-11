CREATE TABLE novelku.collections (
    id          UUID PRIMARY KEY,
    user_id     UUID NOT NULL,
    name        VARCHAR(100) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES auth.users(id) ON DELETE CASCADE,
    UNIQUE (user_id, name)
);

CREATE INDEX idx_collections_user_id ON novelku.collections(user_id);
CREATE INDEX idx_collections_name_trgm ON novelku.collections USING GIN (name gin_trgm_ops);
