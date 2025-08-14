package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var Credentials string
var BaseUrl string = "https://nextcloud.zachmanson.com/index.php/apps/news/api/v1-3"

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "TO IMPLEMENT!\n")
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Not Implemented!")
}


func main() {
	println("Starting up...")
	// read in --credentials from CLI
	flag.StringVar(&Credentials, "credentials", "", "credentials string in base64")
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	// have seen netnewswire use
	http.HandleFunc("/api/greader.php/accounts/ClientLogin", GetClientLogin)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/items/ids", GetStreamItemsIds)
	http.HandleFunc("/api/greader.php/reader/api/0/subscription/list", GetSubscriptionsList) // done maybe?
	http.HandleFunc("/api/greader.php/reader/api/0/subscription/edit", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/subscription/quickadd", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/items/contents", GetStreamItemContents)
	http.HandleFunc("/api/greader.php/reader/api/0/tag/list", GetTagList)
	http.HandleFunc("/api/greader.php/reader/api/0/edit-tag", getRoot)

	// suspected to be netnewswire
	http.HandleFunc("/api/greader.php/reader/api/0/subscription/export", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/subscription/import", getRoot)

	// unknown
	http.HandleFunc("/api/greader.php/check/compatibility", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/contents", getRoot)
	// http.HandleFunc("/api/greader.php/reader/api/0/stream/contents/feed/<include target>", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/contents/feed/user/state/com.google/reading-list", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/contents/feed/user/state/com.google/starred", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/stream/contents/", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/unread-count", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/rename-tag", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/disable-tag", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/mark-all-as-read", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/token", getRoot)
	http.HandleFunc("/api/greader.php/reader/api/0/user-info", getRoot)

	fmt.Printf("Listening on port %d...\n", *port)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ListenAndServe: %s\n", err)
		os.Exit(1)
	}
}
