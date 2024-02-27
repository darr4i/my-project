package main

import (
	"github.com/darr4i/my-project/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/time", controllers.HandleTimeRequest)
	http.ListenAndServe(":8795", nil)
}
