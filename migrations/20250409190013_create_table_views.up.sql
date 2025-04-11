CREATE TABLE novelku.views (
    id           UUID PRIMARY KEY NOT NULL,
    novel_id     UUID NOT NULL,
    ip_address   INET NOT NULL,
    viewed_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (novel_id) REFERENCES novelku.novels(id) ON DELETE CASCADE
);

CREATE INDEX idx_novel_views_novel_id ON novelku.views(novel_id);
CREATE INDEX idx_novel_views_viewed_at ON novelku.views(viewed_at);
