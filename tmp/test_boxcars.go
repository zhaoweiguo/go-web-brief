package main

import (
	"fmt"
	"github.com/azer/boxcars"
	"github.com/azer/boxcars/json-config"
)

func main() {
	filename := "./example.json"
	config := JSONConfig.NewJSONConfig(filename, boxcars.SetupSites)
	fmt.Println(config)
}
