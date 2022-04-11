package model

type User struct {
	ID       uint   `db:"id"`
	UserID   uint64 `db:"user_id"`
	UserName string `db:"username"`
	Password string `db:"password"`
	Phone    string `db:"phone"`
}
