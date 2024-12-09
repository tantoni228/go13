package mapper

import (
	"go13/chats-service/internal/models"
	api "go13/pkg/ogen/chats-service"
)

func ModelsRoleToApiRole(role models.Role) *api.Role {
	return &api.Role{
		ID:                api.RoleId(role.Id),
		Name:              role.Name,
		IsSystem:          role.IsSystem,
		CanBanUsers:       role.CanBanUsers,
		CanEditRoles:      role.CanEditRoles,
		CanDeleteMessages: role.CanDeleteMessages,
		CanGetJoinCode:    role.CanGetJoinCode,
		CanEditChatInfo:   role.CanEditChatInfo,
		CanDeleteChat:     role.CanDeleteChat,
	}
}

func ApiRoleInputToModelsRole(roleInput *api.RoleInput) models.Role {
	return models.Role{
		Name:              roleInput.GetName(),
		CanBanUsers:       roleInput.GetCanBanUsers(),
		CanEditRoles:      roleInput.GetCanEditRoles(),
		CanDeleteMessages: roleInput.GetCanDeleteMessages(),
		CanGetJoinCode:    roleInput.GetCanGetJoinCode(),
		CanEditChatInfo:   roleInput.GetCanEditChatInfo(),
		CanDeleteChat:     roleInput.GetCanDeleteChat(),
	}
}
