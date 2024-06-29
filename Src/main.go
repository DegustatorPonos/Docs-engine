package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	http.HandleFunc("/GetDirectories", GetDirectory)
	http.ListenAndServe(":5000", nil)
}

func GetDirectory(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Directory requested\n")

	absPath, _ := filepath.Abs("../Content")
	fmt.Fprint(writer, ReadDirectory(absPath))
}

// Recursive function to read directoory content and
// the content of all included directories
func ReadDirectory(path string) string {
	var builder strings.Builder

	content, err := os.ReadDir(path)
	//If something went wrong it's more reliable to return nothing IMO
	if err != nil {
		return ""
	}
	fmt.Fprint(&builder, "(")
	var isFirst = true
	//Read the directory
	for _, e := range content {
		if !isFirst {
			fmt.Fprint(&builder, ";")
		}
		fmt.Fprint(&builder, e.Name())
		if e.IsDir() {
			fmt.Fprint(&builder, ReadDirectory(path+"/"+e.Name()))
		}
		isFirst = false
	}
	fmt.Fprint(&builder, ")")
	return builder.String()
}
