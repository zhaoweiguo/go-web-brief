package db

import (
	"fmt"
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

func QueryByDate(table, date string) *sql.Rows {
	db, err := sql.Open("sqlite3", config.DB_FILE)
	error.CheckErr(err)
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT id, title, image, share_url FROM %s where date = %s", table, date))
	error.CheckErr(err)

	return rows
}

func Execs(sqls []string) {
	db, err := sql.Open("sqlite3", config.DB_FILE)
	defer db.Close()
	error.CheckErr(err)
	for _, sql := range sqls {
//		rtn, err := db.Exec(sql)
		_, err = db.Exec(sql)
		error.CheckErr(err)
//		fmt.Println(sql)
	}
	db.Close()
}
func Exec(s string) {
//	fmt.Println(s)
	db, err := sql.Open("sqlite3", config.DB_FILE)
	defer db.Close()
	error.CheckErr(err)
	_, err = db.Exec(s)
	error.CheckErr(err)
	db.Close()
}
