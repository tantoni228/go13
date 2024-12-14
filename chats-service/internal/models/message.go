package models

type Message struct {
	Id            int
	SenderId      string
	Message       string
	Edited        bool
	SendTimestamp int
}
