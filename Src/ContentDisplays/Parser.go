package parser

import (
	"fmt"
	"net/http"
	"os"
)

type Chunk struct {
	openingTag string
	closingTag string
}

const (
	Text int = 0
	Code
	H1
	H2
	H3
	H4
	H5
	H6
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
