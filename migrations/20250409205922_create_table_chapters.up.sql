CREATE TABLE novelku.chapters (
    id          UUID PRIMARY KEY NOT NULL,
    novel_id    UUID NOT NULL,
    slug        VARCHAR(100) NOT NULL UNIQUE,
    title        VARCHAR(150) NOT NULL,
    number      SERIAL NOT NULL,
    content     TEXT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (novel_id) REFERENCES novelku.novels(id)
);

CREATE INDEX idx_novel_chapters_slug ON novelku.chapters(slug);
CREATE INDEX idx_chapters_title_trgm ON novelku.chapters USING GIN (title gin_trgm_ops);