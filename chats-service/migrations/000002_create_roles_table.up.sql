CREATE TABLE IF NOT EXISTS roles (
	id SERIAL PRIMARY KEY,
	chat_id INTEGER NOT NULL REFERENCES chats (id),
	name VARCHAR(255) NOT NULL,
	can_ban_users BOOLEAN NOT NULL,
	can_edit_roles BOOLEAN NOT NULL,
	can_delete_messages BOOLEAN NOT NULL,
	can_get_join_code BOOLEAN NOT NULL,
	can_edit_chat_info BOOLEAN NOT NULL,
	can_delete_chat BOOLEAN NOT NULL,
	UNIQUE (id, chat_id),
	UNIQUE (name, chat_id)
);