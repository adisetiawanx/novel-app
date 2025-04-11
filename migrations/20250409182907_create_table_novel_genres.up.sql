CREATE TABLE novelku.novel_genres (
    novel_id    UUID NOT NULL,
    genre_id    UUID NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (novel_id, genre_id),
    FOREIGN KEY (novel_id) REFERENCES novelku.novels(id) ON DELETE CASCADE,
    FOREIGN KEY (genre_id) REFERENCES novelku.genres(id) ON DELETE CASCADE
);

CREATE INDEX idx_novel_genres_novel_id ON novelku.novel_genres(novel_id);
CREATE INDEX idx_novel_genres_genre_id ON novelku.novel_genres(genre_id);

