CREATE TABLE users IF NOT EXIST (
    uuid TEXT PRIMARY KEY NOT NULL,
    firstName TEXT,
    lastName TEXT,
    nickName TEXT,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    profileStatus BOOLEAN
);