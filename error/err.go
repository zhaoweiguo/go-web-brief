package error

import (
	"fmt"
	"bytes"
)


func CheckErr(args ...interface{} ) {

	if args[0] == nil {
		return
	}
	var format bytes.Buffer
	for _ = range args {
		format.WriteString("%v  ")
	}
	fmt.Printf(format.String()+"  \n", args...)
}



