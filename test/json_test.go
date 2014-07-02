package main

import(
	"fmt"
	"encoding/json"

	"testing"
	"github.com/bmizerany/assert"

)

func TestUnmarshal(t *testing.T) {
	var jsonBlob = []byte(`[
        {"Name": "number 0","Order": "order 0"},
        {"Name": "number 1","Order": "order 1"}
    ]`)
	type Animal struct {
		Name  string
		Order string
	}

	jsonRtn := make([]Animal, 2)

	jsonRtn[0] = Animal{ "number 0", "order 0"}
	jsonRtn[1] = Animal{"number 1", "order 1"}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)

	fmt.Println(err)
	assert.Equal(t, animals, jsonRtn)
}


