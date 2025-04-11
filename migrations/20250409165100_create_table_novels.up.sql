CREATE TABLE novelku.novels (
    id                  UUID PRIMARY KEY NOT NULL,
    post_status         VARCHAR(20) NOT NULL DEFAULT 'publish' CHECK (post_status IN ('publish', 'draft')),
    title               VARCHAR(150) NOT NULL,
    slug                VARCHAR(100) NOT NULL UNIQUE,
    alternative_title   VARCHAR(150),
    synopsis            TEXT,
    status              VARCHAR(20) NOT NULL DEFAULT 'unknown' CHECK (status IN ('unknown', 'ongoing', 'completed', 'hiatus')),
    release_year        SMALLINT NOT NULL,
    country             VARCHAR(20) NOT NULL DEFAULT 'unknown' CHECK (country IN ('unknown', 'china', 'korea', 'japan', 'indonesia')),
    rating_total        NUMERIC(4,1) NOT NULL,
    chapter_total       INT NOT NULL DEFAULT 0,
    comment_total       INT NOT NULL DEFAULT 0,
    vote_total          INT NOT NULL DEFAULT 0,
    bookmark_total      INT NOT NULL DEFAULT 0,
    view_total          INT NOT NULL DEFAULT 0,
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT chk_rating_range CHECK (rating_total BETWEEN 0.0 AND 10.0)
);

CREATE INDEX idx_novels_post_status ON novelku.novels(post_status);
CREATE INDEX idx_novels_status ON novelku.novels(status);
CREATE INDEX idx_novels_slug ON novelku.novels(slug);
CREATE INDEX idx_novels_release_year ON novelku.novels(release_year);
CREATE INDEX idx_novels_country ON novelku.novels(country);
CREATE INDEX idx_novels_rating ON novelku.novels(rating_total);
CREATE INDEX idx_novels_created_at ON novelku.novels(created_at DESC);
CREATE INDEX idx_novels_chapter_total_asc ON novelku.novels(chapter_total ASC);
CREATE INDEX idx_novels_comment_total_asc ON novelku.novels(comment_total ASC);
CREATE INDEX idx_novels_vote_total_asc ON novelku.novels(vote_total ASC);
CREATE INDEX idx_novels_bookmark_total_asc ON novelku.novels(bookmark_total ASC);
CREATE INDEX idx_novels_view_total_asc ON novelku.novels(view_total ASC);
CREATE INDEX idx_novels_chapter_total_desc ON novelku.novels(chapter_total DESC);
CREATE INDEX idx_novels_comment_total_desc ON novelku.novels(comment_total DESC);
CREATE INDEX idx_novels_vote_total_desc ON novelku.novels(vote_total DESC);
CREATE INDEX idx_novels_bookmark_total_desc ON novelku.novels(bookmark_total DESC);
CREATE INDEX idx_novels_view_total_desc ON novelku.novels(view_total DESC);
CREATE INDEX idx_novels_title_trgm ON novelku.novels USING GIN (title gin_trgm_ops);