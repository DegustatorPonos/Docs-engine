package parser

import (
	"slices"
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
	Quote int = 8
	SepLine int = 9
	// Table contents 
	TableGeneral int = 10
	TableRecord int = 11
	TableCell int = 12
	TableHeaderCell int = 13
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
	TagsDict[SepLine] = Chunk{"<hr>", ""}
	TagsDict[Quote] = Chunk{"<div class=\"QuoteElement QuoteElementOverride\">", "</div>"}
	// Table elements 
	TagsDict[TableGeneral] = Chunk{"<table class=\"TableElement TableElementOverride\">", "</table>"} 
	TagsDict[TableRecord] = Chunk{"<tr class=\"TableRecordElement TableRecordElementOverride\">", "</tr>"}
	TagsDict[TableCell] = Chunk{"<td class=\"TableCellElement TableCellElementOverride\">", "</td>"}
	TagsDict[TableHeaderCell] = Chunk{"<td class=\"TableHeaderCellElement TableHeaderCellElementOverride\">", "</td>"}
}

// Checks wheather it is secure to read this file or not
func isPathSecure(path string) bool {
	if(strings.Contains(path, "..")) {
		return false
	}
	return true
}

// Gets the string HTML tags and transforms spectial sumbols to HTML equivalents
func TransformString(input string, globalTag int) string {
	if(len(input) == 0) {
		return "<br>"
	}
	var correctedInput string = input
	correctedInput = strings.ReplaceAll(correctedInput, "<", "&lt")
	correctedInput = strings.ReplaceAll(correctedInput, ">", "&gt")
	// TODO: Find a better way to do it
	if(globalTag == TableGeneral) {
		return StringToTableRecord(correctedInput)
	}
	outp := ""
	outp += correctedInput
	outp += "<br>"
	return outp
}

// =============================== Mode changes =============================== 

// Sets the mode to the current one and modifies the string if needed. Returns true if we need to incude this line
func SetMode_legacy(previousString string, currentString *string, nextString string, contextMode *int) bool {
	// Checking for the comment
	if(CheckForCommentBlock(*currentString)) {
		return false
	}

	// Checking for the table
	if(CheckForTable(&currentString, contextMode)) {
		return true
	}

	return true
}

// Modifies stack according to the current string. Returns true if the string must be included. 
func SetMode(previousString string, currentString *string, nextSeting string, modeStack *ModeStackNode) bool {

	if(CheckForCommentBlock(*currentString)) {
		return false
	}

	if(CheckForQuoteBlock(&currentString, modeStack)) {
		return true
	}

	if(CheckForHeaderBlock(&currentString, nextSeting, modeStack)) {
		return true
	}

	if(CheckForCodeBlock(*currentString, modeStack)) {
		return false
	}

	return true
}

// By specs quote is set by placing '>' symbol before the string
// This function sets the mode value and returns true if the line is a quote block identifier
func CheckForQuoteBlock(currentString **string, contextMode *ModeStackNode) bool { 
	if(strings.HasPrefix(**currentString, "> ")) {
		**currentString = strings.TrimLeft(**currentString, "> ")
		*contextMode = contextMode.Push(Quote)
		return true
	} else {
		if(contextMode.mode == Quote) {
			contextMode.Pull()
		}
	}
	return false
}

// By specs the code block is defined by tripple backticks (```) that are not included.
// This function sets the mode value and returns true if the line is a code block identifier
func CheckForCodeBlock(currentString string, contextMode *ModeStackNode) bool {
	if(currentString == "```") {
		if(*&contextMode.mode != Code) {
			*contextMode = contextMode.Push(Code)
		} else {
			contextMode.Pull()
		}
		return true 
	}
	return false
}

// By specs the header block can be defined by two ways:
// First - '#' symbols before the string. The ammount of strings is the depth. Max 6
// Second - strings that only contains === for H1 and --- for H2. Low priority impl
// This function sets the mode value and returns true if the line is a header block identifier
func CheckForHeaderBlock(currentString **string, nextString string, contextMode *ModeStackNode) bool { 
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
	for slices.Contains([]int{H1, H2, H3, H4, H5, H6}, contextMode.mode) {
		contextMode.Pull() 
	}
	if(depth == 0) {
		return false
	} else {
		*contextMode = contextMode.Push(H1 - 1 + depth )
		**currentString = strings.TrimLeft(**currentString, "#")
		return true
	}
}

// by specs comment is set by placing text between '<!--' and '-->' blocks
// this function sets the returns true if the line is a comment 
func CheckForCommentBlock(currentString string) bool { 
	if(strings.HasPrefix(currentString, "<!-- ") && strings.HasSuffix(currentString, " -->")) {
		return true
	}
	return false
}

// by specs separator is set as '---'
// this function sets the returns true if the line is a comment 
// !IMPORTANT - if i remember correctly this only should apply in hiding list.
// I might or might not be speculating but still
func CheckForSeparartor(currentString string) bool {
	if(len(strings.Trim(currentString, "-")) == 0) {
		return true
	}
	return false
} // TODO

// TODO: Write the table conditions here
// https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/organizing-information-with-tables 
func CheckForTable(currentString **string, contextMode *int) bool {
	if(strings.HasPrefix(**currentString, "|") && strings.HasSuffix(**currentString, "|")) {
		// If we reach this part of code then we are in the table	
		if(*contextMode != TableGeneral) {
			*contextMode = TableGeneral
		} else {
		}
		return true
	} else {
		if(*contextMode == TableGeneral) {
			*contextMode = Text
		}
	}
	return false
}
