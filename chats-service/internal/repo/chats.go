package repo

import (
	"context"
	"go13/chats-service/internal/models"
)

type ChatsRepo interface {
	CreateChat(ctx context.Context, chat models.Chat) (models.Chat, error)
	ListChatsForUser(ctx context.Context, userId string) ([]models.Chat, error)
	GetChatById(ctx context.Context, chatId int) (models.Chat, error)
	UpdateChat(ctx context.Context, chatId int, newChat models.Chat) (models.Chat, error)
	DeleteChat(ctx context.Context, chatId int) error
}
