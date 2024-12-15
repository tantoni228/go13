package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"go13/pkg/postgres"
	"go13/user-service/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type UsersRepo struct {
	db *postgres.Postgres
}

func NewUsersRepo(pg *postgres.Postgres) *UsersRepo {
	return &UsersRepo{
		db: pg,
	}
}

func (ur *UsersRepo) UpdateUser(ctx context.Context, userId string, user models.User) error {
	updateQuery, args, err := sq.Update("users").
		Set("user_name", user.Username).
		Set("user_email", user.Email).
		Set("user_password", user.Password).
		Set("user_bio", user.Bio).
		Where(sq.Eq{"user_id": userId}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = ur.db.DB.ExecContext(ctx, updateQuery, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func (ur *UsersRepo) GetUserById(ctx context.Context, userId string) (models.User, error) {
	var user models.User

	sqlQuery, args, err := sq.Select("user_id", "user_name", "user_email", "user_password", "user_bio").
		From("users").
		Where(sq.Eq{"user_id": userId}).
		ToSql()
	if err != nil {
		return user, fmt.Errorf("failed to build query: %w", err)
	}

	err = ur.db.DB.QueryRowContext(ctx, sqlQuery, args...).Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.Bio)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, models.ErrIDNotFound
		}
		return user, fmt.Errorf("failed to execute query: %w", err)
	}

	return user, nil
}

func (ur *UsersRepo) CheckUser(ctx context.Context, userId string) (bool, error) {
	var exists bool

	existsQuery, args, err := sq.Select("exists(select 1 from users where user_id = ?)", userId).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("failed to build query: %w", err)
	}

	err = ur.db.DB.QueryRowContext(ctx, existsQuery, args...).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}

	return exists, nil
}
