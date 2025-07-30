CREATE TABLE IF NOT EXISTS following (
    followee_id TEXT NOT NULL REFERENCES users(id),
    follower_id TEXT NOT NULL REFERENCES users(id),
    PRIMARY KEY (folowee_id, follower_id)
);