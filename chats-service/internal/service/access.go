package service

import (
	"context"
	"errors"
	"fmt"
	"go13/chats-service/internal/models"
	"go13/chats-service/internal/repo"
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
	chatsRepo      repo.ChatsRepo
	rolesRepo      repo.RolesRepo
	messagesRepo   repo.MessagesRepo
}

func NewAccessService(
	chatsRepo repo.ChatsRepo,
	rolesRepo repo.RolesRepo,
	messagesRepo repo.MessagesRepo,
) *AccessService {
	return &AccessService{
		chatsServer:    &chapi.Server{},
		messagesServer: &mapi.Server{},
		chatsRepo:      chatsRepo,
		rolesRepo:      rolesRepo,
		messagesRepo:   messagesRepo,
	}
}

func (as *AccessService) CheckAccess(ctx context.Context, userId string, method string, u *url.URL) error {
	op := "AccessService.CheckAccess"

	if route, ok := as.chatsServer.FindPath(method, u); ok {
		logger.FromCtx(ctx).Debug(op, zap.String("operation", route.OperationID()))
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

		case chapi.GetMyRoleOperation:
			return as.CheckGetMyRole(ctx, userId, route, u)

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
		logger.FromCtx(ctx).Debug(op, zap.String("operation", route.OperationID()))
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
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	return nil
}

func (as *AccessService) CheckSendMessage(ctx context.Context, userId string, route mapi.Route, u *url.URL) error {
	op := "AccessService.ChackSendMessage"

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
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	return nil
}

func (as *AccessService) CheckGetMessageById(ctx context.Context, userId string, route mapi.Route, u *url.URL) error {
	op := "AccessService.CheckGetMessageById"

	chatId, err := getChatIdFromQuery(u)
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = getMessageIdFromArgs(route.Args())
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
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	return nil
}

func (as *AccessService) CheckUpdateMessage(ctx context.Context, userId string, route mapi.Route, u *url.URL) error {
	op := "AccessService.CheckUpdateMessage"

	chatId, err := getChatIdFromQuery(u)
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	messageId, err := getMessageIdFromArgs(route.Args())
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
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	message, err := as.messagesRepo.GetMessageById(ctx, chatId, messageId)
	if err != nil {
		return fmt.Errorf("%s: get message by id: %w", op, err)
	}

	if message.SenderId != userId {
		return models.ErrAccessForbidden
	}

	return nil
}

func (as *AccessService) CheckDeleteMessage(ctx context.Context, userId string, route mapi.Route, u *url.URL) error {
	op := "AccessService.CheckDeketeMessage"

	chatId, err := getChatIdFromQuery(u)
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	messageId, err := getMessageIdFromArgs(route.Args())
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	role, err := as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	message, err := as.messagesRepo.GetMessageById(ctx, chatId, messageId)
	if err != nil {
		return fmt.Errorf("%s: get message by id: %w", op, err)
	}

	if message.SenderId != userId && !role.CanDeleteMessages {
		return models.ErrAccessForbidden
	}

	return nil
}

func (as *AccessService) CheckBanUser(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckBanUser"

	chatId, memberId, err := getChatIdAndMemberIdFromArgs(route.Args())
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	role, err := as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	if !role.CanManageMembers {
		return models.ErrAccessForbidden
	}

	bannedRole, err := as.rolesRepo.GetRoleForMember(ctx, chatId, memberId)
	if err != nil {
		return fmt.Errorf("%s: get role for member: %w", op, err)
	}

	if bannedRole.Name == models.RoleCreator.Name {
		return models.ErrAccessForbidden
	}

	return nil
}

func (as *AccessService) CheckCreateChat(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	return nil
}

func (as *AccessService) CheckCreateRole(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckCreateRole"

	chatId, err := getChatIdFromQuery(u)
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	role, err := as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	if !role.CanEditRoles {
		return models.ErrAccessForbidden
	}

	return nil
}

func (as *AccessService) CheckDeleteChat(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckDeleteChat"

	chatId, err := getChatIdFromArgs(route.Args())
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	role, err := as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	if !role.CanDeleteChat {
		return models.ErrAccessForbidden
	}

	return nil
}

func (as *AccessService) CheckDeleteRole(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckDeleteRole"

	chatId, err := getChatIdFromQuery(u)
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	role, err := as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	if !role.CanEditRoles {
		return models.ErrAccessForbidden
	}

	roleId, err := getRoleIdFromArgs(route.Args())
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	targetRole, err := as.rolesRepo.GetRoleById(ctx, chatId, roleId)
	if err != nil {
		return fmt.Errorf("%s: get target role: %w", op, err)
	}

	if targetRole.IsSystem {
		return models.ErrAccessForbidden
	}

	return nil
}

func (as *AccessService) CheckGetChatById(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckGetChatById"

	chatId, err := getChatIdFromArgs(route.Args())
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
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	return nil
}

func (as *AccessService) CheckGetJoinCode(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckGetJoinCode"

	chatId, err := getChatIdFromArgs(route.Args())
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	role, err := as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	if !role.CanGetJoinCode {
		return models.ErrAccessForbidden
	}

	return nil
}

func (as *AccessService) CheckGetMyRole(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessServiceCheckGetMyRole"

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
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	return nil
}

func (as *AccessService) CheckGetRoleById(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckGetRoleById"

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
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	return nil
}

func (as *AccessService) CheckJoinChat(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	return nil
}

func (as *AccessService) CheckLeaveChat(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	return nil
}

func (as *AccessService) CheckListBannedUsers(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckListBannedUsers"

	chatId, err := getChatIdFromArgs(route.Args())
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
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	return nil
}

func (as *AccessService) CheckListChats(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	return nil
}

func (as *AccessService) CheckListMembers(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckListMembers"

	chatId, err := getChatIdFromArgs(route.Args())
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
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	return nil
}

func (as *AccessService) CheckListRoles(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckListRoles"

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
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	return nil
}

func (as *AccessService) CheckSetRole(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckSetRole"

	chatId, memberId, err := getChatIdAndMemberIdFromArgs(route.Args())
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	role, err := as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	if !role.CanManageMembers {
		return models.ErrAccessForbidden
	}

	targetUserRole, err := as.rolesRepo.GetRoleForMember(ctx, chatId, memberId)
	if err != nil {
		return fmt.Errorf("%s: get role for member: %w", op, err)
	}

	if targetUserRole.Name == models.RoleCreator.Name {
		return models.ErrAccessForbidden
	}

	return nil
}

func (as *AccessService) CheckUnbanUser(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckUnbanUser"

	chatId, _, err := getChatIdAndMemberIdFromArgs(route.Args())
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	role, err := as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	if !role.CanManageMembers {
		return models.ErrAccessForbidden
	}

	return nil
}

func (as *AccessService) CheckUpdateChat(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckUpdateChat"

	chatId, err := getChatIdFromArgs(route.Args())
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	role, err := as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	if !role.CanEditChatInfo {
		return models.ErrAccessForbidden
	}

	return nil
}

func (as *AccessService) CheckUpdateRole(ctx context.Context, userId string, route chapi.Route, u *url.URL) error {
	op := "AccessService.CheckListMembers"

	chatId, err := getChatIdFromQuery(u)
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	_, err = as.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	role, err := as.rolesRepo.GetRoleForMember(ctx, chatId, userId)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			return models.ErrAccessForbidden
		}
		return fmt.Errorf("%s: get role for user: %w", op, err)
	}

	if !role.CanEditRoles {
		return models.ErrAccessForbidden
	}

	roleId, err := getRoleIdFromArgs(route.Args())
	if err != nil {
		return models.ErrInvalidRouteParams
	}

	targetRole, err := as.rolesRepo.GetRoleById(ctx, chatId, roleId)
	if err != nil {
		return fmt.Errorf("%s: get target role: %w", op, err)
	}

	if targetRole.Name == models.RoleCreator.Name {
		return models.ErrAccessForbidden
	}

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
	if len(args) < 1 {
		return 0, errors.New("invalid args count")
	}
	messageIdStr := args[0]
	messageId, err := strconv.Atoi(messageIdStr)
	if err != nil {
		return 0, err
	}
	return messageId, nil
}

func getChatIdFromArgs(args []string) (int, error) {
	if len(args) < 1 {
		return 0, errors.New("invalid args count")
	}
	chatIdStr := args[0]
	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		return 0, err
	}
	return chatId, nil
}

func getChatIdAndMemberIdFromArgs(args []string) (int, string, error) {
	if len(args) < 2 {
		return 0, "", errors.New("invalid args count")
	}
	chatId, err := getChatIdFromArgs(args)
	if err != nil {
		return 0, "", err
	}
	memberId := args[1]
	return chatId, memberId, nil
}

func getRoleIdFromArgs(args []string) (int, error) {
	if len(args) < 1 {
		return 0, errors.New("invalid args count")
	}
	roleIdStr := args[0]
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		return 0, err
	}
	return roleId, nil
}
