package parser

import (
	"fmt"
	"net/http"
)

func PreprocessingParse(writer http.ResponseWriter, request *http.Request) {
	//I truly despise CORS with passion
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintln(writer, "mode 2")
		fmt.Fprintln(writer, "Preprocessing parser invoked")
}
