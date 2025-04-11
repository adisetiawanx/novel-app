CREATE TABLE novelku.ratings (
    user_id     UUID NOT NULL,
    novel_id    UUID NOT NULL,
    rating      NUMERIC(4,1) NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id, novel_id),
    FOREIGN KEY (novel_id) REFERENCES novelku.novels(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES auth.users(id) ON DELETE CASCADE,

    CONSTRAINT chk_rating_range CHECK (rating BETWEEN 0.0 AND 10.0)
);

CREATE INDEX idx_novel_ratings_novel_id ON novelku.ratings(novel_id);
CREATE INDEX idx_novel_ratings_user_id ON novelku.ratings(user_id);
