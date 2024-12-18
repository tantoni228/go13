package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go13/pkg/postgres"
	"go13/user-service/internal/dto"
	"go13/user-service/internal/models"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UsersRepo struct {
	db *sqlx.DB
	sq squirrel.StatementBuilderType
}

func NewUsersRepo(pg *postgres.Postgres) *UsersRepo {
	return &UsersRepo{
		db: pg.DB,
		sq: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (ur *UsersRepo) AddUser(ctx context.Context, user models.User) (models.User, error) {
	op := "UsersRepo.AddUser"

	sql, args, err := ur.sq.
		Insert("users").
		Columns("username", "email", "hashed_password", "bio").
		Values(user.Username, user.Email, user.HashedPassword, user.Bio).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("%s: build query: %w", op, err)
	}

	var createduser models.User
	err = ur.db.GetContext(ctx, &createduser, sql, args...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23505":
				return models.User{}, models.ErrEmailIsTaken
			}
		}
		return models.User{}, fmt.Errorf("%s: GetContext: %w", op, err)
	}

	return createduser, nil
}

func (ur *UsersRepo) GetUserById(ctx context.Context, userId string) (models.User, error) {
	op := "UsersRepo.GetUserById"

	query, args, err := ur.sq.
		Select("*").
		From("users").
		Where(sq.Eq{"id": userId}).
		ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("%s: build query: %w", op, err)
	}

	var user models.User
	err = ur.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, models.ErrUserNotFound
		}
	}

	return user, nil
}

func (ur *UsersRepo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	op := "UsersRepo.GetUserById"

	query, args, err := ur.sq.
		Select("*").
		From("users").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("%s: build query: %w", op, err)
	}

	var user models.User
	err = ur.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, models.ErrUserNotFound
		}
	}

	return user, nil
}

func (ur *UsersRepo) UpdateUser(ctx context.Context, userId string, input dto.UpdateUserInput) (models.User, error) {
	op := "UserRepo.UpdateUser"

	builder := ur.sq.Update("users")
	if input.Username.IsSet() {
		builder = builder.Set("username", input.Username.GetValue())
	}
	if input.Email.IsSet() {
		builder = builder.Set("email", input.Email.GetValue())
	}
	if input.HashedPassword.IsSet() {
		builder = builder.Set("hashed_password", input.HashedPassword.GetValue())
	}
	if input.Bio.IsSet() {
		builder = builder.Set("bio", input.Bio.GetValue())
	}
	query, args, err := builder.
		Where(sq.Eq{"id": userId}).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("%s: build query: %w", op, err)
	}

	var user models.User
	err = ur.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code {
			case "23505":
				return models.User{}, models.ErrEmailIsTaken
			}
		}
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, models.ErrUserNotFound
		}
	}

	return user, nil
}
