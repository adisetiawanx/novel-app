CREATE TABLE novelku.collection_items (
    collection_id   UUID NOT NULL,
    novel_id        UUID NOT NULL,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (collection_id, novel_id),
    FOREIGN KEY (collection_id) REFERENCES novelku.collections(id) ON DELETE CASCADE,
    FOREIGN KEY (novel_id) REFERENCES novelku.novels(id) ON DELETE CASCADE
);

CREATE INDEX idx_collection_items_novel_id ON novelku.collection_items(novel_id);
CREATE INDEX idx_collection_items_collection_id ON novelku.collection_items(collection_id);
