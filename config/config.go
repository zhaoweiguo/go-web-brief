package config

const (
//	PATH_BASE "/var/"
	PATH_DB = "/tmp/go-web-brief/db/"
	PATH_FILES = "/tmp/go-web-brief/files/"
	PATH_BASE = "/Users/zhaoweiguo/test/go/src/github.com/zhaoweiguo/go-web-brief/"
	PATH_BASEWEB = "/Users/zhaoweiguo/test/go/src/github.com/zhaoweiguo/go-web-brief/web/"
	PATH_IMG = ""
)

const (
	TAB_MAIN = "mainpage"
	TAB_PREFIX = "ga_prefix"
)

const (
	DB_FILE = "file:" + PATH_DB + "go-web-brief.db" + "?cache=shared&mode=rwc"
	DB_UNITTEST_FILE = PATH_DB + "go-web-brief_unittest.db"
)


const (
//	PORT = "7111"
	PORT = "80"
)
