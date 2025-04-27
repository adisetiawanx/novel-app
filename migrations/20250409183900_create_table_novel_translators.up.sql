CREATE TABLE novelku.novel_translators (
    novel_id    UUID NOT NULL,
    translator_id   UUID NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (novel_id, translator_id),
    FOREIGN KEY (novel_id) REFERENCES novelku.novels(id) ON DELETE CASCADE,
    FOREIGN KEY (translator_id) REFERENCES novelku.translators(id) ON DELETE CASCADE
);

CREATE INDEX idx_novel_translators_novel_id ON novelku.novel_translators(novel_id);
CREATE INDEX idx_novel_translators_translators_id ON novelku.novel_translators(translator_id);
