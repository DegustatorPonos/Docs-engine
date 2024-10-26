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
	var modeStack ModeStackNode = ModeStackNode{}
	var modeStackBuf ModeStackNode = ModeStackNode{}  // The buffer used to detect changes in stack

	// Going through the file's strings
	for index, el := range FileStrings {
		var prevString = "<Null>"
		outp := strings.Trim(strings.ReplaceAll(el, "\n", ""), " ") 
		if(index == -1) { // It doesnt build without it. Might not be useful
			continue;
		}
		if(SetMode(prevString, &outp, "", &modeStack)) {
			// If there is a difference on the stack
			if(!modeStack.EqualsTo(modeStackBuf)) {

				// Closing/opening default tag
				if(modeStackBuf.depth == 0 && index != 0) {
					fmt.Fprint(writer, TagsDict[Text].closingTag)
				}

				// Handling difference
				var toPull, toPush = modeStack.CalculateBiggestDifference(modeStackBuf)
				for _, val := range toPush {
					fmt.Fprint(writer, TagsDict[val.mode].closingTag)
				}
				for _, val := range toPull {
					fmt.Fprint(writer, TagsDict[val.mode].openingTag)
				}

				if(modeStack.depth == 0) {
					fmt.Fprint(writer, TagsDict[Text].openingTag)
				}
				// Eqalization of the stacks
				modeStackBuf = modeStack.Clone()
			}
			fmt.Fprint(writer, TransformString(outp, modeStack.mode))
		}
	}
}
