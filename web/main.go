package main

import (
	"flag"
	"fmt"
	"http"
	"github.com/zhaoweiguo/go-web-brief/config"
	"github.com/golang/glog"
	"github.com:martini-contrib/render"
	"github.com/go-martini/martini"
)


func main() {
	fmt.Println("go-web-brief started...")
	flag.Parse()
	glog.Infof("web ver: [%s] start...", config.Version)
	defer glog.Flush()
	
	m := martini.Classic()
	m.Get("/", func(r render.Render) string {
		return "hello world"
	})
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
	m.Run()



}


