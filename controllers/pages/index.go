package pages

import (
	"net/http"
	"fmt"
)

func Index(rw http.ResponseWriter) {
	fmt.Fprint(rw, "Hello World")
}