package model

type Category struct {
	ID       uint   `db:"id" ,json:"id"`
	Name     string `db:"name" ,json:"name" `
	CreateAt Time   `db:"create_time" ,json:"create_time" `
	UpdateAt Time   `db:"update_time" ,json:"update_time" `
}
