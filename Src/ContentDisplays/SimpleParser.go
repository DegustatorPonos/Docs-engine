package parser

import (
	"fmt"
	"net/http"
)

func SimpleParse(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Simple parser invoked")
}
