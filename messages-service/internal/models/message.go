package models

type Message struct {
	Id             int64      `json:"id" db:"id"`
	Sender_id      int64      `json:"sender_id" db:"user_id"`
	Chat_id        string     `json: "chat_id" db:"chat_id"`
	Message        string     `json:"message" db:"message"`
	Edited         bool       `json:"edited" db:"edited"`
	Send_timestamp int64      `json:"send_timestamp" db:"send_timestamp"`
}