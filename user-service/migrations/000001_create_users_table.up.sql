CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    user_id UUID NOT NULL UNIQUE,
    user_name VARCHAR NOT NULL UNIQUE,
    user_email VARCHAR NOT NULL UNIQUE,
    user_password VARCHAR NOT NULL,
    user_bio TEXT,
);