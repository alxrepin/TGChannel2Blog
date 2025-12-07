CREATE TABLE posts (
    id BIGINT NOT NULL PRIMARY KEY,
    group_id BIGINT NOT NULL,
    url TEXT,
    text TEXT,
    seo_title TEXT,
    seo_description TEXT,
    seo_keywords TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);
