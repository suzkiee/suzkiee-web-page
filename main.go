package main

import (
	"fmt"
	"net/http"
)

const portNum = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf(fmt.Sprintf("Starting	application on port# %s", portNum))

	_ = http.ListenAndServe(portNum, nil)
}
