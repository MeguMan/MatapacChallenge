package storage

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) CreateUser(ctx context.Context, user User) error {
	query, params, err := sq.Insert(usersTableName).
		Columns(
			"tg_id",
			"tg_username",
			"sol_public_key",
			"tg_chat_id",
		).
		Values(user.TgID, user.TgUsername, user.SolPublicKey, user.TgChatID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, s.db.Rebind(query), params...)
	return handleErr(err)
}
