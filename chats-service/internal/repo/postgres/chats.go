package postgres

import (
	"context"
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
