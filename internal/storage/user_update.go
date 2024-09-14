package storage

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) UpdateUser(ctx context.Context, user User) error {
	query, params, err := sq.Update(usersTableName).
		SetMap(map[string]interface{}{
			"sol_public_key": user.SolPublicKey,
			"attempt":        "attempt + 1",
		}).
		Where(sq.Eq{"tg_id": user.TgID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, s.db.Rebind(query), params...)
	return handleErr(err)
}
