package utils

import (
	"fmt"
	"net/http"
	"os"
)

func WriteBytes(w http.ResponseWriter, b []byte, full bool) {
	fmt.Println("Writing Res:")
	if full {
		// fmt.Print(string(b))

		toFile(b, "freshnews_res.dump")
		fmt.Print(CapString(b, 300))

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

func toFile(bytes []byte, filename string) {
	e := os.WriteFile(filename, bytes, 0644)
	if e != nil {
		panic(e)
	}
	// f, err := os.Create("/tmp/dat2")
	// if e != nil {
	// 	panic(e)
	// }
	// defer f.Close()
}
