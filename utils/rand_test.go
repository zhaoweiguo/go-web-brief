package utils

import(
	"math/rand"

	"testing"
	"github.com/bmizerany/assert"
)

func TestRand(t *testing.T) {
	assert.Equal(t, 1, rand.Intn(10))  //每次结果一样


}
