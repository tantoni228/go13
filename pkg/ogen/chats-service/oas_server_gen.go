// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	ChatsHandler
	RolesHandler
}

// ChatsHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Chats
type ChatsHandler interface {
	// BanUser implements banUser operation.
	//
	// Ban user in chat.
	//
	// POST /chats/{chatId}/members/{userId}/ban
	BanUser(ctx context.Context, params BanUserParams) (BanUserRes, error)
	// CreateChat implements createChat operation.
	//
	// Create new chat.
	//
	// POST /chats
	CreateChat(ctx context.Context, req *ChatInput) (CreateChatRes, error)
	// DeleteChat implements deleteChat operation.
	//
	// Delete chat by id.
	//
	// DELETE /chats/{chatId}
	DeleteChat(ctx context.Context, params DeleteChatParams) (DeleteChatRes, error)
	// GetChatById implements getChatById operation.
	//
	// Get chat info by id.
	//
	// GET /chats/{chatId}
	GetChatById(ctx context.Context, params GetChatByIdParams) (GetChatByIdRes, error)
	// GetJoinCode implements getJoinCode operation.
	//
	// Get join code for chat by id.
	//
	// GET /chats/{chatId}/join-code
	GetJoinCode(ctx context.Context, params GetJoinCodeParams) (GetJoinCodeRes, error)
	// JoinChat implements joinChat operation.
	//
	// Join chat.
	//
	// POST /chats/join
	JoinChat(ctx context.Context, req *JoinChatReq) (JoinChatRes, error)
	// LeaveChat implements leaveChat operation.
	//
	// Leave chat by id.
	//
	// POST /chats/{chatId}/leave
	LeaveChat(ctx context.Context, params LeaveChatParams) (LeaveChatRes, error)
	// ListChats implements listChats operation.
	//
	// Get chats infos for user.
	//
	// GET /chats
	ListChats(ctx context.Context) (ListChatsRes, error)
	// ListMembers implements listMembers operation.
	//
	// Get members for chat.
	//
	// GET /chats/{chatId}/members
	ListMembers(ctx context.Context, params ListMembersParams) (ListMembersRes, error)
	// SetRole implements setRole operation.
	//
	// Set role to user.
	//
	// POST /chats/{chatId}/members/{userId}/set-role
	SetRole(ctx context.Context, req *SetRoleReq, params SetRoleParams) (SetRoleRes, error)
	// UpdateChat implements updateChat operation.
	//
	// Update chat info.
	//
	// PUT /chats/{chatId}
	UpdateChat(ctx context.Context, req *ChatInput, params UpdateChatParams) (UpdateChatRes, error)
}

// RolesHandler handles operations described by OpenAPI v3 specification.
//
// x-ogen-operation-group: Roles
type RolesHandler interface {
	// CheckAccess implements CheckAccess operation.
	//
	// Check access to uri with method.
	//
	// GET /roles/check-access
	CheckAccess(ctx context.Context, params CheckAccessParams) (CheckAccessRes, error)
	// CreateRole implements createRole operation.
	//
	// Create role in Chat.
	//
	// POST /roles
	CreateRole(ctx context.Context, req *RoleInput, params CreateRoleParams) (CreateRoleRes, error)
	// DeleteRole implements deleteRole operation.
	//
	// Delete role in chat.
	//
	// DELETE /roles/{roleId}
	DeleteRole(ctx context.Context, params DeleteRoleParams) (DeleteRoleRes, error)
	// GetRoleById implements getRoleById operation.
	//
	// Get role in Chat.
	//
	// GET /roles/{roleId}
	GetRoleById(ctx context.Context, params GetRoleByIdParams) (GetRoleByIdRes, error)
	// ListRoles implements listRoles operation.
	//
	// Get roles for chat.
	//
	// GET /roles
	ListRoles(ctx context.Context, params ListRolesParams) (ListRolesRes, error)
	// UpdateRole implements updateRole operation.
	//
	// Update role in chat.
	//
	// PUT /roles/{roleId}
	UpdateRole(ctx context.Context, req *RoleInput, params UpdateRoleParams) (UpdateRoleRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}