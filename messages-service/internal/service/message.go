package service

import (
	"context"
	api "go13/pkg/ogen/messages-service"
)

type MessagesRepo interface {
	DeleteMessage(ctx context.Context, params api.DeleteMessageParams) (error)
	GetMessageById(ctx context.Context, params api.GetMessageByIdParams) (api.GetMessageByIdRes, error)
	ListMessages(ctx context.Context, params api.ListMessagesParams) ([]*api.Message, error)
	SendMessage(ctx context.Context, req *api.MessageInput, params api.SendMessageParams) (api.SendMessageRes, error)
	UpdateMessage(ctx context.Context, req *api.MessageInput, params api.UpdateMessageParams) (api.UpdateMessageRes, error)
}

type MessagesService struct {
	Repo MessagesRepo
}

func NewMessageService(repo MessagesRepo) *MessagesService {
	return &MessagesService{repo}
}

func (s *MessagesService) DeleteMessage(ctx context.Context, params api.DeleteMessageParams) (error) {
	return s.Repo.DeleteMessage(ctx, params)
}

func (s *MessagesService) GetMessageById(ctx context.Context, params api.GetMessageByIdParams) (api.GetMessageByIdRes, error) {
	return s.Repo.GetMessageById(ctx, params)
}

func (s *MessagesService) ListMessages(ctx context.Context, params api.ListMessagesParams) ([]*api.Message, error) {
	return s.Repo.ListMessages(ctx, params)
}

func (s *MessagesService) SendMessage(ctx context.Context, req *api.MessageInput, params api.SendMessageParams) (api.SendMessageRes, error) {
	return s.Repo.SendMessage(ctx, req, params)
}

func (s *MessagesService) UpdateMessage(ctx context.Context, req *api.MessageInput, params api.UpdateMessageParams) (api.UpdateMessageRes, error) {
	return s.Repo.UpdateMessage(ctx, req, params)
}
