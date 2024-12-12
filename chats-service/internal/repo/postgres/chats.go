package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go13/chats-service/internal/models"
	"go13/pkg/postgres"

	"github.com/Masterminds/squirrel"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
)

type ChatsRepo struct {
	db     *sqlx.DB
	sq     squirrel.StatementBuilderType
	getter *trmsqlx.CtxGetter
}

func NewChatsRepo(pg *postgres.Postgres) *ChatsRepo {
	return &ChatsRepo{
		db:     pg.DB,
		sq:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		getter: trmsqlx.DefaultCtxGetter,
	}
}

func (cr *ChatsRepo) CreateChat(ctx context.Context, chat models.Chat) (models.Chat, error) {
	op := "ChatsRepo.CreateChat"
	sql, args, err := cr.sq.
		Insert("chats").
		Columns("name", "description").
		Values(chat.Name, chat.Description).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return models.Chat{}, fmt.Errorf("%s: build query: %w", op, err)
	}

	var id int
	err = cr.getter.DefaultTrOrDB(ctx, cr.db).QueryRowContext(ctx, sql, args...).Scan(&id)
	if err != nil {
		return models.Chat{}, fmt.Errorf("%s: QueryRowContext: %w", op, err)
	}

	chat.Id = id
	return chat, nil
}

func (cr *ChatsRepo) GetChatById(ctx context.Context, chatId int) (models.Chat, error) {
	op := "ChatsRepo.GetChatById"

	query, args, err := cr.sq.
		Select("id", "name", "description").
		From("chats").
		Where(squirrel.Eq{"id": chatId}).
		ToSql()
	if err != nil {
		return models.Chat{}, fmt.Errorf("%s: build query: %w", op, err)
	}

	var chat models.Chat
	err = cr.getter.DefaultTrOrDB(ctx, cr.db).GetContext(ctx, &chat, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Chat{}, models.ErrChatNotFound
		}

		return models.Chat{}, fmt.Errorf("%s: GetContext: %w", op, err)
	}

	return chat, nil
}

func (cr *ChatsRepo) DeleteChat(ctx context.Context, chatId int) error {
	op := "ChatsRepo.DeleteChat"

	sql, args, err := cr.sq.
		Delete("chats").
		Where(squirrel.Eq{"id": chatId}).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s: building query: %w", op, err)
	}

	cmd, err := cr.getter.DefaultTrOrDB(ctx, cr.db).ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("%s: ExecContext: %w", op, err)
	}

	affected, err := cmd.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: cmd.RowsAffected: %w", op, err)
	}

	if affected == 0 {
		return models.ErrChatNotFound
	}

	return nil
}
