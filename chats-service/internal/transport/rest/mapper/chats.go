package mapper

import (
	"go13/chats-service/internal/models"
	api "go13/pkg/ogen/chats-service"
)

func ModelsChatToApiChat(chat models.Chat) *api.Chat {
	return &api.Chat{
		ID:          api.ChatId(chat.Id),
		Name:        chat.Name,
		Description: chat.Description,
	}
}

func ApiChatInputToModelsChat(chatInput *api.ChatInput) models.Chat {
	return models.Chat{
		Name:        chatInput.Name,
		Description: chatInput.Description,
	}
}
