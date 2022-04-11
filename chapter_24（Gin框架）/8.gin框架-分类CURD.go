package main

import "time"

type Category struct {
	ID       uint      `json:"id,omitempty" ,db:"id"`
	Name     string    `json:"name,omitempty" ,db:"name"`
	CreateAt time.Time `json:"create_at" ,db:"create_time"`
	UpdateAt time.Time `json:"update_at" ,db:"update_time"`
}
