package parser

import (
	"fmt"
	"net/http"
)

func PreprocessingParse(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Preprocessing parser invoked")
}
