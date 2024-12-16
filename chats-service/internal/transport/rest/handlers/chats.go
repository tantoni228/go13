package handlers

import (
	"context"
	"errors"
	"go13/chats-service/internal/models"
	"go13/chats-service/internal/transport/rest/auth"
	"go13/pkg/logger"
	api "go13/pkg/ogen/chats-service"

	"go.uber.org/zap"
)

type ChatsService interface {
	CreateChat(ctx context.Context, creatorId string, chat models.Chat) (models.Chat, error)
	DeleteChat(ctx context.Context, chatId int) error
	GetJoinCode(ctx context.Context, chatId int) (string, error)
	JoinChat(ctx context.Context, userId string, joinCode string) error
	LeaveChat(ctx context.Context, chatId int, userId string) error
	SetRole(ctx context.Context, chatId int, userId string, roleId int) error
	BanUser(ctx context.Context, chatId int, userId string) error
	UnbanUser(ctx context.Context, chatId int, userId string) error
	ListBannedMembers(ctx context.Context, chatId int) ([]string, error)
}

type ChatsHandler struct {
	chatsService ChatsService
}

func NewChatsHandler(chatsService ChatsService) *ChatsHandler {
	return &ChatsHandler{
		chatsService: chatsService,
	}
}

// BanUser implements banUser operation.
//
// Ban user in chat.
//
// POST /chats/{chatId}/members/{userId}/ban
func (ch *ChatsHandler) BanUser(ctx context.Context, params api.BanUserParams) (api.BanUserRes, error) {
	err := ch.chatsService.BanUser(ctx, int(params.ChatId), string(params.UserId))
	if err != nil {
		if errors.Is(err, models.ErrChatNotFound) {
			return &api.BanUserNotFound{}, nil
		}
		if errors.Is(err, models.ErrMemberNotFound) {
			return &api.BanUserNotFound{}, nil
		}
		if errors.Is(err, models.ErrUserAlreadyBanned) {
			return &api.BanUserConflict{}, nil
		}
		logger.FromCtx(ctx).Error("ban user", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.BanUserNoContent{}, nil
}

// UnbanUser implements UnbanUser operation.
//
// Unban user in chat.
//
// POST /chats/{chatId}/members/{userId}/unban
func (ch *ChatsHandler) UnbanUser(ctx context.Context, params api.UnbanUserParams) (api.UnbanUserRes, error) {
	err := ch.chatsService.UnbanUser(ctx, int(params.ChatId), string(params.UserId))
	if err != nil {
		if errors.Is(err, models.ErrChatNotFound) {
			return &api.UnbanUserNotFound{}, nil
		}
		if errors.Is(err, models.ErrMemberNotFound) {
			return &api.UnbanUserConflict{}, nil
		}

		logger.FromCtx(ctx).Error("unban user", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.UnbanUserNoContent{}, nil
}

// ListBannedUsers implements listBannedUsers operation.
//
// Get banned members for chat.
//
// GET /chats/{chatId}/members/banned
func (ch *ChatsHandler) ListBannedUsers(ctx context.Context, params api.ListBannedUsersParams) (api.ListBannedUsersRes, error) {
	bannedMembers, err := ch.chatsService.ListBannedMembers(ctx, int(params.ChatId))
	if err != nil {
		logger.FromCtx(ctx).Error("list banned members", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	apiBannedMember := make([]api.BannedMembersResponseItem, len(bannedMembers))
	for i, memberId := range bannedMembers {
		apiBannedMember[i] = api.BannedMembersResponseItem{UserID: api.UserId(memberId)}
	}
	resp := api.ListBannedUsersOKApplicationJSON(apiBannedMember)
	return &resp, err
}

// CreateChat implements createChat operation.
//
// Create new chat.
//
// POST /chats
func (ch *ChatsHandler) CreateChat(ctx context.Context, req *api.ChatInput) (api.CreateChatRes, error) {
	userId := auth.UserIdFromCtx(ctx)
	chat := models.Chat{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}
	chat, err := ch.chatsService.CreateChat(ctx, userId, chat)
	if err != nil {
		logger.FromCtx(ctx).Error("ChatsHandler.CreateChat", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.Chat{ID: api.ChatId(chat.Id), Name: chat.Name, Description: chat.Description}, nil
}

// DeleteChat implements deleteChat operation.
//
// Delete chat by id.
//
// DELETE /chats/{chatId}
func (ch *ChatsHandler) DeleteChat(ctx context.Context, params api.DeleteChatParams) (api.DeleteChatRes, error) {
	err := ch.chatsService.DeleteChat(ctx, int(params.ChatId))
	if err != nil {
		if errors.Is(err, models.ErrChatNotFound) {
			return &api.ChatNotFoundResponse{}, nil
		}
		logger.FromCtx(ctx).Error("ChatsHandler.DeleteChat", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

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
	joinCode, err := ch.chatsService.GetJoinCode(ctx, int(params.ChatId))
	if err != nil {
		logger.FromCtx(ctx).Error("get join code", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}
	return &api.JoinCodeResponse{JoinCode: joinCode}, nil
}

// JoinChat implements joinChat operation.
//
// Join chat.
//
// POST /chats/join
func (ch *ChatsHandler) JoinChat(ctx context.Context, req *api.JoinChatReq) (api.JoinChatRes, error) {
	err := ch.chatsService.JoinChat(ctx, auth.UserIdFromCtx(ctx), req.GetJoinCode())
	if err != nil {
		if errors.Is(err, models.ErrInvalidJoinCode) {
			return &api.InvalidInputResponse{
				Message: "invalid join code",
			}, nil
		}
		if errors.Is(err, models.ErrChatNotFound) {
			return &api.ChatNotFoundResponse{}, nil
		}
		if errors.Is(err, models.ErrUserIsBanned) {
			return &api.UnauthorizedResponse{}, nil
		}
		if errors.Is(err, models.ErrUserAlreadyInChat) {
			return &api.JoinChatConflict{}, nil
		}

		logger.FromCtx(ctx).Error("join chat", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

	return &api.JoinChatNoContent{}, nil
}

// LeaveChat implements leaveChat operation.
//
// Leave chat by id.
//
// POST /chats/{chatId}/leave
func (ch *ChatsHandler) LeaveChat(ctx context.Context, params api.LeaveChatParams) (api.LeaveChatRes, error) {
	err := ch.chatsService.LeaveChat(ctx, int(params.ChatId), auth.UserIdFromCtx(ctx))
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return &api.ChatNotFoundResponse{}, nil
		}
		logger.FromCtx(ctx).Error("leave chat", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

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
	err := ch.chatsService.SetRole(ctx, int(params.ChatId), string(params.UserId), int(req.GetRoleID()))
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return &api.SetRoleNotFound{}, nil
		}
		if errors.Is(err, models.ErrRoleNotFound) {
			return &api.SetRoleNotFound{}, nil
		}

		logger.FromCtx(ctx).Error("set role", zap.Error(err))
		return &api.InternalErrorResponse{}, nil
	}

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
