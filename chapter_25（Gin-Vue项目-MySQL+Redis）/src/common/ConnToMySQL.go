package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"util"
)

func ConnectToMySQL() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/blog")
	util.HandleError(err, "sqlx.Open")
	return db
}
