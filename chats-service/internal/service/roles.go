package service

import (
	"context"
	"errors"
	"fmt"
	"go13/chats-service/internal/models"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
)

type RolesService struct {
	rolesRepo   RolesRepo
	membersRepo MembersRepo
	trManager   trm.Manager
}

func NewRolesService(
	rolesRepo RolesRepo,
	membersRepo MembersRepo,
	trManager trm.Manager,
) *RolesService {
	return &RolesService{
		rolesRepo:   rolesRepo,
		membersRepo: membersRepo,
		trManager:   trManager,
	}
}

func (rs *RolesService) CreateRole(ctx context.Context, chatId int, role models.Role) (models.Role, error) {
	op := "RolesService.CreateRole"

	role.IsSystem = false
	created, err := rs.rolesRepo.CreateRole(ctx, chatId, role)
	if err != nil {
		return models.Role{}, fmt.Errorf("%s: %w", op, err)
	}

	return created, nil
}

func (rs *RolesService) ListRoles(ctx context.Context, chatId int) ([]models.Role, error) {
	op := "RolesService.ListRoles"

	roles, err := rs.rolesRepo.ListRoles(ctx, chatId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return roles, nil
}

func (rs *RolesService) GetRoleById(ctx context.Context, chatId int, roleId int) (models.Role, error) {
	op := "RolesService.GetRoleById"

	role, err := rs.rolesRepo.GetRoleById(ctx, chatId, roleId)
	if err != nil {
		return models.Role{}, fmt.Errorf("%s: %w", op, err)
	}

	return role, nil
}

func (rs *RolesService) DeleteRole(ctx context.Context, chatId int, roleId int) error {
	op := "RolesService.DeleteRole"

	err := rs.trManager.Do(ctx, func(ctx context.Context) error {
		memberRoleId, err := rs.rolesRepo.GetMemberRoleId(ctx, chatId)
		if err != nil {
			if errors.Is(err, models.ErrChatNotFound) {
				return fmt.Errorf("get member role id: %w", models.ErrChatOrRoleNotFound)
			}
			return fmt.Errorf("get member role id: %w", err)
		}

		if err := rs.membersRepo.UnsetRole(ctx, chatId, roleId, memberRoleId); err != nil {
			return fmt.Errorf("unset role: %w", err)
		}

		if err := rs.rolesRepo.DeleteRole(ctx, chatId, roleId); err != nil {
			return fmt.Errorf("delete role: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("%s: trManager.Do: %w", op, err)
	}

	return nil
}
