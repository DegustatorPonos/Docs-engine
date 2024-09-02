package parser

import (
	"PaketikDocsEngine/config"
	"fmt"
	"net/http"
	"os"
)

func SimpleParse(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintln(writer, "mode 0")
	// File reading
	filepath := request.URL.Query().Get("path")
	content, err := os.ReadFile(config.GetSourceDirectoryPath("../Content/") + filepath)
	if err != nil {
		fmt.Fprint(writer, "Error while reading the file. \nError: "+err.Error())
	}

	strings := string(content)
	fmt.Fprint(writer, strings)
}
