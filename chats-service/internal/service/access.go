package service

import (
	"context"
	"errors"
	"fmt"
	"go13/chats-service/internal/models"
	"go13/pkg/logger"
	chapi "go13/pkg/ogen/chats-service"
	mapi "go13/pkg/ogen/messages-service"
	"net/url"
	"strconv"

	"go.uber.org/zap"
)

type AccessService struct {
	chatsServer    *chapi.Server
	messagesServer *mapi.Server
	chatsRepo      ChatsRepo
	rolesRepo      RolesRepo
}

func NewAccessService(chatsRepo ChatsRepo, rolesRepo RolesRepo) *AccessService {
	return &AccessService{
		chatsServer:    &chapi.Server{},
		messagesServer: &mapi.Server{},
		chatsRepo:      chatsRepo,
		rolesRepo:      rolesRepo,
	}
}

func (as *AccessService) CheckAccess(ctx context.Context, userId string, method string, u *url.URL) error {
	if route, ok := as.chatsServer.FindPath(method, u); ok {
		switch route.Name() {
		case chapi.BanUserOperation:
			return as.CheckBanUser(ctx, userId, route, u)

		case chapi.CheckAccessOperation:
			return nil

		case chapi.CreateChatOperation:
			return as.CheckCreateChat(ctx, userId, route, u)

		case chapi.CreateRoleOperation:
			return as.CheckCreateRole(ctx, userId, route, u)

		case chapi.DeleteChatOperation:
			return as.CheckDeleteChat(ctx, userId, route, u)

		case chapi.DeleteRoleOperation:
			return as.CheckDeleteRole(ctx, userId, route, u)

		case chapi.GetChatByIdOperation:
			return as.CheckGetChatById(ctx, userId, route, u)

		case chapi.GetJoinCodeOperation:
			return as.CheckGetJoinCode(ctx, userId, route, u)

		case chapi.GetRoleByIdOperation:
			return as.CheckGetRoleById(ctx, userId, route, u)

		case chapi.JoinChatOperation:
			return as.CheckJoinChat(ctx, userId, route, u)

		case chapi.LeaveChatOperation:
			return as.CheckLeaveChat(ctx, userId, route, u)

		case chapi.ListBannedUsersOperation:
			return as.CheckListBannedUsers(ctx, userId, route, u)

		case chapi.ListChatsOperation:
			return as.CheckListChats(ctx, userId, route, u)

		case chapi.ListMembersOperation:
			return as.CheckListMembers(ctx, userId, route, u)

		case chapi.ListRolesOperation:
			return as.CheckListRoles(ctx, userId, route, u)

		case chapi.SetRoleOperation:
			return as.CheckSetRole(ctx, userId, route, u)

		case chapi.UnbanUserOperation:
			return as.CheckUnbanUser(ctx, userId, route, u)

		case chapi.UpdateChatOperation:
			return as.CheckUpdateChat(ctx, userId, route, u)

		case chapi.UpdateRoleOperation:
			return as.CheckUpdateRole(ctx, userId, route, u)

		}
	}
	if route, ok := as.messagesServer.FindPath(method, u); ok {
		switch route.Name() {
		case mapi.ListMessagesOperation:
			return as.CheckListMessages(ctx, userId, route, u)
		case mapi.SendMessageOperation:
			return as.CheckSendMessage(ctx, userId, route, u)
		case mapi.GetMessageByIdOperation:
			return as.CheckGetMessageById(ctx, userId, route, u)
		case mapi.UpdateMessageOperation:
			return as.CheckUpdateMessage(ctx, userId, route, u)
		case mapi.DeleteMessageOperation:
			return as.CheckDeleteMessage(ctx, userId, route, u)
		}
	}
	return models.ErrEndpointNotFound
}

func (as *AccessService) CheckListMessages(ctx context.Context, userId string, route mapi.Route, u *url.URL) error {
	op := "AccessService.CheckListMessages"

	logger.FromCtx(ctx).Debug(op, zap.String("operation", route.OperationID()))
	chatId, err := getChatIdFromQuery(u)
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	_, err = as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return err
	}

	return nil
}

func (as *AccessService) CheckSendMessage(ctx context.Context, userId string, route mapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckGetMessageById(ctx context.Context, userId string, route mapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckUpdateMessage(ctx context.Context, userId string, route mapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckDeleteMessage(ctx context.Context, userId string, route mapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckBanUser(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckCreateChat(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckCreateRole(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckDeleteChat(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckDeleteRole(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckGetChatById(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckGetJoinCode(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckGetRoleById(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckJoinChat(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckLeaveChat(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckListBannedUsers(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckListChats(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckListMembers(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckListRoles(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckSetRole(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckUnbanUser(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckUpdateChat(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func (as *AccessService) CheckUpdateRole(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	logger.FromCtx(ctx).Info("check", zap.String("operation", route.OperationID()))
	return nil
}

func getChatIdFromQuery(u *url.URL) (int, error) {
	chatIdStr := u.Query().Get("chatId")
	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		return 0, err
	}

	return chatId, nil
}

func getMessageIdFromArgs(args []string) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("invalid args count")
	}
	messageIdStr := args[0]
	messageId, err := strconv.Atoi(messageIdStr)
	if err != nil {
		return 0, err
	}
	return messageId, nil
}
