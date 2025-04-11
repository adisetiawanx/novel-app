CREATE TABLE novelku.novel_artists (
    novel_id    UUID NOT NULL,
    artist_id   UUID NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (novel_id, artist_id),
    FOREIGN KEY (novel_id) REFERENCES novelku.novels(id) ON DELETE CASCADE,
    FOREIGN KEY (artist_id) REFERENCES novelku.artists(id) ON DELETE CASCADE
);

CREATE INDEX idx_novel_artists_novel_id ON novelku.novel_artists(novel_id);
CREATE INDEX idx_novel_artists_artist_id ON novelku.novel_artists(artist_id);
