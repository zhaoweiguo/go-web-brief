package db

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	PATH_DB = "/tmp/go-web-brief/db/"
	PATH_FILES = "/tmp/go-web-brief/files/"
)

func init() {
	fmt.Println("init.go init func...")
	filepath := PATH_FILES + "dbinit.flag"

	initFlag, err := ioutil.ReadFile(filepath)
	if err != nil {  // 如还没有进行数据初使化
		InitDB()
		err = ioutil.WriteFile(filepath, [byte(123)], 0644)
	}
}

func InitDB() {
	path := "./go-web-brief.db"
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()

	sql := `create table if not exists goWebBriefTest (
               id integer primary key,
               name text
           )`
	fmt.Println(sql)
	_, err = db.Exec(sql)
	checkErr(err, sql)

	tx, err := db.Begin()

	stmt, err := tx.Prepare("insert into goWebBriefTest values(?, ?)")
	checkErr(err)
	defer stmt.Close()
	for i := 0; i<10; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("新溪-gordon %03d", i))
		checkErr(err)
	}
	tx.Commit()

	rows, err := db.Query("select * from goWebBriefTest")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
	rows.Close()

/*
	_, err = db.Exec("delete from goWebBriefTest")
	checkErr(err)
*/
}


func checkErr(args ...interface{} ) {

	if args[0] == nil {
		return
	}
	var format bytes.Buffer
	for _ = range args {
		format.WriteString("%v  ")
	}
	fmt.Printf(format.String()+"  \n", args...)
}


