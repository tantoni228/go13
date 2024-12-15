package service

import (
	"context"
	"errors"
	"fmt"
	"go13/chats-service/internal/models"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/golang-jwt/jwt/v5"
)

var (
	chatIdKey = "chat_id"
)

type ChatsService struct {
	chatsRepo      ChatsRepo
	rolesRepo      RolesRepo
	membersRepo    MembersRepo
	joinCodeSecret string
	trManager      trm.Manager
}

func NewChatsService(
	chatsRepo ChatsRepo,
	rolesRepo RolesRepo,
	membersRepo MembersRepo,
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
