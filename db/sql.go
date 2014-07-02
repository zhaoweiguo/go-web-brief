package db

import (
	"fmt"
	"database/sql"
	"github.com/zhaoweiguo/go-web-brief/error"
	"github.com/zhaoweiguo/go-web-brief/config"
)

func Query(table string) *sql.Rows {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s%s", config.PATH_DB, config.DB_FILE))
	error.CheckErr(err)
	fmt.Printf("%smain.db", config.PATH_DB)
	rows, err := db.Query("SELECT * FROM " + table)
	error.CheckErr(err)
	db.Close()

	return rows

}
