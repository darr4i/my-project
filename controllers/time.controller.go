package controllers

import (
	"fmt"
	"github.com/darr4i/my-project/services"
	"io"
	"net/http"
)

func HandleTimeRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received time request")
	fmt.Println("Request method:", r.Method)
	data := services.GetCurrentTime()
	io.WriteString(w, string(data))
}
