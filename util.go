package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

func ParseTime(timeString string) time.Time {
	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		log.Println("Error parsing time:", err)
	}
	return t
}
