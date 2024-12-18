CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT (gen_random_uuid()) PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    hashed_password VARCHAR(255) NOT NULL,
    bio VARCHAR(255) DEFAULT '' NOT NULL
);
