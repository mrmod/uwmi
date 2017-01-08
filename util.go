package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(writer http.ResponseWriter, v interface{}) {
	if body, err := json.Marshal(v); err != nil {
		ServerError(writer, err)
	} else {
		writer.Write(body)
	}
}

func ServerError(writer http.ResponseWriter, err error) {
	fmt.Println("Error:", err)
	writer.WriteHeader(500)
}

func Unauthorized(writer http.ResponseWriter, v interface{}) {
	writer.WriteHeader(http.StatusUnauthorized)
}
