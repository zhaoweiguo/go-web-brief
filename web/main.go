package main

import (
	"flag"
	"fmt"
	"net/http"


	"github.com/zhaoweiguo/go-web-brief/config"
	"github.com/zhaoweiguo/go-web-brief/db"
	"github.com/zhaoweiguo/go-web-brief/error"

	"github.com/golang/glog"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
)

type MainItem struct {
	Id int
	Title string
	Pic string
	Url string
}
type MainData struct {
	Date string
	MainPage []MainItem
}

type PageData struct {
}

type FinalData struct {
	MainData []MainData
	PageData []PageData
}

func main() {
	fmt.Println("go-web-brief started...")
	flag.Parse()
	glog.Infof("web ver: [%s] start...", config.Version)
	defer glog.Flush()
	
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "content", []interface{}{renderPages(1)})
	})
	// a test demo with http res and http req
	m.Get("/test.html", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
	})
	// a 404 page
	m.Get("/404.html", func() (int, string) {
		return 404, "you are in the wrong page"
	})
	// just a hello world
	m.Get("/hello.html", func() string {
		return "hello world"
	})
	http.ListenAndServe("0.0.0.0:7111", m)    // 修改监听端口
	m.Run()



}

func renderPages(page int) FinalData {

	var finalData FinalData
	var mainItemList []MainItem

	var mainDataList []MainData
	var pageDataList []PageData

	mainItemList = getMainData(page)

	fmt.Println(mainItemList)

	finalData.MainData = mainDataList
	finalData.PageData = pageDataList
	return finalData
}


func getMainData(page int) []MainItem {
	rows := db.Query(config.TAB_MAIN)
	var id int
	var title string
	var pic string
	var url string

	fmt.Println(rows)

	mainData := make([]MainItem, 10)

	i := 1;
	for rows.Next() {
		fmt.Println(i)
		i++
		err := rows.Scan(&id, &title, &pic, &url)
		error.CheckErr(err)
		mainData[i] = MainItem{id, title, pic, url}
	}

	fmt.Println(mainData)

	return []MainItem{}
}
