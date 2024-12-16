package handlers

import (
	"context"
	"errors"
	"go13/chats-service/internal/models"
	"go13/chats-service/internal/transport/rest/auth"
	"go13/chats-service/internal/transport/rest/mapper"
	"go13/pkg/logger"
	api "go13/pkg/ogen/chats-service"
	"net/url"

	"go.uber.org/zap"
)

type RolesService interface {
	CreateRole(ctx context.Context, chatId int, role models.Role) (models.Role, error)
	ListRoles(ctx context.Context, chatId int) ([]models.Role, error)
	GetRoleForUser(ctx context.Context, chatId int, userId string) (models.Role, error)
	GetRoleById(ctx context.Context, chatId int, roleId int) (models.Role, error)
	UpdateRole(ctx context.Context, chatId int, roleId int, newRole models.Role) (models.Role, error)
	DeleteRole(ctx context.Context, chatId int, roleId int) error
}

type AccessService interface {
	CheckAccess(ctx context.Context, userId string, method string, u *url.URL) error
}

type RolesHandler struct {
	rolesService  RolesService
	accessService AccessService
}

func NewRolesHandler(rolesService RolesService, accessService AccessService) *RolesHandler {
	return &RolesHandler{
		rolesService:  rolesService,
		accessService: accessService,
	}
}

// CheckAccess implements CheckAccess operation.
//
// Check access to uri with method.
//
// GET /roles/check-access
func (rh *RolesHandler) CheckAccess(ctx context.Context, params api.CheckAccessParams) (api.CheckAccessRes, error) {
	userId := auth.UserIdFromCtx(ctx)
	logger.FromCtx(ctx).Debug(
		"check access",
		zap.String("uri", params.XTargetURI),
		zap.String("method", string(params.XTargetMethod)),
		zap.String("user_id", userId),
	)
	u, err := url.ParseRequestURI(params.XTargetURI)
	if err != nil {
		return &api.InvalidInputResponse{
			Message: "invalid X-Target-Uri parameter",
		}, nil
	}
	err = rh.accessService.CheckAccess(ctx, userId, string(params.XTargetMethod), u)
	if err != nil {
		notFoundErrors := []error{
			models.ErrEndpointNotFound,
			models.ErrRoleNotFound,
			models.ErrChatNotFound,
			models.ErrMemberNotFound,
			models.ErrMessageNotFound,
		}
		for _, targerErr := range notFoundErrors {
			if errors.Is(err, targerErr) {
				return &api.CheckAccessNotFound{}, nil
			}
		}
		if errors.Is(err, models.ErrAccessForbidden) {
			return &api.UnauthorizedResponse{}, nil
		}
		if errors.Is(err, models.ErrInvalidRouteParams) {
			return &api.InvalidInputResponse{
				Message: "invalid route parameters",
			}, nil
		}
		logger.FromCtx(ctx).Error("check access", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.CheckAccessNoContent{}, nil
}

// CreateRole implements createRole operation.
//
// Create role in Chat.
//
// POST /roles
func (rh *RolesHandler) CreateRole(ctx context.Context, req *api.RoleInput, params api.CreateRoleParams) (api.CreateRoleRes, error) {
	role := mapper.ApiRoleInputToModelsRole(req)
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

	return mapper.ModelsRoleToApiRole(created), nil
}

// DeleteRole implements deleteRole operation.
//
// Delete role in chat.
//
// DELETE /roles/{roleId}
func (rh *RolesHandler) DeleteRole(ctx context.Context, params api.DeleteRoleParams) (api.DeleteRoleRes, error) {
	err := rh.rolesService.DeleteRole(ctx, int(params.ChatId), int(params.RoleId))
	if err != nil {
		if errors.Is(err, models.ErrChatNotFound) {
			return &api.DeleteRoleNotFound{}, nil
		}
		if errors.Is(err, models.ErrRoleNotFound) {
			return &api.DeleteRoleNotFound{}, nil
		}
		logger.FromCtx(ctx).Error("delete role", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.DeleteRoleNoContent{}, nil
}

// GetMyRole implements getMyRole operation.
//
// Get my role in chat.
//
// GET /roles/my
func (rh *RolesHandler) GetMyRole(ctx context.Context, params api.GetMyRoleParams) (api.GetMyRoleRes, error) {
	role, err := rh.rolesService.GetRoleForUser(ctx, int(params.ChatId), auth.UserIdFromCtx(ctx))
	if err != nil {
		logger.FromCtx(ctx).Error("get my role", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return mapper.ModelsRoleToApiRole(role), nil
}

// GetRoleById implements getRoleById operation.
//
// Get role in Chat.
//
// GET /roles/{roleId}
func (rh *RolesHandler) GetRoleById(ctx context.Context, params api.GetRoleByIdParams) (api.GetRoleByIdRes, error) {
	role, err := rh.rolesService.GetRoleById(ctx, int(params.ChatId), int(params.RoleId))
	if err != nil {
		if errors.Is(err, models.ErrRoleNotFound) {
			return &api.GetRoleByIdNotFound{}, nil
		}
		logger.FromCtx(ctx).Error("get role by id", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return mapper.ModelsRoleToApiRole(role), nil
}

// ListRoles implements listRoles operation.
//
// Get roles for chat.
//
// GET /roles
func (rh *RolesHandler) ListRoles(ctx context.Context, params api.ListRolesParams) (api.ListRolesRes, error) {
	roles, err := rh.rolesService.ListRoles(ctx, int(params.ChatId))
	if err != nil {
		logger.FromCtx(ctx).Error("list roles", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	apiRoles := make([]api.Role, len(roles))
	for i, role := range roles {
		apiRoles[i] = *mapper.ModelsRoleToApiRole(role)
	}

	res := api.ListRolesOKApplicationJSON(apiRoles)
	return &res, nil
}

// UpdateRole implements updateRole operation.
//
// Update role in chat.
//
// PUT /roles/{roleId}
func (rh *RolesHandler) UpdateRole(ctx context.Context, req *api.RoleInput, params api.UpdateRoleParams) (api.UpdateRoleRes, error) {
	newRole := mapper.ApiRoleInputToModelsRole(req)
	updatedRole, err := rh.rolesService.UpdateRole(ctx, int(params.ChatId), int(params.RoleId), newRole)
	if err != nil {
		if errors.Is(err, models.ErrRoleNotFound) {
			return &api.UpdateRoleNotFound{}, nil
		}
		if errors.Is(err, models.ErrRoleAlreadyExists) {
			return &api.UpdateRoleConflict{}, nil
		}

		logger.FromCtx(ctx).Error("update role", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return mapper.ModelsRoleToApiRole(updatedRole), nil
}
