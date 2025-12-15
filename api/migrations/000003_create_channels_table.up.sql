CREATE TABLE channels (
    id BIGINT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    avatar TEXT,
    subscriptions BIGINT NOT NULL DEFAULT 0
);