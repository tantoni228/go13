package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	// "log"

	// "go13/messages-service/internal/models"
	"go13/messages-service/internal/models"
	"go13/messages-service/internal/transport/rest/auth"
	api "go13/pkg/ogen/messages-service"
	"go13/pkg/postgres"

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

	fmt.Println(auth.UserIdFromCtx(ctx))

	err := sq.Insert("messages").
		Columns("message", "edited", "user_id", "send_timestamp", "chat_id").
		Values(&req.Message, false, auth.UserIdFromCtx(ctx), 5, params.ChatId).
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

func (s MessageRepository) DeleteMessage(ctx context.Context, params api.DeleteMessageParams) (error) {
	deleteBuilder := sq.Delete("messages").
		Where(sq.Eq{"id": params.MessageId, "chat_id": params.ChatId}).
		PlaceholderFormat(sq.Dollar)
	
	_, err := deleteBuilder.RunWith(s.db.DB.DB).Exec()
	if err != nil {
		return fmt.Errorf("repository.DeleteMessage")
	}
	return  nil
}

func (s MessageRepository) GetMessageById(ctx context.Context, params api.GetMessageByIdParams) (api.GetMessageByIdRes, error) {
	var result api.Message
	
	err := sq.Select("id", "user_id", "message", "edited", "send_timestamp").
        From("messages").
        Where(sq.Eq{"id": params.MessageId}).

		PlaceholderFormat(sq.Dollar).
		RunWith(s.db.DB.DB).
		QueryRow().
		Scan(&result.ID, &result.SenderID, &result.Message, &result.Edited, &result.SendTimestamp)
	
	if err != nil {
		return nil, fmt.Errorf("repository.GetMessageById: %w", err)
	}

	return &result, nil
}

func (s MessageRepository) ListMessages(ctx context.Context, params api.ListMessagesParams) ([]*api.Message, error) {
	var result []*api.Message
	var chat_id int64

    query := sq.Select("*").
		From("messages").
		Where(sq.Eq{"chat_id": params.ChatId}).
		Limit(uint64(params.Limit)).
		Offset(uint64(params.Offset)).
		PlaceholderFormat(sq.Dollar)

    rows, err := query.RunWith(s.db.DB.DB).Query()
    if err != nil {
        return nil, fmt.Errorf("repository.ListMessages: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
		var message api.Message
        if err := rows.Scan(&message.ID, 
			&message.SenderID, 
			&chat_id, 
			&message.Message, 
			&message.Edited,
			&message.SendTimestamp); err != nil {
            return nil, fmt.Errorf("repository.ListMessages: failed to scan row: %w", err)
        }
        result = append(result, &message)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("repository.ListMessages: failed to iterate over rows: %w", err)
    }

    return result, nil
}


func (s MessageRepository) UpdateMessage(ctx context.Context, req *api.MessageInput, params api.UpdateMessageParams) (api.UpdateMessageRes, error) {
	var result api.Message
	var chat_id int64
	err := sq.Update("messages").
        Set("message", req.Message).
        Set("edited", true).
        Where(sq.Eq{"id": params.MessageId, "chat_id": params.ChatId}).
        Suffix("returning *").
        PlaceholderFormat(sq.Dollar).
        RunWith(s.db.DB.DB).
        QueryRow().
        Scan(&result.ID, &result.SenderID, &chat_id, &result.Message, &result.Edited, &result.SendTimestamp)

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) { // Проверка на отсутствие записей
            return nil, fmt.Errorf("repository.UpdateMessage: no row found with ID %s", params.MessageId)
        }
        return nil, fmt.Errorf("repository.UpdateMessage: %w", err)
    }

    return &result, nil
}

