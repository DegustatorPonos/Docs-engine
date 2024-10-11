package parser

import (
	"strings"
)

// Parses the string that contains table row to its html representation
func StringToTableRecord(rawString string) string {
	var builder strings.Builder
	builder.WriteString(TagsDict[TableRecord].openingTag)

	if(len(strings.ReplaceAll(strings.ReplaceAll(rawString, "|", ""), "-", "")) == 0) {
		// that string represents the end of the header
		return ""
	}

	for _, cell := range strings.Split(rawString, "|") {
		if(len(cell) == 0) {
			continue
		}
		builder.WriteString(TagsDict[TableCell].openingTag)
		builder.WriteString(cell)
		builder.WriteString(TagsDict[TableCell].closingTag)
	}
	builder.WriteString(TagsDict[TableRecord].closingTag)
	return builder.String()
}
