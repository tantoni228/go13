package handlers

import (
	"context"
	"errors"
	"go13/chats-service/internal/models"
	"go13/pkg/logger"
	api "go13/pkg/ogen/chats-service"

	"go.uber.org/zap"
)

type RolesService interface {
	CreateRole(ctx context.Context, chatId int, role models.Role) (models.Role, error)
	DeleteRole(ctx context.Context, chatId int, roleId int) error
}

type RolesHandler struct {
	rolesService RolesService
}

func NewRolesHandler(rolesService RolesService) *RolesHandler {
	return &RolesHandler{
		rolesService: rolesService,
	}
}

// CheckAccess implements CheckAccess operation.
//
// Check access to uri with method.
//
// GET /roles/check-access
func (rh *RolesHandler) CheckAccess(ctx context.Context, params api.CheckAccessParams) (api.CheckAccessRes, error) {
	return &api.CheckAccessNoContent{}, nil
}

// CreateRole implements createRole operation.
//
// Create role in Chat.
//
// POST /roles
func (rh *RolesHandler) CreateRole(ctx context.Context, req *api.RoleInput, params api.CreateRoleParams) (api.CreateRoleRes, error) {
	role := models.Role{
		Name:              req.Name,
		CanBanUsers:       req.CanBanUsers,
		CanEditRoles:      req.CanEditRoles,
		CanDeleteMessages: req.CanDeleteMessages,
		CanGetJoinCode:    req.CanGetJoinCode,
		CanEditChatInfo:   req.CanEditChatInfo,
		CanDeleteChat:     req.CanDeleteChat,
	}
	created, err := rh.rolesService.CreateRole(ctx, int(params.ChatId), role)
	if err != nil {
		if errors.Is(err, models.ErrChatNotFound) {
			return &api.ChatNotFoundResponse{}, nil
		}
		if errors.Is(err, models.ErrRoleAlreadyExists) {
			return &api.CreateRoleConflict{}, nil
		}
		logger.FromCtx(ctx).Error("create role", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.Role{
		ID:                api.RoleId(created.Id),
		Name:              created.Name,
		CanBanUsers:       created.CanBanUsers,
		CanEditRoles:      created.CanEditRoles,
		CanDeleteMessages: created.CanDeleteMessages,
		CanGetJoinCode:    created.CanGetJoinCode,
		CanEditChatInfo:   created.CanEditChatInfo,
		CanDeleteChat:     created.CanDeleteChat,
	}, nil
}

// DeleteRole implements deleteRole operation.
//
// Delete role in chat.
//
// DELETE /roles/{roleId}
func (rh *RolesHandler) DeleteRole(ctx context.Context, params api.DeleteRoleParams) (api.DeleteRoleRes, error) {
	err := rh.rolesService.DeleteRole(ctx, int(params.ChatId), int(params.RoleId))
	if err != nil {
		if errors.Is(err, models.ErrChatNotFound) || errors.Is(err, models.ErrRoleNotFound) {
			return &api.DeleteRoleNotFound{}, nil
		}
		logger.FromCtx(ctx).Error("delete role", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.DeleteRoleNoContent{}, nil
}

// GetRoleById implements getRoleById operation.
//
// Get role in Chat.
//
// GET /roles/{roleId}
func (rh *RolesHandler) GetRoleById(ctx context.Context, params api.GetRoleByIdParams) (api.GetRoleByIdRes, error) {
	return &api.Role{}, nil
}

// ListRoles implements listRoles operation.
//
// Get roles for chat.
//
// GET /roles
func (rh *RolesHandler) ListRoles(ctx context.Context, params api.ListRolesParams) (api.ListRolesRes, error) {
	return &api.ListRolesOKApplicationJSON{}, nil
}

// UpdateRole implements updateRole operation.
//
// Update role in chat.
//
// PUT /roles/{roleId}
func (rh *RolesHandler) UpdateRole(ctx context.Context, req *api.RoleInput, params api.UpdateRoleParams) (api.UpdateRoleRes, error) {
	return &api.Role{}, nil
}
