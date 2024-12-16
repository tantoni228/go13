package service

import (
	"context"
	"errors"
	"fmt"
	"go13/chats-service/internal/models"
	"go13/chats-service/internal/repo"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/golang-jwt/jwt/v5"
)

var (
	chatIdKey = "chat_id"
)

type ChatsService struct {
	chatsRepo      repo.ChatsRepo
	rolesRepo      repo.RolesRepo
	membersRepo    repo.MembersRepo
	joinCodeSecret string
	trManager      trm.Manager
}

func NewChatsService(
	chatsRepo repo.ChatsRepo,
	rolesRepo repo.RolesRepo,
	membersRepo repo.MembersRepo,
	trManager trm.Manager,
	joinCodeSecret string,
) *ChatsService {
	return &ChatsService{
		chatsRepo:      chatsRepo,
		rolesRepo:      rolesRepo,
		membersRepo:    membersRepo,
		trManager:      trManager,
		joinCodeSecret: joinCodeSecret,
	}
}

func (cs *ChatsService) CreateChat(ctx context.Context, creatorId string, chat models.Chat) (models.Chat, error) {
	op := "ChatsService.CreateChat"

	var resChat models.Chat
	err := cs.trManager.Do(ctx, func(ctx context.Context) error {
		createdChat, err := cs.chatsRepo.CreateChat(ctx, chat)
		if err != nil {
			return fmt.Errorf("create chat: %w", err)
		}

		_, err = cs.rolesRepo.CreateRole(ctx, createdChat.Id, models.RoleMember)
		if err != nil {
			return fmt.Errorf("create member role: %w", err)
		}

		creatorRole, err := cs.rolesRepo.CreateRole(ctx, createdChat.Id, models.RoleCreator)
		if err != nil {
			return fmt.Errorf("create creator role: %w", err)
		}

		err = cs.membersRepo.AddMember(ctx, createdChat.Id, models.Member{
			UserId: creatorId,
			RoleId: creatorRole.Id,
		})
		if err != nil {
			return fmt.Errorf("add member: %w", err)
		}

		resChat = createdChat
		return nil
	})

	if err != nil {
		return models.Chat{}, fmt.Errorf("%s: trmanger.Do: %w", op, err)
	}

	return resChat, nil
}

func (cs *ChatsService) ListChatsForUser(ctx context.Context, userId string) ([]models.Chat, error) {
	op := "ChatsService.ListChatsForUser"

	chats, err := cs.chatsRepo.ListChatsForUser(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return chats, nil
}

func (cs *ChatsService) GetChatById(ctx context.Context, chatId int) (models.Chat, error) {
	op := "ChatsService.GetChatById"

	chat, err := cs.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return models.Chat{}, fmt.Errorf("%s: %w", op, err)
	}

	return chat, nil
}

func (cs *ChatsService) UpdateChat(ctx context.Context, chatId int, newChat models.Chat) (models.Chat, error) {
	op := "ChatsService.UpdateChat"

	updatedChat, err := cs.chatsRepo.UpdateChat(ctx, chatId, newChat)
	if err != nil {
		return models.Chat{}, fmt.Errorf("%s: %w", op, err)
	}

	return updatedChat, nil
}

func (cs *ChatsService) DeleteChat(ctx context.Context, chatId int) error {
	op := "ChatsService.DeleteChat"

	err := cs.trManager.Do(ctx, func(ctx context.Context) error {
		if err := cs.membersRepo.DeleteMembersForChat(ctx, chatId); err != nil {
			return fmt.Errorf("delete members: %w", err)
		}
		if err := cs.rolesRepo.DeleteRolesForChat(ctx, chatId); err != nil {
			return fmt.Errorf("delete roles: %w", err)
		}
		if err := cs.chatsRepo.DeleteChat(ctx, chatId); err != nil {
			return fmt.Errorf("delete chat: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%s: trManager.Do: %w", op, err)
	}

	return nil
}

func (cs *ChatsService) ListMembers(ctx context.Context, chatId int) ([]models.Member, error) {
	op := "ChatsService.ListMembers"

	members, err := cs.membersRepo.ListMembers(ctx, chatId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return members, nil
}

func (cs *ChatsService) GetJoinCode(ctx context.Context, chatId int) (string, error) {
	op := "ChatsService.GetJoinCode"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		chatIdKey: chatId,
	})

	signed, err := token.SignedString([]byte(cs.joinCodeSecret))
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return signed, nil
}

func (cs *ChatsService) JoinChat(ctx context.Context, userId string, joinCode string) error {
	op := "ChatsService.JoinChat"

	token, err := jwt.Parse(joinCode, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(cs.joinCodeSecret), nil
	})
	if err != nil {
		return models.ErrInvalidJoinCode
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return models.ErrInvalidJoinCode
	}

	chatIdAny, ok := claims[chatIdKey]
	if !ok {
		return models.ErrInvalidJoinCode
	}

	chatIdFloat64, ok := chatIdAny.(float64)
	if !ok {
		return models.ErrInvalidJoinCode
	}

	chatId := int(chatIdFloat64)
	_, err = cs.chatsRepo.GetChatById(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get chat by id: %w", op, err)
	}

	banned, err := cs.membersRepo.CheckMemberIsBanned(ctx, chatId, userId)
	if err != nil {
		return fmt.Errorf("%s: check member is banned: %w", op, err)
	}

	if banned {
		return models.ErrUserIsBanned
	}

	roleId, err := cs.rolesRepo.GetMemberRoleId(ctx, chatId)
	if err != nil {
		return fmt.Errorf("%s: get member role id: %w", op, err)
	}

	err = cs.membersRepo.AddMember(ctx, chatId, models.Member{UserId: userId, RoleId: roleId})
	if err != nil {
		return fmt.Errorf("%s: add member: %w", op, err)
	}

	return nil
}

func (cs *ChatsService) LeaveChat(ctx context.Context, chatId int, userId string) error {
	op := "ChatsService.Leave"

	err := cs.membersRepo.DeleteMember(ctx, chatId, userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (cs *ChatsService) SetRole(ctx context.Context, chatId int, userId string, roleId int) error {
	op := "ChatsService.SetRole"

	err := cs.membersRepo.SetRoleForMember(ctx, chatId, userId, roleId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (cs *ChatsService) BanUser(ctx context.Context, chatId int, userId string) error {
	op := "ChatsService.BanUser"

	err := cs.trManager.Do(ctx, func(ctx context.Context) error {
		if err := cs.membersRepo.AddMemberToBanned(ctx, chatId, userId); err != nil {
			return fmt.Errorf("add member to banned: %w", err)
		}
		if err := cs.membersRepo.DeleteMember(ctx, chatId, userId); err != nil {
			return fmt.Errorf("delete member: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%s: trmanager.Do: %w", op, err)
	}

	return nil
}

func (cs *ChatsService) UnbanUser(ctx context.Context, chatId int, userId string) error {
	op := "ChatsService.UnbanUser"

	err := cs.trManager.Do(ctx, func(ctx context.Context) error {
		if err := cs.membersRepo.DeleteMemberFromBanned(ctx, chatId, userId); err != nil {
			return fmt.Errorf("delete member from banned: %w", err)
		}
		roleId, err := cs.rolesRepo.GetMemberRoleId(ctx, chatId)
		if err != nil {
			return fmt.Errorf("get member role id: %w", err)
		}
		if err := cs.membersRepo.AddMember(ctx, chatId, models.Member{
			UserId: userId,
			RoleId: roleId,
		}); err != nil {
			return fmt.Errorf("add member: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("%s: trmanager.Do: %w", op, err)
	}

	return nil
}

func (cs *ChatsService) ListBannedMembers(ctx context.Context, chatId int) ([]string, error) {
	op := "ChatsService.ListBannedMembers"

	bannedMembers, err := cs.membersRepo.ListBannedMembers(ctx, chatId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return bannedMembers, nil
}
