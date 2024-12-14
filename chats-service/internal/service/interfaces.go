package service

import (
	"context"
	"go13/chats-service/internal/models"
)

type ChatsRepo interface {
	CreateChat(ctx context.Context, chat models.Chat) (models.Chat, error)
	GetChatById(ctx context.Context, chatId int) (models.Chat, error)
	DeleteChat(ctx context.Context, chatId int) error
}

type RolesRepo interface {
	CreateRole(ctx context.Context, chatId int, role models.Role) (models.Role, error)
	ListRoles(ctx context.Context, chatId int) ([]models.Role, error)
	GetRoleById(ctx context.Context, chatId int, roleId int) (models.Role, error)
	UpdateRole(ctx context.Context, chatId int, roleId int, newRole models.Role) (models.Role, error)
	DeleteRole(ctx context.Context, chatId int, roleId int) error
	DeleteRolesForChat(ctx context.Context, chatId int) error
	GetRoleForMember(ctx context.Context, chatId int, userId string) (models.Role, error)
	GetMemberRoleId(ctx context.Context, chatId int) (int, error)
}

type MembersRepo interface {
	AddMember(ctx context.Context, chatId int, member models.Member) error
	DeleteMembersForChat(ctx context.Context, chatId int) error
	UnsetRole(ctx context.Context, chatId int, oldRoleId, newRoleId int) error
	CheckMemberIsBanned(ctx context.Context, chatid int, userId string) (bool, error)
}

type MessagesRepo interface {
	GetMessageById(ctx context.Context, chatId int, messageId int) (models.Message, error)
}
