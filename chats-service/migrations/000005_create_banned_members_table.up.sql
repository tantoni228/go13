ALTER TABLE IF EXISTS members DROP COLUMN IF EXISTS banned;
CREATE TABLE IF NOT EXISTS banned_members (
    user_id UUID NOT NULL,
    chat_id INTEGER NOT NULL,
    FOREIGN KEY (chat_id) REFERENCES chats (id)
)