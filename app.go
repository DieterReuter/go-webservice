package main

import (
	"fmt"
)

var API_PORT int
var CONFIG_FILE string
var DATASTORE_DIR string

func main() {
	fmt.Println("API_PORT=", API_PORT)
	fmt.Println("CONFIG_FILE=", CONFIG_FILE)
	fmt.Println("DATASTORE_DIR=", DATASTORE_DIR)
}
