package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/nazrawigedion123/user-service/internal/adapters/database"
	"github.com/nazrawigedion123/user-service/internal/domain"
)

type PostgreRepository struct {
	db *database.Queries
}

func NewPostgreRepository(db *database.Queries) *PostgreRepository {
	return &PostgreRepository{
		db: db,
	}

}

func (r *PostgreRepository) CreateUser(ctx context.Context, user *domain.User) error {
	arg := database.CreateUserParams{
		ID:        user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     sql.NullString{String: user.Phone, Valid: true},
	}

	_, err := r.db.CreateUser(ctx, arg)
	if err != nil {
		return err
	}

	return nil

}
func (r *PostgreRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := r.db.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone.String,
	}, nil
}

func (r *PostgreRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := r.db.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone.String,
	}, nil
}

func (r *PostgreRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := r.db.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone.String,
	}, nil

}

// func (r *PostgreRepository) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
// 	user, err := r.db.GetUserByPhone(ctx, phone)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &domain.User{
// 		ID:        user.ID,
// 		UserName:  user.UserName,
// 		FirstName: user.FirstName,
// 		LastName:  user.LastName,
// 		Email:     user.Email,
// 		Phone:     user.Phone.String,
// 	}, nil
// }

func (r *PostgreRepository) ListUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := r.db.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []*domain.User
	for _, user := range users {
		result = append(result, &domain.User{
			ID:        user.ID,
			UserName:  user.UserName,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone.String,
		})
	}

	return result, nil
}

func (r *PostgreRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	arg := database.UpdateUserParams{
		ID:        user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     sql.NullString{String: user.Phone, Valid: true},
	}

	_, err := r.db.UpdateUser(ctx, arg)
	if err != nil {
		return err
	}

	return nil

}

func (r *PostgreRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := r.db.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
