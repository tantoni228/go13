CREATE TABLE IF NOT EXISTS users (
    "id" UUID NOT NULL UNIQUE,
    "user_name" VARCHAR NOT NULL UNIQUE,
    "user_email" VARCHAR NOT NULL UNIQUE,
    "user_password" VARCHAR NOT NULL,
    "user_bio" TEXT,
    PRIMARY KEY("id")
);
