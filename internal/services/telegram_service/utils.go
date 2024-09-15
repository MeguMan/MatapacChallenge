package telegram_service

import (
	"context"
	"fmt"
	"sort"
	"time"
)

func (s *service) getTopMsg() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.topMsg
}

func (s *service) setTopMsg(text string) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	s.topMsg = text
}

func (s *service) syncUsersBalanceTicker(ctx context.Context) error {
	textMsg, err := s.calculateUsersBalance(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	s.setTopMsg(textMsg)

	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				textMsg, err = s.calculateUsersBalance(ctx)
				if err != nil {
					fmt.Println(err)
					continue
				}

				s.setTopMsg(textMsg)
			}
		}
	}()

	return nil
}

func (s *service) calculateUsersBalance(ctx context.Context) (string, error) {
	users, err := s.storageService.GetUsersSolAccounts(ctx)
	if err != nil {
		return "", err
	}

	mpUserNameByPublicKey := make(map[string]string, len(users))
	publicKeys := make([]string, 0, len(users))
	for _, user := range users {
		mpUserNameByPublicKey[user.SolPublicKey] = user.TgUsername
		publicKeys = append(publicKeys, user.SolPublicKey)
	}

	accounts, err := s.chainstackService.GetAccountsBalance(ctx, publicKeys)
	if err != nil {
		return "", err
	}

	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].Sol > accounts[j].Sol
	})

	textMsg := ""
	for i, account := range accounts {
		textMsg += fmt.Sprintf("%d. <a href='https://solscan.io/account/%s'>%s</a> - %f\n", i+1, account.PublicKey, mpUserNameByPublicKey[account.PublicKey], account.Sol)
	}

	return textMsg, nil
}
