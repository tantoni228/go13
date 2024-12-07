package service

import (
	"context"
	"fmt"
	"go13/chats-service/internal/models"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
)

type ChatsService struct {
	chatsRepo   ChatsRepo
	rolesRepo   RolesRepo
	membersRepo MembersRepo
	trManager   trm.Manager
}

func NewChatsService(
	chatsRepo ChatsRepo,
	rolesRepo RolesRepo,
	membersRepo MembersRepo,
	trManager trm.Manager,
) *ChatsService {
	return &ChatsService{
		chatsRepo:   chatsRepo,
		rolesRepo:   rolesRepo,
		membersRepo: membersRepo,
		trManager:   trManager,
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
