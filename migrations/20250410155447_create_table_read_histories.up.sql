CREATE TABLE novelku.read_histories (
    user_id     UUID NOT NULL,
    chapter_id  UUID NOT NULL,
    novel_id    UUID NOT NULL,
    read_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id, chapter_id),

    FOREIGN KEY (user_id) REFERENCES auth.users(id) ON DELETE CASCADE,
    FOREIGN KEY (chapter_id) REFERENCES novelku.chapters(id) ON DELETE CASCADE,
    FOREIGN KEY (novel_id) REFERENCES novelku.novels(id) ON DELETE CASCADE
);

CREATE INDEX idx_read_histories_user_id ON novelku.read_histories(user_id);
CREATE INDEX idx_read_histories_novel_id ON novelku.read_histories(novel_id);
CREATE INDEX idx_read_histories_chapter_id ON novelku.read_histories(chapter_id);