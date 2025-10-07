package utils

import (
	"fmt"
	"net/http"
)

func WriteBytes(w http.ResponseWriter, b []byte, full bool) {
	fmt.Println("Writing Res:")
	if full {
		fmt.Print(string(b))
	} else {
		fmt.Print(CapString(b, 300))
	}
	w.Write(b)
}

func WriteString(w http.ResponseWriter, s string, full bool) {
	WriteBytes(w, []byte(s), full)
}

func WriteError(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
	fmt.Println("Error:", err)
}
