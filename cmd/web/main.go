package main

import (
	"fmt"
	"net/http"

	"github.com/suzkiee/web-page/pkg/handlers"
)

const portNum = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf(fmt.Sprintf("Starting	application on port# %s", portNum))

	_ = http.ListenAndServe(portNum, nil)
}
