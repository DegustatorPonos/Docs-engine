package parser

import (
	"fmt"
	"net/http"
)

func GradialParse(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Gradial parser invoked")
}
