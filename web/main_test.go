package main

import(
	"fmt"
	"testing"
)


func TestRenderPages(t *testing.T) {
	rtn := renderPages(1)
	fmt.Println(rtn)
}
