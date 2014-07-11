package crawler

import (
	"fmt"
	"time"
	"github.com/zhaoweiguo/go-web-brief/db"
	"github.com/zhaoweiguo/go-web-brief/error"
	"github.com/zhaoweiguo/go-web-brief/config"
)

func init() {
	fmt.Println("crawler init/0")
	updateData()
//	ticker := time.NewTicker(time.Hour*2)   // 2小时更新一次
	ticker := time.NewTicker(time.Minute)  //测试
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
			updateData()
		}
	}()
}

func updateData() {
	fmt.Println("crawler updateData()")
	rows := db.Query(config.TAB_PREFIX)

	rows.Next()

	var ga_prefix string
	var id int

	err := rows.Scan(&id, &ga_prefix)
	error.CheckErr(err)

	fmt.Println("-----------")
	fmt.Println(ga_prefix)
	fmt.Println(id)
	CrawlerDataAfter(ga_prefix)
}

func CrawlerDataAfter(ga_prefix string) {
	fmt.Println("ticker CrawlerDataAfter/0)")
	data := getTodayData()
//	fmt.Println(data)
	date := data.Date
	items := data.News
	sqls := make([]string, 5, 100)

	max_prefix := ""

	fmt.Printf("ga_prefix:%s\n", ga_prefix)
	for _, item := range items {
		fmt.Println(item.Ga_prefix)
		if(item.Ga_prefix < ga_prefix) {
			continue
		}
		id := item.Id
		title := item.Title
		image := item.Image
		url := item.Url
		share_url := item.Share_url
		thumbnail := item.Thumbnail

		s := fmt.Sprintf("INSERT INTO %s (id, title, image, url ,share_url, thumbnail, date) values(%d, '%s', '%s', '%s', '%s', '%s', '%s')",  config.TAB_MAIN, id, title, image, url, share_url, thumbnail, date)
//		fmt.Println(s)
		sqls = append(sqls, s)

		if item.Ga_prefix > max_prefix {
			max_prefix = item.Ga_prefix
		}
	}
	db.Execs(sqls)
	if(max_prefix != "") {
		// @todo
		fmt.Println("update the prefix")
	}
}


