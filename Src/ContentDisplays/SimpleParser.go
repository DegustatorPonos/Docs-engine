package parser

import (
	"PaketikDocsEngine/config"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func SimpleParse(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintln(writer, "mode 0") // Throw mode to the frontend

	// File reading
	filepath := request.URL.Query().Get("path")
	content, err := os.ReadFile(config.GetSourceDirectoryPath("../Content/") + filepath)
	if err != nil {
		fmt.Fprint(writer, "Error while reading the file. \nError: "+err.Error())
	}
	FileStrings := strings.Split(string(content), "\n")

	// File parsing

	// The type of content the string represents
	var globalMode int = Text 
	// var bufMode int = Text

	// Going through the file's strings
	for index, el := range FileStrings {
		var prevString = "<Null>"
		outp := strings.Trim(strings.ReplaceAll(el, "\n", ""), " ") 
		if(len(outp) == 0 || index == -1) {
			continue;
		}
		// For explanation check the defenition of SetMoee() function
		if(SetMode(prevString, &outp, "", &globalMode)) {
			fmt.Fprintln(writer, TransformString((string)(outp), globalMode))
		}
	}
}
