package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(writer http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "Hello World")
}
