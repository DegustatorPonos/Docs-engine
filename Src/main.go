package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/GetDirectories", GetDirectory)
	http.ListenAndServe(":5000", nil)
}

func GetDirectory(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Directory requested")
}
