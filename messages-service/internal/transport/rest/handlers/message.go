package handlers

import (
	"context"
	"fmt"
	// "go13/messages-service/internal/models"
	"go13/messages-service/internal/service"
	api "go13/pkg/ogen/messages-service"

)

type MessagesHandler struct {

	service *service.MessagesService
}

func NewMessagesHandler(srv *service.MessagesService) *MessagesHandler {
	return &MessagesHandler{service: srv}
}

// DeleteMessage implements deleteMessage operation.
//
// Delete message from chat.
//
// DELETE /messages/{messageId}
func (m *MessagesHandler) DeleteMessage(ctx context.Context, params api.DeleteMessageParams) (api.DeleteMessageRes, error) {
	resp, err := m.service.DeleteMessage(ctx, api.DeleteMessageParams{MessageId: params.MessageId, ChatId: params.ChatId})
	if err != nil {
		return nil, fmt.Errorf("DeleteMessage: %w", err)
	}

	return resp, nil
}
// GetMessageById implements getMessageById operation.
//
// Get message in chat.
//
// GET /messages/{messageId}
func (m *MessagesHandler)  GetMessageById(ctx context.Context, params api.GetMessageByIdParams) (api.GetMessageByIdRes, error) {
	resp, err := m.service.GetMessageById(ctx, api.GetMessageByIdParams{MessageId: params.MessageId,ChatId: params.ChatId})
	if err != nil {
		return nil, fmt.Errorf("GetMessageById: %w", err)
	}

	return resp, nil
}
// ListMessages implements listMessages operation.
//
// Get messages for chat.
//
// GET /messages
func (m *MessagesHandler)  ListMessages(ctx context.Context, params api.ListMessagesParams) (api.ListMessagesRes, error) {
	resp, err := m.service.ListMessages(ctx, api.ListMessagesParams{ChatId: params.ChatId, Limit: params.Limit, Offset: params.Offset})
	if err != nil {
		return nil, fmt.Errorf("ListMessages: %w", err)
	}

	return resp, nil
}
// SendMessage implements sendMessage operation.
//
// Send new message to chat.
//
// POST /messages
func (m *MessagesHandler) SendMessage(ctx context.Context, req *api.MessageInput, params api.SendMessageParams) (api.SendMessageRes, error) {
    resp, err := m.service.SendMessage(ctx, &api.MessageInput{Message: req.Message}, api.SendMessageParams{ChatId: params.ChatId})
    if err != nil {
        return nil, fmt.Errorf("SendMessage: %w", err)
    }

    // Приведение типа resp к messageResponse
    response, ok := resp.(*api.Message)
    if !ok {
        return nil, fmt.Errorf("unexpected type for resp: %T", resp)
    }

    return &api.Message{
        ID:        response.ID,
        SenderID:  response.SenderID,
        Message:   response.Message,
    }, nil
}
// UpdateMessage implements updateMessage operation.
//
// Update message in chat.
//
// PUT /messages/{messageId}
func (m *MessagesHandler)  UpdateMessage(ctx context.Context, req *api.MessageInput, params api.UpdateMessageParams) (api.UpdateMessageRes, error) {
	resp, err := m.service.UpdateMessage(ctx, &api.MessageInput{Message: req.Message}, api.UpdateMessageParams{MessageId: params.MessageId, ChatId: params.ChatId})
	if err != nil {
		return nil, fmt.Errorf("UpdateMessage: %w", err)
	}

	return resp, nil
}