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
	if(!isPathSecure(filepath)) {
		return;
	}
	content, err := os.ReadFile(config.GetSourceDirectoryPath("../Content/") + filepath)
	if err != nil {
		fmt.Fprint(writer, "Error while reading the file. \nError: "+err.Error())
	}
	FileStrings := strings.Split(string(content), "\n")

	// File parsing

	// The type of content the string represents
	var globalMode int = Text 
	var bufMode int = -1 // Set it to -1 so the first opening tag will be open

	// Going through the file's strings
	for index, el := range FileStrings {
		var prevString = "<Null>"
		outp := strings.Trim(strings.ReplaceAll(el, "\n", ""), " ") 
		if(index == -1) { // It doesnt build without it
			continue;
		}
		if(SetMode(prevString, &outp, "", &globalMode)) {
			if(bufMode != globalMode) {
				// fmt.Printf("Difference between %v and %v\n", globalMode, bufMode)
				if(bufMode != -1) {
					fmt.Fprint(writer, TagsDict[bufMode].closingTag) // Closing last section
				}
				fmt.Fprint(writer, TagsDict[globalMode].openingTag) // Opening new tag
				bufMode = globalMode
			}
			// fmt.Println(globalMode)
			fmt.Fprint(writer, TransformString((string)(outp), globalMode))
		}
	}
	// fmt.Println("Done")
}
