package application

import (
	"context"
	"errors"
	"os"

	"github.com/google/uuid"
	"github.com/nazrawigedion123/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrSaltNotSet = errors.New("SALT environment variable not set")
)

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	user.ID = uuid.New()
	salt := os.Getenv("SALT")
	if salt == "" {
		return ErrSaltNotSet
	}

	hasedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password+salt), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hasedpassword)

	return s.userRepository.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.userRepository.GetUserByID(ctx, id)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.userRepository.GetUserByEmail(ctx, email)
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	return s.userRepository.GetUserByUsername(ctx, username)
}

func (s *UserService) ListUsers(ctx context.Context) ([]*domain.User, error) {
	return s.userRepository.ListUsers(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) error {
	return s.userRepository.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.userRepository.DeleteUser(ctx, id)
}
