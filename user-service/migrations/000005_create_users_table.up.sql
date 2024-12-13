CREATE TABLE IF NOT EXISTS users (
    id VARCHAR PRIMARY KEY,
    user_id UUID NOT NULL
    user_name VARCHAR NOT NULL,
    user_email VARCHAR NOT NULL,
    user_password VARCHAR NOT NULL,
    user_bio TEXT
    UNIQUE(user_id, user_name, user_email)
);