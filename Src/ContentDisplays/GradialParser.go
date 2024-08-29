package parser

import (
	"fmt"
	"net/http"
)

func GradialParse(writer http.ResponseWriter, request *http.Request) {
	//I truly despise CORS with passion
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintln(writer, "mode 1")
	fmt.Fprintln(writer, "Gradial parser invoked")
}
