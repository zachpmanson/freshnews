package main

import (
	"flag"
	"fmt"
	"freshnews/config"
	"freshnews/handlers"
	"freshnews/utils"
	"net/http"
	"net/http/httputil"
	"os"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

func dump(f HttpHandler) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		dump, _ := httputil.DumpRequest(r, true)
		fmt.Println("\n\n\n==========================")
		fmt.Println(utils.CapString(dump, 1000))

		f(w, r) // call the original function
	}
}

func attachHandler(path string, f HttpHandler) {
	http.HandleFunc(path, dump(f))
}

func main() {
	println("Starting up...")
	// read in --credentials from CLI
	flag.StringVar(&config.Credentials, "credentials", "", "credentials string in base64")
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()
	// have seen netnewswire use
	attachHandler("/", handlers.NotImplemented)
	attachHandler("/api/greader.php/accounts/ClientLogin", handlers.GetClientLogin)
	attachHandler("/api/greader.php/reader/api/0/stream/items/ids", handlers.GetStreamItemsIds) // TODO still not visible in NNN, something wrong with IDs?
	attachHandler("/api/greader.php/reader/api/0/subscription/list", handlers.GetSubscriptionsList)
	attachHandler("/api/greader.php/reader/api/0/subscription/edit", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/subscription/quickadd", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/stream/items/contents", handlers.GetStreamItemContents)
	attachHandler("/api/greader.php/reader/api/0/tag/list", handlers.GetTagList)
	attachHandler("/api/greader.php/reader/api/0/edit-tag", handlers.NotImplemented)

	// suspected to be netnewswire
	attachHandler("/api/greader.php/reader/api/0/subscription/export", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/subscription/import", handlers.NotImplemented)

	// unknown
	attachHandler("/api/greader.php/check/compatibility", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/stream/contents", handlers.NotImplemented)
	// attachHandler.HandleFunc("/api/greader.php/reader/api/0/stream/contents/feed/<include target>", getRoot)
	attachHandler("/api/greader.php/reader/api/0/stream/contents/feed/user/state/com.google/reading-list", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/stream/contents/feed/user/state/com.google/starred", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/stream/contents/", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/unread-count", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/rename-tag", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/disable-tag", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/mark-all-as-read", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/token", handlers.NotImplemented)
	attachHandler("/api/greader.php/reader/api/0/user-info", handlers.NotImplemented)

	fmt.Printf("Listening on port %d...\n", *port)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ListenAndServe: %s\n", err)
		os.Exit(1)
	}
}
