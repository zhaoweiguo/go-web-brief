package db

import (
	"fmt"
	"io/ioutil"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zhaoweiguo/go-web-brief/error"
	"github.com/zhaoweiguo/go-web-brief/config"
)

func init() {
	fmt.Println("init.go init func...")
	filepath := fmt.Sprintf("%sdbinit.flag", config.PATH_FILES)

	_, err := ioutil.ReadFile(filepath)
	if err != nil {  // 如还没有进行数据初使化
		InitDB()
		err = ioutil.WriteFile(filepath, []byte(""), 0644)
	} else {
		fmt.Println("你已经初使化过了...")
	}
	
}

func InitDB() {
	path := fmt.Sprintf("%sgo-web-brief.db", config.PATH_DB)
	db, err := sql.Open("sqlite3", path)
	error.CheckErr(err)
	defer db.Close()

	sql := fmt.Sprintf(`create table if not exists %s (
               id integer primary key,
               title char(200),
               pic char(200),
               url char(200)
            	)`, config.TAB_MAIN)

	_, err = db.Exec(sql)
	error.CheckErr(err, sql)
	initData(db)
}


func initData(db *sql.DB) {
	for i := 0; i<10; i++ {
		// fmt.Sprintf("新溪-gordon %d ", i)
		sql := fmt.Sprintf("insert into %s values (%d, '%s', '%s', '%s')", config.TAB_MAIN, i, "新溪gordon测试用例", "http://pic1.zhimg.com/89e91c96f515b2e96d5b65b9d873b208.jpg", "http://daily.zhihu.com/story/3996747")
		fmt.Println(sql)
		_, err := db.Exec(sql)
		error.CheckErr(err)
	}
	db.Close()

}

func test() {
	path := fmt.Sprintf("%sgo-web-brief.db", config.PATH_DB)
	db, err := sql.Open("sqlite3", path)
	rows, err := db.Query(fmt.Sprintf("select * from %s", config.TAB_MAIN))
	error.CheckErr(err)
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
	CheckErr(err)
*/

}

