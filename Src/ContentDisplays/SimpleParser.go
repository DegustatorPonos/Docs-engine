package parser

import (
	"fmt"
	"net/http"
)

func SimpleParse(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintln(writer, "mode 0")
	fmt.Fprintln(writer, "Simple parser invoked")
}
