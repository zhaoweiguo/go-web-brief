package sqlite3

import(
	"os"
	"fmt"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/zhaoweiguo/go-web-brief/config"

	"testing"
	"github.com/bmizerany/assert"
)

func TestCreateTable(t *testing.T) {
	os.Remove(config.DB_UNITTEST_FILE)
	db, err := sql.Open("sqlite3", config.DB_UNITTEST_FILE)
	if err != nil {
		t.Error("can't open the database file")
	}
	defer db.Close()

	sql := "create table gordon (id int primary key, value text);"
	_, err = db.Exec(sql)
	if err != nil {
		t.Error("can't create table")
	}
}

func TestInsertTable(t *testing.T) {
	sqls := make([]string, 10)
	for i:=0; i<10; i++ {
		sqls[i] = fmt.Sprintf("insert into gordon (id, value) values(%d, 'aaaa%d');", i, i)
	}
//	t.Log(sqls)
	db, _ := sql.Open("sqlite3", config.DB_UNITTEST_FILE)
	defer db.Close()
	for _,s := range sqls {
		rtn, _ := db.Exec(s)
		t.Log(rtn.LastInsertId())
	}

}

func TestSelectTable(t *testing.T) {
	db, _ := sql.Open("sqlite3", config.DB_UNITTEST_FILE)
	defer db.Close()
	sql := "select id, value from gordon;"
	rows, _ := db.Query(sql)
	defer rows.Close()
	i :=0
	for rows.Next() {
		var id int
		var value string
		rows.Scan(&id, &value)
		assert.Equal(t, i, id)
		assert.Equal(t, fmt.Sprintf("aaaa%d", i), value)
		i++
	}

}

func TestDeleteTable(t *testing.T) {
	db, _ := sql.Open("sqlite3", config.DB_UNITTEST_FILE)
	defer db.Close()
	sql := "delete from gordon;"
	rtn, err := db.Exec(sql)
	if err != nil {
		t.Error(err)
	}
	t.Log(rtn.RowsAffected())
}

func TestSelectTable2(t *testing.T) {
	db, _ :=sql.Open("sqlite3", config.DB_UNITTEST_FILE)
	defer db.Close()
	rows, _ := db.Query("select id ,value from gordon")
	defer rows.Close()

	assert.Equal(t, rows.Next(), false)

}

func TestDropTable(t *testing.T) {
	db, _ := sql.Open("sqlite3", config.DB_UNITTEST_FILE)
	defer db.Close()
	Result, err := db.Exec("update gordon set value='123'")
	if err != nil {
		t.Error(err)
	}
}
