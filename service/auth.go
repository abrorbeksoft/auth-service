package service

import (
	"context"
	"database/sql"
	"github.com/abrorbeksoft/auth-service/genproto/auth"
	"github.com/abrorbeksoft/auth-service/storage"
)

type authService struct {
	storage storage.StorageI
}

func NewAuthService(db *sql.DB) *authService {
	return &authService{

	}
}

func (auth *authService) Create(ctx context.Context, user *auth.User) (*auth.UserResponce, error) {
	panic("implement me")
}

func (auth *authService) Find(ctx context.Context, request *auth.GetRequest) (*auth.UserResponce, error) {
	panic("implement me")
}
