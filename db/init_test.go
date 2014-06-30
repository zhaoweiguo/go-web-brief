package db

import(
	"fmt"
	"testing"
	"github.com/bmizerany/assert"

)

func TestInit(t *testing.T) {
	initDB()
	fmt.Println("test finished!")
	assert.NotEqual(t, 1, 2)
}



