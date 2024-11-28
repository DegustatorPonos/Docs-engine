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
	// parser.Test() // <- Test the stack

	// Parsing a config file
	err := config.ReadConfigFile()
	if err != nil {
		fmt.Printf("Failed to read the nonfig file. Message: %s\n", err.Error())
		return
	}

	// Enabling configured parsing mode
	mode := config.GetParsingMode(0)
	parser.InitializeDict()
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
	http.HandleFunc("/GetAppName", GetAppName)

	port := ":" + config.GetPort("5000")
	fmt.Println("Listener opened on port " + port)
	if servererror := http.ListenAndServe(":5000", nil); !errors.Is(servererror, os.ErrClosed) {
		panic(servererror)
	}
}

func GetAppName(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprint(writer, config.GetAppName("Documentation"))
}
