package main

import (
	"flag"
	"fmt"
	"net/http"
	"github.com/zhaoweiguo/go-web-brief/config"
	"github.com/golang/glog"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
)

type MainItem struct {
	Id int
	Title string
	ShareImage string
}
type MainData struct {
	Date string
	MainPage []MainItem
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
	var mainDataList []MainData
	var pageDataList []PageData

	mainDataList = getMainData(page)

	finalData.MainData = mainDataList
	finalData.PageData = pageDataList
	return finalData
}


func getMainData(page int) []MainData {
	
}
