package handlers

import (
	"context"
	api "go13/pkg/ogen/chats-service"
)

type RolesHandler struct {
}

func NewRolesHandler() *RolesHandler {
	return &RolesHandler{}
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
	return &api.Role{}, nil
}

// DeleteRole implements deleteRole operation.
//
// Delete role in chat.
//
// DELETE /roles/{roleId}
func (rh *RolesHandler) DeleteRole(ctx context.Context, params api.DeleteRoleParams) (api.DeleteRoleRes, error) {
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
