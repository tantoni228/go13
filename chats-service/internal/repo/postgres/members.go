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

type MembersRepo struct {
	db     *sqlx.DB
	sq     squirrel.StatementBuilderType
	getter *trmsqlx.CtxGetter
}

func NewMembersRepo(pg *postgres.Postgres) *MembersRepo {
	return &MembersRepo{
		db:     pg.DB,
		sq:     squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		getter: trmsqlx.DefaultCtxGetter,
	}
}

func (mr *MembersRepo) AddMember(ctx context.Context, chatId int, member models.Member) error {
	op := "MembersRepo.AddMember"

	sql, args, err := mr.sq.
		Insert("members").
		Columns("chat_id", "user_id", "role_id").
		Values(chatId, member.UserId, member.RoleId).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s: build query: %w", op, err)
	}

	_, err = mr.getter.DefaultTrOrDB(ctx, mr.db).ExecContext(ctx, sql, args...)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code {
			case "23505":
				return models.ErrUserAlreadyInChat
			case "23503":
				return models.ErrChatNotFound
			}
		}
		return fmt.Errorf("%s: ExecContext: %w", op, err)
	}

	return nil
}

func (mr *MembersRepo) ListMembers(ctx context.Context, chatId int) ([]models.Member, error) {
	op := "MembersRepo.ListMembers"

	sql, args, err := mr.sq.
		Select("user_id", "role_id").
		From("members").
		Where(squirrel.Eq{"chat_id": chatId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: build query: %w", op, err)
	}

	var members []models.Member
	err = mr.getter.DefaultTrOrDB(ctx, mr.db).SelectContext(ctx, &members, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: SelectContext: %w", op, err)
	}

	return members, nil
}

func (mr *MembersRepo) DeleteMember(ctx context.Context, chatId int, userId string) error {
	op := "MembersRepo.DeleteMember"

	sql, args, err := mr.sq.
		Delete("members").
		Where(squirrel.And{
			squirrel.Eq{"chat_id": chatId},
			squirrel.Eq{"user_id": userId},
		}).ToSql()
	if err != nil {
		return fmt.Errorf("%s: build query: %w", op, err)
	}

	cmd, err := mr.getter.DefaultTrOrDB(ctx, mr.db).ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("%s: ExecContext: %w", op, err)
	}

	affected, err := cmd.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: RowsAffected: %w", op, err)
	}

	if affected == 0 {
		return models.ErrMemberNotFound
	}

	return nil
}

func (mr *MembersRepo) SetRoleForMember(ctx context.Context, chatId int, userId string, roleId int) error {
	op := "MembersRepo.SetRoleForMember"

	sql, args, err := mr.sq.
		Update("members").
		Set("role_id", roleId).
		Where(squirrel.And{
			squirrel.Eq{"chat_id": chatId},
			squirrel.Eq{"user_id": userId},
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s: build query: %w", op, err)
	}

	cmd, err := mr.getter.DefaultTrOrDB(ctx, mr.db).ExecContext(ctx, sql, args...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23503":
				return models.ErrRoleNotFound
			}
		}
		return fmt.Errorf("%s: ExecContext: %w", op, err)
	}

	affected, err := cmd.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: RowsAffected: %w", op, err)
	}

	if affected == 0 {
		return models.ErrMemberNotFound
	}

	return nil
}

func (mr *MembersRepo) DeleteMembersForChat(ctx context.Context, chatId int) error {
	op := "MembersRepo.DeleteMembersForChat"

	sql, args, err := mr.sq.
		Delete("members").
		Where(squirrel.Eq{"chat_id": chatId}).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s: building query: %w", op, err)
	}

	_, err = mr.getter.DefaultTrOrDB(ctx, mr.db).ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("%s: ExecContext: %w", op, err)
	}

	return nil
}

func (mr *MembersRepo) UnsetRole(ctx context.Context, chatId int, oldRoleId, newRoleId int) error {
	op := "MembersRepo.UnsetRole"

	sql, args, err := mr.sq.
		Update("members").
		Set("role_id", newRoleId).
		Where(squirrel.And{
			squirrel.Eq{"chat_id": chatId},
			squirrel.Eq{"role_id": oldRoleId},
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s: build query: %w", op, err)
	}

	_, err = mr.getter.DefaultTrOrDB(ctx, mr.db).ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("%s: ExecContext: %w", op, err)
	}

	return nil
}

func (mr *MembersRepo) CheckMemberIsBanned(ctx context.Context, chatId int, userId string) (bool, error) {
	op := "MembersRepo.CheckUserIsBanned"

	query, args, err := mr.sq.
		Select("count(*)").
		From("banned_members").
		Where(squirrel.And{
			squirrel.Eq{"chat_id": chatId},
			squirrel.Eq{"user_id": userId},
		}).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("%s: build query: %w", op, err)
	}

	var count int
	err = mr.getter.DefaultTrOrDB(ctx, mr.db).QueryRowxContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("%s: QueryContext: %w", op, err)
	}

	return count != 0, nil
}

func (mr *MembersRepo) AddMemberToBanned(ctx context.Context, chatId int, userId string) error {
	op := "MembersRepo.AddMemberToBanned"

	sql, args, err := mr.sq.
		Insert("banned_members").
		Columns("chat_id", "user_id").
		Values(chatId, userId).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s: build query: %w", op, err)
	}

	_, err = mr.getter.DefaultTrOrDB(ctx, mr.db).ExecContext(ctx, sql, args...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23505":
				return models.ErrUserAlreadyBanned
			case "23503":
				return models.ErrChatNotFound
			}
		}
		return fmt.Errorf("%s: ExecContext: %w", op, err)
	}

	return nil
}

func (mr *MembersRepo) DeleteMemberFromBanned(ctx context.Context, chatId int, userId string) error {
	op := "MembersRepo.DeleteMemberFromBanned"

	sql, args, err := mr.sq.
		Delete("banned_members").
		Where(squirrel.And{
			squirrel.Eq{"chat_id": chatId},
			squirrel.Eq{"user_id": userId},
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s: build query: %w", op, err)
	}

	cmd, err := mr.getter.DefaultTrOrDB(ctx, mr.db).ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("%s: ExecContext: %w", op, err)
	}

	affected, err := cmd.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: RowsAffected: %w", op, err)
	}

	if affected == 0 {
		return models.ErrMemberNotFound
	}

	return nil
}

func (mr *MembersRepo) ListBannedMembers(ctx context.Context, chatId int) ([]string, error) {
	op := "MembersRepo.ListbannedMembers"

	sql, args, err := mr.sq.
		Select("user_id").
		From("banned_members").
		Where(squirrel.Eq{"chat_id": chatId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: build query: %w", op, err)
	}

	var bannedmembers []string
	err = mr.getter.DefaultTrOrDB(ctx, mr.db).SelectContext(ctx, &bannedmembers, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: SelectContext: %w", op, err)
	}

	return bannedmembers, nil
}
