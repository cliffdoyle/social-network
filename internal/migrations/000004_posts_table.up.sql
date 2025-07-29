CREATE TABLE IF NOT EXISTS posts (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    group_id TEXT REFERENCES groups(id) ON DELETE SET NULL,
    title TEXT,
    content TEXT,
    media_url TEXT,
    media_type TEXT CHECK(media_type IS NULL OR media_type IN ('image', 'gif')),
    privacy TEXT NOT NULL CHECK(privacy IN ('public', 'followers', 'private')),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);