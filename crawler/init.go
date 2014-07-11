package crawler

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/zhaoweiguo/go-web-brief/config"
	"github.com/zhaoweiguo/go-web-brief/db"
	"github.com/zhaoweiguo/go-web-brief/error"
)

type ZhihuItem struct {
	Title string
	Url string
	Image string
	Share_url string
	Thumbnail string
	Ga_prefix string
	Id int
}

type ZhihuData struct {
	Date string
	News []ZhihuItem
}

const(
	urlnow = "http://news.at.zhihu.com/api/1.2/news/latest"
	urlbefore = "http://news.at.zhihu.com/api/1.2/news/before/"
)

func getTodayData() ZhihuData{
	if resp, err := http.Get(urlnow); err == nil {
		defer resp.Body.Close()
		var jsonBodys ZhihuData
		if body, err := ioutil.ReadAll(resp.Body); err == nil {

			if err = json.Unmarshal(body, &jsonBodys); err == nil {
				return jsonBodys
			} else {
				error.CheckErr(err)
			}
		} else {
			error.CheckErr(err)
		}
	}
	fmt.Println("crawler init/0 no data found")
	return ZhihuData{}
}

func CrawlerData() {
	data := getTodayData()

	date := data.Date
	items := data.News
	sqls := make([]string, 10)

	for _, item := range items {
		id := item.Id
		title := item.Title
		image := item.Image
		url := item.Url
		share_url := item.Share_url
		thumbnail := item.Thumbnail

		sql := fmt.Sprintf("INSERT INTO %s (id, title, image, url ,share_url, thumbnail, date) values(%d, '%s', '%s', '%s', '%s', '%s', '%s')",  config.TAB_MAIN, id, title, image, url, share_url, thumbnail, date)
		sqls = append(sqls, sql)
	}
//	fmt.Println(sqls)
	db.Execs(sqls)
}


