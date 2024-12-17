package postgres

import (
	"context"
	"database/sql"
	"fmt"

	api "go13/pkg/ogen/users-service"
	"go13/pkg/postgres"
	"go13/user-service/internal/models"
	"go13/user-service/internal/transport/rest/auth"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type UserRepo struct {
	db *postgres.Postgres
}

func NewUsersRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{db: pg}
}

func GenerateRandomUUID() string {
	id := uuid.New()
	return id.String()
}

func (ur *UserRepo) SignUp(ctx context.Context, req *api.SignUpReq) (api.SignUpRes, error) {
	_, err := sq.Insert("users").
		Columns("id", "user_name", "user_email", "user_password").
		Values(uuid.New(), req.Username, req.Email, req.Password).
		Suffix("returning *").
		PlaceholderFormat(sq.Dollar).
		RunWith(ur.db.DB.DB).
		Exec()
	if err != nil {
		return nil, fmt.Errorf("repository.SendMessage: %w", err)
	}
	return &api.SignUpNoContent{}, nil
}

func (ur *UserRepo) SignIn(ctx context.Context, req *api.SignInReq) (api.SignInRes, error) {
	var res api.SignInRes

	// Prepare the select query to find the user
	sqlQuery, args, err := sq.Select("user_id", "user_name").
		From("users").
		Where(sq.Eq{"user_email": req.Email, "user_password": req.Password}).
		ToSql()
	if err != nil {
		return res, fmt.Errorf("failed to build select query: %w", err)
	}

	var UserId, Username int
	err = ur.db.DB.QueryRowContext(ctx, sqlQuery, args...).Scan(&UserId, &Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, fmt.Errorf("invalid credentials")
		}
		return res, fmt.Errorf("failed to execute select query: %w", err)
	}

	return res, nil
}

func (ur *UserRepo) ChangePassword(ctx context.Context, req *api.ChangePasswordReq) (api.ChangePasswordRes, error) {
	var res api.ChangePasswordRes

	updateQuery, args, err := sq.Update("users").
		Set("user_password", req.NewPassword).
		Where(sq.Eq{"user_id": auth.UserIdFromCtx(ctx)}).
		ToSql()
	if err != nil {
		return res, fmt.Errorf("failed to build update query: %w", err)
	}

	// Execute the update query
	_, err = ur.db.DB.ExecContext(ctx, updateQuery, args...)
	if err != nil {
		return res, fmt.Errorf("failed to execute update query: %w", err)
	}

	return res, nil
}

// func (ur *UserRepo) ChangePassword(ctx context.Context, req *api.ChangePasswordReq, UserId string) (models.Password, error) {

// 	updateQuery, args, err := sq.Update("users").
// 		Set("user_password", req.NewPassword).
// 		Where(sq.Eq{"user_id": UserId}).
// 		ToSql()
// 	if err != nil {
// 		return models.Password(req.OldPassword), fmt.Errorf("failed to build update query: %w", err)
// 	}

// 	// Execute the update query
// 	_, err = ur.db.DB.ExecContext(ctx, updateQuery, args...)
// 	if err != nil {
// 		return models.Password(req.OldPassword), fmt.Errorf("failed to execute update query: %w", err)
// 	}

// 	return models.Password(req.NewPassword), nil
// }

func (ur *UserRepo) UpdateUser(ctx context.Context, userId string, user models.User) error {
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

// func (ur *UserRepo) GetUserById(ctx context.Context, userId string) (models.User, error) {
// 	var user models.User

// 	sqlQuery, args, err := sq.Select("user_id", "user_name", "user_email", "user_password", "user_bio").
// 		From("users").
// 		Where(sq.Eq{"user_id": userId}).
// 		ToSql()
// 	if err != nil {
// 		return user, fmt.Errorf("failed to build query: %w", err)
// 	}

// 	err = ur.db.DB.QueryRowContext(ctx, sqlQuery, args...).Scan(&user.UserId, &user.Username, &user.Email, &user.Password, &user.Bio)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return user, models.ErrIDNotFound
// 		}
// 		return user, fmt.Errorf("failed to execute query: %w", err)
// 	}

// 	return user, nil
// }

func (ur *UserRepo) CheckUserByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool

	existsQuery, _, err := sq.Select("EXISTS(SELECT 1 FROM users WHERE email = ?)").
		ToSql()
	if err != nil {
		return false, fmt.Errorf("failed to build query: %w", err)
	}

	err = ur.db.DB.QueryRowContext(ctx, existsQuery, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to execute query: %w", err)
	}

	return exists, nil
}
