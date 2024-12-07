package service

import (
	"context"
	"go13/chats-service/internal/models"
)

type ChatsRepo interface {
	CreateChat(ctx context.Context, chat models.Chat) (models.Chat, error)
}

type RolesRepo interface {
	CreateRole(ctx context.Context, chatId int, role models.Role) (models.Role, error)
}

type MembersRepo interface {
	AddMember(ctx context.Context, chatId int, member models.Member) error
}
