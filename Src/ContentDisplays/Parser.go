package parser

import (
	"fmt"
	"net/http"
	"os"
)

func GetFile(writer http.ResponseWriter, request *http.Request) {
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
