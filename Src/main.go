package main

import (
	parser "PaketikDocsEngine/ContentDisplays"
	directories "PaketikDocsEngine/DirectoriesControllers"
	"PaketikDocsEngine/config"
	"fmt"
	"net/http"
)

func main() {
	// Parsing a config file
	config.ReadConfigFile()
	fmt.Println("Listener opened on port " + config.GetPort("5000"))
	http.HandleFunc("/GetDirectories", directories.GetDirectory)
	http.HandleFunc("/ReadFile", parser.GetFile)
	http.ListenAndServe(":"+config.GetPort("5000"), nil)
}
