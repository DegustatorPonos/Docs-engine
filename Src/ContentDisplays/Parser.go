package parser

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Chunk struct {
	openingTag string
	closingTag string
}

const (
	Text int = 0
	Code int = 1
	H1 int = 2
	H2 int = 3
	H3 int = 4
	H4 int = 5
	H5 int = 6
	H6 int = 7
)

// A dictionary of opening and closing tags of elements
var TagsDict = map[int]Chunk{}

// Reads the file specified py 'path' query param
func ReadFile(writer http.ResponseWriter, request *http.Request) {
	filepath := request.URL.Query().Get("path")
	fmt.Printf("Parsing file" + "\n")
	//I truly despise CORS with passion
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	contents, err := os.ReadFile("../Content/" + filepath)
	if err != nil {
		fmt.Fprint(writer, "There is nod such file as "+filepath)
		return
	}
	fmt.Fprint(writer, string(contents))
}

// Initializes the TagsDict. Probably full of magical strings. And it looks scary.
func InitializeDict() {
	TagsDict[Text] = Chunk{"<p>", "</p>"}
	TagsDict[Code] = Chunk{"<div class=\"CodeElement CodeElementOverride\">", "</div>"}
	TagsDict[H1] = Chunk{"<h1>", "</h1>"}
	TagsDict[H2] = Chunk{"<h2>", "</h2>"}
	TagsDict[H3] = Chunk{"<h3>", "</h3>"}
	TagsDict[H4] = Chunk{"<h4>", "</h4>"}
	TagsDict[H5] = Chunk{"<h5>", "</h5>"}
	TagsDict[H6] = Chunk{"<h6>", "</h6>"}
}

// Gets the string HTML tags and transforms spectial sumbols to HTML equivalents
func TransformString(input string, globalTag int, includeClosingTag bool, includeOpeningTag bool) string {
	var correctedInput string = input
	correctedInput = strings.ReplaceAll(correctedInput, "<", "&lt")
	correctedInput = strings.ReplaceAll(correctedInput, ">", "&gt")

	modeTags := TagsDict[globalTag]
	outp := ""
	if (includeOpeningTag) {
		outp += modeTags.openingTag
	}
	outp += correctedInput
	outp += "<br>"
	if (includeClosingTag) {
		outp += modeTags.closingTag
	}
	return outp
}

// =============================== Mode changes =============================== 

// Sets the mode to the current one and modifies the string if needed. Returns true if we need to incude this line
func SetMode(previousString string, currentString *string, nextString string, contextMode *int) bool {
	// Checking for the block code
	if(CheckForCodeBlock(*currentString, contextMode)) {
		return false
	}

	return true
}

// By specs the code block is defined by tripple backticks (```) that are not included.
// This function sets the mode value and returns true if the line is a code block identifier
func CheckForCodeBlock(currentString string, contextMode *int) bool {
	if(currentString == "```") {
		if(*contextMode != Code) {
			*contextMode = Code
		} else {
			*contextMode = Text
		}
		return true 
	}
	return false
}
