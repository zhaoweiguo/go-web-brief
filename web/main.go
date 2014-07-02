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


func main() {
	fmt.Println("go-web-brief started...")
	flag.Parse()
	glog.Infof("web ver: [%s] start...", config.Version)
	defer glog.Flush()
	
	m := martini.Classic()
//	fmt.Println(config.PATH_BASE + "templates")
	m.Use(martini.Static(config.PATH_BASEWEB + "static"))
	m.Use(render.Renderer(render.Options{
		Directory : config.PATH_BASEWEB + "templates",
		Extensions: []string{".tmpl", ".html"},
		Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".

	}))



	m.Get("/", func(r render.Render) {
		r.HTML(200, "content", []interface{}{renderPages(1)})
	})

	m.Get("/testtpl.html", func(r render.Render) {
		r.HTML(200, "testtpl", nil)
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
	http.ListenAndServe("0.0.0.0:" + config.PORT, m)    // 修改监听端口
	m.Run()

}

