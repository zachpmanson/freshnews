package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

var Credentials string
var BaseUrl string = "https://nextcloud.zachmanson.com/index.php/apps/news/api/v1-3"

func notImplemented(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Not Implemented!\n")
}

type HttpHandler func(w http.ResponseWriter, r *http.Request)

func dump(f HttpHandler) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		dump, _ := httputil.DumpRequest(r, true)
		fmt.Println("\n==========================")
		fmt.Println(string(dump))

		f(w, r) // call the original function
	}
}

func attachHandler(path string, f HttpHandler) {
	http.HandleFunc(path, dump(f))
}

func main() {
	println("Starting up...")
	// read in --credentials from CLI
	flag.StringVar(&Credentials, "credentials", "", "credentials string in base64")
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()
	// have seen netnewswire use
	attachHandler("/api/greader.php/accounts/ClientLogin", GetClientLogin)
	attachHandler("/api/greader.php/reader/api/0/stream/items/ids", GetStreamItemsIds)
	attachHandler("/api/greader.php/reader/api/0/subscription/list", GetSubscriptionsList) // done maybe?
	attachHandler("/api/greader.php/reader/api/0/subscription/edit", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/subscription/quickadd", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/stream/items/contents", GetStreamItemContents)
	attachHandler("/api/greader.php/reader/api/0/tag/list", GetTagList)
	attachHandler("/api/greader.php/reader/api/0/edit-tag", notImplemented)

	// suspected to be netnewswire
	attachHandler("/api/greader.php/reader/api/0/subscription/export", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/subscription/import", notImplemented)

	// unknown
	attachHandler("/api/greader.php/check/compatibility", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/stream/contents", notImplemented)
	// attachHandler.HandleFunc("/api/greader.php/reader/api/0/stream/contents/feed/<include target>", getRoot)
	attachHandler("/api/greader.php/reader/api/0/stream/contents/feed/user/state/com.google/reading-list", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/stream/contents/feed/user/state/com.google/starred", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/stream/contents/", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/unread-count", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/rename-tag", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/disable-tag", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/mark-all-as-read", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/token", notImplemented)
	attachHandler("/api/greader.php/reader/api/0/user-info", notImplemented)

	fmt.Printf("Listening on port %d...\n", *port)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ListenAndServe: %s\n", err)
		os.Exit(1)
	}
}
