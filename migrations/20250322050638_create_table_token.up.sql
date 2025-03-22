CREATE TABLE tokens (
    id            CHAR(36) PRIMARY KEY NOT NULL,
    refresh_token VARCHAR(255) UNIQUE NOT NULL,
    expires_at    TIMESTAMP NOT NULL,
    user_id       CHAR(36) NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_tokens_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB;