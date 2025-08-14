package main

import (
	"io"
	"net/http"
)
func GetTagList(w http.ResponseWriter, r *http.Request) {
	// if output=json set
	if r.URL.Query().Get("output") == "json" {
		// if output=json, return json
		io.WriteString(w, "[{\"id\":\"/user/zach\",\"title\":\"My News Feed\"}]")
		return
	} else {
		notImplemented(w, r)
	}
	io.WriteString(w, `{
		"tags": [
		  { "id": "user/-/state/com.google/starred" },
		  { "id": "user/-/label/m", "type": "folder" }
		]
	  }`)
}