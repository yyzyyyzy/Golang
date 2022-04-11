package model

import "database/sql"

type Post struct {
	ID         int  `json:"id" ,db:"id"`
	PostID     uint `json:"post_id" ,db:"post_id"`
	UserID     uint `json:"user_id" ,db:"user_id"`
	CategoryID uint `json:"category_id" ,db:"category_id"`
	Category   *Category
	Title      string `json:"title" ,db:"title"`
	HeadImg    string `json:"head_img" ,db:"head_img"`
	Content    string `json:"content" ,db:"content"`
	CreateAt   Time   `json:"create_time" ,db:"create_time"`
	UpdateAt   Time   `json:"update_time" ,db:"update_time" `
}

type Postman struct {
	ID sql.NullInt64 `json:"id," ,db:"id"`
}
