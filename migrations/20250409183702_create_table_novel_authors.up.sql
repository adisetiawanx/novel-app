CREATE TABLE novelku.novel_authors (
    novel_id    UUID NOT NULL,
    author_id   UUID NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (novel_id, author_id),
    FOREIGN KEY (novel_id) REFERENCES novelku.novels(id) ON DELETE CASCADE,
    FOREIGN KEY (author_id) REFERENCES novelku.authors(id) ON DELETE CASCADE
);

CREATE INDEX idx_novel_authors_novel_id ON novelku.novel_authors(novel_id);
CREATE INDEX idx_novel_authors_author_id ON novelku.novel_authors(author_id);

