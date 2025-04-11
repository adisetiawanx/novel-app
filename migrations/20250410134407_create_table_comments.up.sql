CREATE TABLE novelku.comments (
    id              UUID PRIMARY KEY NOT NULL,
    user_id         UUID NOT NULL,
    novel_id        UUID NOT NULL,
    chapter_id      UUID,
    parent_id       UUID,
    content         TEXT NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_novel_id FOREIGN KEY (novel_id) REFERENCES novelku.novels(id) ON DELETE CASCADE,
    CONSTRAINT fk_chapter_id FOREIGN KEY (chapter_id) REFERENCES novelku.chapters(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES auth.users(id) ON DELETE CASCADE,
    CONSTRAINT fk_parent_comment FOREIGN KEY (parent_id) REFERENCES novelku.comments(id) ON DELETE CASCADE
);

CREATE INDEX idx_comments_novel_id ON novelku.comments(novel_id);
CREATE INDEX idx_comments_chapter_id ON novelku.comments(chapter_id);
CREATE INDEX idx_comments_user_id ON novelku.comments(user_id);
CREATE INDEX idx_comments_parent_id ON novelku.comments(parent_id);
CREATE INDEX idx_comments_created_at ON novelku.comments(created_at DESC);
