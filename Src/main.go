package main

import (
	parser "PaketikDocsEngine/ContentDisplays"
	directories "PaketikDocsEngine/DirectoriesControllers"
	"PaketikDocsEngine/config"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Parsing a config file
	config.ReadConfigFile()
	fmt.Println("Listener opened on port " + config.GetPort("5000"))

	// Enabling configured parsing mode
	mode := config.GetParsingMode(0)
	fmt.Println("Selected parsing mode: " + strconv.Itoa(mode))
	switch mode {
	case 0: // Basic aka simple
		http.HandleFunc("/ReadFile", parser.SimpleParse)
	case 1: // Gradial
		http.HandleFunc("/ReadFile", parser.GradialParse)
	case 2: // Preprocessing
		http.HandleFunc("/ReadFile", parser.PreprocessingParse)
	}
	http.HandleFunc("/GetDirectories", directories.GetDirectory)

	port := ":" + config.GetPort("5000")
	fmt.Println(port)
	if servererror := http.ListenAndServe(port, nil); !errors.Is(servererror, os.ErrClosed) {
		panic(servererror)
	}
}
