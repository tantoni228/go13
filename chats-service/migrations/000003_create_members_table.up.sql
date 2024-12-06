CREATE TABLE IF NOT EXISTS members (
    id SERIAL PRIMARY KEY,
    chat_id INTEGER NOT NULL,
    user_id UUID NOT NULL,
    role_id INTEGER NOT NULL,
    banned BOOLEAN NOT NULL DEFAULT false,
    FOREIGN KEY (chat_id, role_id) REFERENCES roles (chat_id, id),
    UNIQUE (chat_id, user_id)
);