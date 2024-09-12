package storage

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetUsersSolAccounts(ctx context.Context) ([]User, error) {
	query, params, err := sq.Select("tg_id", "tg_username", "sol_public_key", "tg_chat_id").
		From(usersTableName).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var dest []User

	err = s.db.SelectContext(ctx, &dest, s.db.Rebind(query), params...)
	if err != nil {
		return nil, err
	}

	return dest, nil
}
