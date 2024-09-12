package storage

type User struct {
	TgID         int    `db:"tg_id"`
	TgUsername   string `db:"tg_username"`
	SolPublicKey string `db:"sol_public_key"`
}
