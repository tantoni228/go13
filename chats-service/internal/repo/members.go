package repo

import (
	"context"
	"go13/chats-service/internal/models"
)

type MembersRepo interface {
	AddMember(ctx context.Context, chatId int, member models.Member) error
	DeleteMember(ctx context.Context, chatId int, userId string) error
	SetRoleForMember(ctx context.Context, chatId int, userId string, roleId int) error
	DeleteMembersForChat(ctx context.Context, chatId int) error
	UnsetRole(ctx context.Context, chatId int, oldRoleId, newRoleId int) error
	CheckMemberIsBanned(ctx context.Context, chatid int, userId string) (bool, error)
	AddMemberToBanned(ctx context.Context, chatId int, userId string) error
	DeleteMemberFromBanned(ctx context.Context, chatId int, userId string) error
	ListBannedMembers(ctx context.Context, chatId int) ([]string, error)
}
