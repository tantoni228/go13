package repo

import (
	"context"
	"go13/chats-service/internal/models"
)

type MessagesRepo interface {
	GetMessageById(ctx context.Context, chatId int, messageId int) (models.Message, error)
}
