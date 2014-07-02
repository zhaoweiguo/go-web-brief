package db

import (
	"database/sql"
	"github.com/zhaoweiguo/go-web-brief/error"
	"github.com/zhaoweiguo/go-web-brief/config"
)

func Query(table string) *sql.Rows {
	db, err := sql.Open("sqlite3", config.DB_FILE)
	error.CheckErr(err)
	rows, err := db.Query("SELECT * FROM " + table)
	error.CheckErr(err)
	db.Close()

	return rows

}
