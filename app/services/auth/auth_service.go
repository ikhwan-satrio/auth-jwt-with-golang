package auth

import (
	"context"
	"fmt"

	"github.com/ikhwan-satrio/auth-golang/app/db"
	"github.com/ikhwan-satrio/auth-golang/app/services/tokens"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(ctx context.Context, name, email, password string) error
	Login(ctx context.Context, email, password string) (string, error)
}

type authService struct {
	db  *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{db: db}
}

func (s *authService) Register(ctx context.Context, name, email, password string) error {
	if email == "" || password == "" || name == "" {
		return fmt.Errorf("all fields required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	user := db.User{
		Name:     &name,
		Email:    email,
		Password: string(hash),
	}

	return gorm.G[db.User](s.db).Create(ctx, &user)
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	if email == "" || password == "" {
		return "", fmt.Errorf("all fields required")
	}

	user, err := gorm.G[db.User](s.db).Where("email = ?", email).First(ctx)
	if err != nil {
		return "", fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password")
	}

	return tokens.CreateToken(user.Email)
}
