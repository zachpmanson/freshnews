package handlers

import (
	"io"
	"net/http"
)

func NotImplemented(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Not Implemented!\n")
}
