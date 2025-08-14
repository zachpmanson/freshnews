package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetClientLogin(w http.ResponseWriter, r *http.Request) {
	// read email and passwd from request query params
	email := r.URL.Query().Get("Email")
	passwd := r.URL.Query().Get("Passwd")

	fmt.Println("got /api/greader.php/accounts/ClientLogin request", email, passwd)
	// check if email and passwd are valid
	if email == "zach" && passwd == "password" {
		// if valid, return a valid token
		io.WriteString(w, "SID=zach/{some id code}\nLSID=null\nAuth=zach/{some id code}")
	}
}
