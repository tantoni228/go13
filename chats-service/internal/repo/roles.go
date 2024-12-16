package repo

import (
	"context"
	"go13/chats-service/internal/models"
)

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
