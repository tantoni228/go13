package postgres

import (
	"context"
	"fmt"
	"go13/chats-service/internal/models"
	"go13/pkg/postgres"

	"github.com/Masterminds/squirrel"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type RolesRepo struct {
	db     *sqlx.DB
	sq     squirrel.StatementBuilderType
	getter *trmsqlx.CtxGetter
}

func NewRolesRepo(pg *postgres.Postgres) *RolesRepo {
	return &RolesRepo{
		db:     pg.DB,
		sq:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		getter: trmsqlx.DefaultCtxGetter,
	}
}

func (rr *RolesRepo) CreateRole(ctx context.Context, chatId int, role models.Role) (models.Role, error) {
	op := "RolesRepo.CreateRole"
	sql, args, err := rr.sq.
		Insert("roles").
		Columns(
			"chat_id",
			"name",
			"is_system",
			"can_ban_users",
			"can_edit_roles",
			"can_delete_messages",
			"can_get_join_code",
			"can_edit_chat_info",
			"can_delete_chat",
		).
		Values(
			chatId,
			role.Name,
			role.IsSystem,
			role.CanBanUsers,
			role.CanEditRoles,
			role.CanDeleteMessages,
			role.CanGetJoinCode,
			role.CanEditChatInfo,
			role.CanDeleteChat,
		).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return models.Role{}, fmt.Errorf("%s: build query: %w", op, err)
	}

	var id int
	err = rr.getter.DefaultTrOrDB(ctx, rr.db).QueryRowContext(ctx, sql, args...).Scan(&id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23503":
				return models.Role{}, models.ErrChatNotFound
			case "23505":
				return models.Role{}, models.ErrRoleAlreadyExists
			}
		}
		return models.Role{}, fmt.Errorf("%s: QueryRowContext: %w", op, err)
	}

	role.Id = id
	return role, nil
}
