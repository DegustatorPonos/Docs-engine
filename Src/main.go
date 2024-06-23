package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func ReadConfigFile() []byte {
	content, err := os.ReadFile("Config.cfg")
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func main() {
	http.HandleFunc("/GetDirectories", GetDirectory)
	http.ListenAndServe(":5000", nil)
}

func GetDirectory(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Directory requested")

	absPath, _ := filepath.Abs("../Content")

	content, err := os.ReadDir(absPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range content {
		fmt.Fprint(writer, e.Name()+" ")
	}

}
