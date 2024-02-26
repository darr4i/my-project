package controllers

import (
	"fmt"
	"github.com/darr4i/my-project/services"
	"io"
	"net/http"
)

func HandleTimeRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Received time request\n")
    fmt.Printf("Request method:", r.Method)
	data := services.GetCurrentTime()
	io.WriteString(w, string(data))
}

