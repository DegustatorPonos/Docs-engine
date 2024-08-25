package main

import (
	"PaketikDocsEngine/config"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Parsing a config file
	config.ReadConfigFile()
	fmt.Println("Listener opened on port " + config.GetPort("5000"))
	http.HandleFunc("/GetDirectories", GetDirectory)
	http.ListenAndServe(":"+config.GetPort("5000"), nil)
}

func GetDirectory(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Directory requested\n")
	//I truly despise CORS with passion
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var requestedPath = request.URL.Query().Get("path")
	absPath, _ := filepath.Abs("../Content/" + requestedPath)
	fmt.Fprint(writer, ReadDirectory(absPath))
}

//================== DIRECTORY MANAGER ==================

// Recursive function to read directoory content and
// the content of all included directories
func ReadDirectory(path string) string {
	var builder strings.Builder

	content, err := os.ReadDir(path)
	//If something went wrong it's more reliable to return nothing IMO
	if err != nil {
		return ""
	}
	var isFirst = true
	//Read the directory
	for _, e := range content {
		//TODO: Get rid of this flag and do it some better way i don't know about yet
		if !isFirst {
			fmt.Fprint(&builder, ";")
		}

		//Output type separation
		if !e.IsDir() {
			fmt.Fprint(&builder, strings.Replace(e.Name(), ".md", " -f", 1))
		} else {
			fmt.Fprint(&builder, e.Name()+" -d")
		}
		isFirst = false
	}
	return builder.String()
}
