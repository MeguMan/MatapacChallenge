package chainstask_service

import "context"

type Service interface {
	GetAccountsBalance(ctx context.Context, publicKeys []string) ([]SolAccount, error)
}

type service struct {
	chainstackURL string
}

func New(chainstackURL string) Service {
	return &service{
		chainstackURL: chainstackURL,
	}
}
