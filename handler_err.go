package main

import (
	"net/http"
)

func handlerErr(w http.ResponseWriter, _ *http.Request) {
	resposeWithErr(w, 400, "something went wrong")
}
