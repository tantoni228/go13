package handlers

import (
	"context"
	api "go13/pkg/ogen/chats-service"
)

type ChatsHandler struct {
}

func NewChatsHandler() *ChatsHandler {
	return &ChatsHandler{}
}

// BanUser implements banUser operation.
//
// Ban user in chat.
//
// POST /chats/{chatId}/members/{userId}/ban
func (ch *ChatsHandler) BanUser(ctx context.Context, params api.BanUserParams) (api.BanUserRes, error) {
	return &api.BanUserNoContent{}, nil
}

// CreateChat implements createChat operation.
//
// Create new chat.
//
// POST /chats
func (ch *ChatsHandler) CreateChat(ctx context.Context, req *api.ChatInput) (api.CreateChatRes, error) {
	return &api.Chat{}, nil
}

// DeleteChat implements deleteChat operation.
//
// Delete chat by id.
//
// DELETE /chats/{chatId}
func (ch *ChatsHandler) DeleteChat(ctx context.Context, params api.DeleteChatParams) (api.DeleteChatRes, error) {
	return &api.DeleteChatNoContent{}, nil
}

// GetChatById implements getChatById operation.
//
// Get chat info by id.
//
// GET /chats/{chatId}
func (ch *ChatsHandler) GetChatById(ctx context.Context, params api.GetChatByIdParams) (api.GetChatByIdRes, error) {
	return &api.Chat{}, nil
}

// GetJoinCode implements getJoinCode operation.
//
// Get join code for chat by id.
//
// GET /chats/{chatId}/join-code
func (ch *ChatsHandler) GetJoinCode(ctx context.Context, params api.GetJoinCodeParams) (api.GetJoinCodeRes, error) {
	return &api.JoinCodeResponse{}, nil
}

// JoinChat implements joinChat operation.
//
// Join chat.
//
// POST /chats/join
func (ch *ChatsHandler) JoinChat(ctx context.Context, req *api.JoinChatReq) (api.JoinChatRes, error) {
	return &api.JoinChatNoContent{}, nil
}

// LeaveChat implements leaveChat operation.
//
// Leave chat by id.
//
// POST /chats/{chatId}/leave
func (ch *ChatsHandler) LeaveChat(ctx context.Context, params api.LeaveChatParams) (api.LeaveChatRes, error) {
	return &api.LeaveChatNoContent{}, nil
}

// ListChats implements listChats operation.
//
// Get chats infos for user.
//
// GET /chats
func (ch *ChatsHandler) ListChats(ctx context.Context) (api.ListChatsRes, error) {
	return &api.ListChatsOKApplicationJSON{}, nil
}

// ListMembers implements listMembers operation.
//
// Get members for chat.
//
// GET /chats/{chatId}/members
func (ch *ChatsHandler) ListMembers(ctx context.Context, params api.ListMembersParams) (api.ListMembersRes, error) {
	return &api.ListMembersOKApplicationJSON{}, nil
}

// SetRole implements setRole operation.
//
// Set role to user.
//
// POST /chats/{chatId}/members/{userId}/set-role
func (ch *ChatsHandler) SetRole(ctx context.Context, req *api.SetRoleReq, params api.SetRoleParams) (api.SetRoleRes, error) {
	return &api.SetRoleNoContent{}, nil
}

// UpdateChat implements updateChat operation.
//
// Update chat info.
//
// PUT /chats/{chatId}
func (ch *ChatsHandler) UpdateChat(ctx context.Context, req *api.ChatInput, params api.UpdateChatParams) (api.UpdateChatRes, error) {
	return &api.Chat{}, nil
}
