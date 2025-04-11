CREATE TABLE auth.tokens (
    id            UUID PRIMARY KEY NOT NULL,
    refresh_token VARCHAR(255) UNIQUE NOT NULL,
    expires_at    TIMESTAMP NOT NULL,
    user_id       UUID NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_tokens_users FOREIGN KEY (user_id) REFERENCES auth.users(id) ON DELETE CASCADE
);

CREATE INDEX idx_tokens_user_id ON auth.tokens(user_id);
