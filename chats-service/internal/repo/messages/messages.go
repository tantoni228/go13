package messages

import (
	"context"
	"fmt"
	"go13/chats-service/internal/models"
	api "go13/pkg/ogen/messages-service"
)

type securitySource struct {
	Token string
}

func (ss *securitySource) BearerAuth(ctx context.Context, operationName api.OperationName) (api.BearerAuth, error) {
	return api.BearerAuth{
		Token: ss.Token,
	}, nil
}

type Config struct {
	Host  string `yaml:"host" env:"MESSAGES_HOST" env-default:"localhost"`
	Port  int    `yaml:"port" env:"MESSAGES_PORT" env-default:"8080"`
	Token string `yaml:"token" env:"MESSAGES_TOKEN"`
}

type MessagesRepo struct {
	cli *api.Client
}

func NewMessagesRepo(cfg Config) (*MessagesRepo, error) {
	cli, err := api.NewClient(fmt.Sprintf("http://%s:%d", cfg.Host, cfg.Port), &securitySource{Token: cfg.Token})
	if err != nil {
		return nil, err
	}

	return &MessagesRepo{
		cli: cli,
	}, nil
}

func (mr *MessagesRepo) GetMessageById(ctx context.Context, chatId int, messageId int) (models.Message, error) {
	op := "MessagesRepo.GetMessageById"

	resp, err := mr.cli.GetMessageById(ctx, api.GetMessageByIdParams{
		ChatId:    api.ChatId(chatId),
		MessageId: api.MessageId(messageId),
	})

	if err != nil {
		return models.Message{}, fmt.Errorf("%s: %w", op, err)
	}

	switch resp.(type) {
	case *api.InvalidInputResponse:
		return models.Message{}, fmt.Errorf("%s: invalid input response", op)
	case *api.UnauthenticatedResponse:
		return models.Message{}, fmt.Errorf("%s: unaunthenticated response", op)
	case *api.UnauthorizedResponse:
		return models.Message{}, fmt.Errorf("%s: unauthorized response", op)
	case *api.GetMessageByIdNotFound:
		return models.Message{}, models.ErrMessageNotFound
	case *api.InternalErrorResponse:
		return models.Message{}, fmt.Errorf("%s: internal error response", op)
	}

	messageResp := resp.(*api.Message)
	return models.Message{
		Id:            int(messageResp.GetID()),
		SenderId:      string(messageResp.GetSenderID()),
		Message:       messageResp.GetMessage(),
		Edited:        messageResp.GetEdited(),
		SendTimestamp: messageResp.GetSendTimestamp(),
	}, nil
}
