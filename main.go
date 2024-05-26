package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getClientLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /api/greader.php/accounts/ClientLogin request\n")
	// read email and passwd from request query params
	email := r.URL.Query().Get("Email")
	passwd := r.URL.Query().Get("Passwd")

	// check if email and passwd are valid
	if email == "test@example.com" && passwd == "password" {
		// if valid, return a valid token
		io.WriteString(w, "SID=zach/{some id code}\nLSID=null\nAuth=zach/{some id code}")
	}
}

func main() {
	http.HandleFunc("/api/greader.php/check/compatibility", getRoot)
	http.HandleFunc("/api/greader.php/accounts/ClientLogin", getClientLogin)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/contents", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/contents/feed/<include target>", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/contents/feed/user/state/com.google/reading-list", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/contents/feed/user/state/com.google/starred", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/contents/", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/items/ids", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/items/contents", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/tag/list", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/subscription/export", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/subscription/import", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/subscription/list", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/subscription/edit", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/subscription/quickadd", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/unread-count", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/edit-tag", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/rename-tag", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/disable-tag", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/mark-all-as-read", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/token", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/user-info", getRoot)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ListenAndServe: %s\n", err)
		os.Exit(1)
	}
}
