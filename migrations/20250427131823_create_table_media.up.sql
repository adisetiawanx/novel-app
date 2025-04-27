CREATE TABLE novelku.media (
    id         UUID PRIMARY KEY NOT NULL,
    url        VARCHAR(255) UNIQUE NOT NULL,
    name       VARCHAR(255) NOT NULL,
    size       INT,
    type       VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
