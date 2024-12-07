package service

import (
	"context"
	"fmt"

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

func (rs *RolesService) DeleteRole(ctx context.Context, chatId int, roleId int) error {
	op := "RolesService.DeleteRole"

	err := rs.trManager.Do(ctx, func(ctx context.Context) error {
		memberRoleId, err := rs.rolesRepo.GetMemberRoleId(ctx, chatId)
		if err != nil {
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
