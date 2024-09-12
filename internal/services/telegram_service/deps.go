package telegram_service

import (
	"context"
	"github.com/MeguMan/MatapacChallenge/internal/services/chainstask_service"
	"github.com/MeguMan/MatapacChallenge/internal/storage"
)

type storageService interface {
	GetUsersSolAccounts(ctx context.Context) ([]storage.User, error)
	CreateUser(ctx context.Context, user storage.User) error
}

type chainstackService interface {
	GetAccountsBalance(ctx context.Context, publicKeys []string) ([]chainstask_service.SolAccount, error)
}
