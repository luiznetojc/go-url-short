CREATE KEYSPACE IF NOT EXISTS urlshortener
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE urlshortener;

CREATE TABLE IF NOT EXISTS urls (
    short_code TEXT PRIMARY KEY,
    original_url TEXT,
    user_id UUID,
    created_at TIMESTAMP,
    expires_at TIMESTAMP
);
