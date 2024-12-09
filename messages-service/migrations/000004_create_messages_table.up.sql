CREATE TABLE IF NOT EXISTS messages (
	id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
	chat_id INTEGER NOT NULL,
	message TEXT NOT NULL,
	edited BOOLEAN DEFAULT false,
    send_timestamp INTEGER NOT NULL,
	FOREIGN KEY (chat_id) REFERENCES chats (id)
);