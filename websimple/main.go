package main

import(
	"net/http"
	"log"
	"fmt"

	"github.com/zhaoweiguo/go-web-brief/config"	
)

func justHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("form:", r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println("param1:", r.Form["param1"])

	for k, v := range r.Form {
		fmt.Println("k:", k)
		fmt.Println("v:", v)
	}
	fmt.Fprintf(w, "just say hello")
}

func main() {
	http.HandleFunc("/", justHello)
	err := http.ListenAndServe("0.0.0.0:" + config.PORT, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
