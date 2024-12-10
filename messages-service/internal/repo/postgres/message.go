package postgres

import (
	"context"
	"fmt"
	// "log"

	// "go13/messages-service/internal/models"
	api "go13/pkg/ogen/messages-service"
	"go13/pkg/postgres"
	"go13/messages-service/internal/models"

	sq "github.com/Masterminds/squirrel"
	// "github.com/google/uuid"
)


type MessageRepository struct {
	db *postgres.Postgres
}

func NewMessageRepository(db *postgres.Postgres) *MessageRepository {
	return &MessageRepository{db: db}
}

// func GenerateRandomUUID() (string) {
// 	id := uuid.New()
//     return id.String()
// }

func (s MessageRepository) SendMessage(ctx context.Context, req *api.MessageInput, params api.SendMessageParams) (api.SendMessageRes, error) {
	var chat_id models.Message
	var result api.Message
	err := sq.Insert("messages").
		Columns("message", "edited", "user_id", "send_timestamp", "chat_id").
		Values(&req.Message, false, "74ddd237-507a-4fad-a20a-4b98a43d7ffb", 5, params.ChatId).
		Suffix("returning *").
		PlaceholderFormat(sq.Dollar).
		RunWith(s.db.DB.DB).
		QueryRow().
		Scan(&result.ID, &result.SenderID, &chat_id.Chat_id, &result.Message, &result.Edited, &result.SendTimestamp)
	if err != nil {
		return nil, fmt.Errorf("repository.SendMessage: %w", err)
	}
	return &result, nil
}