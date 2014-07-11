package test

import(
	"testing"
	"fmt"
)

func TestPrintf(t *testing.T) {

	var abc string
	abc = "aaa"
	t.Log(fmt.Sprintf("fmt:%s, %s\n", abc, "abcc"))
}


