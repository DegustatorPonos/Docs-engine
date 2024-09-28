package parser

import (
	"fmt"
	"slices"
	"strings"
)

type Chunk struct {
	openingTag string
	closingTag string
}

const (
	Text int = 0
	Quote int = 8
	Code int = 1
	H1 int = 2
	H2 int = 3
	H3 int = 4
	H4 int = 5
	H5 int = 6
	H6 int = 7
	// Table contents 
	TableGeneral int = 10
	TableRecord int = 11
	TableCell int = 12
)

// A dictionary of opening and closing tags of elements
var TagsDict = map[int]Chunk{}

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
	TagsDict[Quote] = Chunk{"<div class=\"QuoteElement QuoteElementOverride\">", "</div>"}
	// Table elements 
	TagsDict[TableGeneral] = Chunk{"<table>", "</table>"} // TODO: add style classes
	TagsDict[TableRecord] = Chunk{"<tr>", "</tr>"}
	TagsDict[TableCell] = Chunk{"<td>", "</td>"}
}

// Checks wheather it is secure to read this file or not
func isPathSecure(path string) bool {
	if(strings.Contains(path, "..")) {
		return false
	}
	return true
}

// Gets the string HTML tags and transforms spectial sumbols to HTML equivalents
func TransformString(input string, globalTag int, includeClosingTag bool, includeOpeningTag bool) string {
	if(len(input) == 0) {
		return "<br>"
	}
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
	// 	outp += modeTags.closingTag
	}
	return outp
}

// =============================== Mode changes =============================== 

// Sets the mode to the current one and modifies the string if needed. Returns true if we need to incude this line
func SetMode(previousString string, currentString *string, nextString string, contextMode *int) bool {
	// Checking for the comment
	if(CheckForCommentBlock(*currentString)) {
		return false
	}

	// Checking for the code block
	if(CheckForCodeBlock(*currentString, contextMode)) {
		return false
	}

	// Checking for the header
	if(CheckForHeaderBlock(&currentString, nextString, contextMode)) {
		return true
	}

	// Checking for the quote 
	if(CheckForQuoteBlock(&currentString, contextMode)) {
		return true
	}

	// Checking for the table
	if(CheckForTable(&currentString, contextMode)) {
		return true
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

// By specs the header block can be defined by two ways:
// First - '#' symbols before the string. The ammount of strings is the depth. Max 6
// Second - strings that only contains === for H1 and --- for H2. Low priority impl
// This function sets the mode value and returns true if the line is a header block identifier
func CheckForHeaderBlock(currentString **string, nextString string, contextMode *int) bool { 
	var builder strings.Builder
	var depth int = 0
	for i := range 6 {
		builder.WriteRune('#')
		// If the string contains this depth's prefix
		if(strings.HasPrefix(**currentString, builder.String())) {
			depth = i + 1
		} else {
			break
		}
	}
	if(depth == 0) {
		if(slices.Contains([]int{H1, H2, H3, H4, H5, H6}, *contextMode)) {
			*contextMode = Text
		}
		return false
	} else {
		*contextMode = (H1 - 1 + depth )
		**currentString = strings.TrimLeft(**currentString, "#")
		return true
	}
}

// By specs quote is set by placing '>' symbol before the string
// This function sets the mode value and returns true if the line is a quote block identifier
func CheckForQuoteBlock(currentString **string, contextMode *int) bool { 
	if(strings.HasPrefix(**currentString, "> ")) {
		**currentString = strings.TrimLeft(**currentString, "> ")
		*contextMode = Quote
		return true
	} else {
		if(*contextMode == Quote) {
			*contextMode = Text
		}
	}
	return false
}

// by specs comment is set by placing text between '<!--' and '-->' blocks
// this function sets the returns true if the line is a comment 
func CheckForCommentBlock(currentString string) bool { 
	if(strings.HasPrefix(currentString, "<!-- ") && strings.HasSuffix(currentString, " -->")) {
		return true
	}
	return false
}

// by specs separator is set as ''
// this function sets the returns true if the line is a comment 
//  func CheckForSeparartor(currentString string) bool {
//  } // TODO

// TODO: Write the table conditions here
// https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/organizing-information-with-tables 
func CheckForTable(currentString **string, contextMode *int) bool {
	if(strings.HasPrefix(**currentString, "|") && strings.HasSuffix(**currentString, "|")) {
		// If we reach this part of code then we are in the table	
		if(*contextMode != TableGeneral) {
			fmt.Println("Table opening")
			*contextMode = TableGeneral
		} else {
			fmt.Println("Table body")
		}
		return true
	} else {
		if(*contextMode == TableGeneral) {
			*contextMode = Text
		}
	}
	return false
}

// Parses the string by '|' symbols and transforms it as HTML
// table record (<tr>)
func StringToTableRecord(rawString string) string {
	var builder strings.Builder
	builder.WriteString(TagsDict[TableRecord].openingTag)
	for _, cell := range strings.Split(rawString, "|") {
		if(len(cell) == 0) {
			continue
		}
		builder.WriteString(TagsDict[TableCell].openingTag)
		builder.WriteString(cell)
		builder.WriteString(TagsDict[TableCell].openingTag)
	}
	builder.WriteString(TagsDict[TableRecord].closingTag)
	return builder.String()
}
