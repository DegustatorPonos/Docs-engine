package main

import (
	"PaketikDocsEngine/config"
	"PaketikDocsEngine/directories"
	"fmt"
	"net/http"
)

func main() {
	// Parsing a config file
	config.ReadConfigFile()
	fmt.Println("Listener opened on port " + config.GetPort("5000"))
	http.HandleFunc("/GetDirectories", directories.GetDirectory)
	http.ListenAndServe(":"+config.GetPort("5000"), nil)
}
