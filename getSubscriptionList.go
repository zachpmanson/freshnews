package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type item struct {
	// Define the fields for the items if there are any
	// Leaving it empty since the provided JSON shows an empty array
}

type ncFeed struct {
	ID               int         `json:"id"`
	URL              string      `json:"url"`
	Title            string      `json:"title"`
	FaviconLink      string      `json:"faviconLink"`
	Added            int64       `json:"added"`
	FolderID         int         `json:"folderId"`
	UnreadCount      int         `json:"unreadCount"`
	Ordering         int         `json:"ordering"`
	Link             string      `json:"link"`
	Pinned           bool        `json:"pinned"`
	UpdateErrorCount int         `json:"updateErrorCount"`
	LastUpdateError  interface{} `json:"lastUpdateError"`
	Items            []item      `json:"items"`
}

type ncFeedResponse struct {
	StarredCount int      `json:"starredCount"`
	Feeds        []ncFeed `json:"feeds"`
	NewestItemID int      `json:"newestItemId"`
}

type category struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// Define the structs
type frSubscription struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	Categories []category `json:"categories"`
	URL        string     `json:"url"`
	HTMLUrl    string     `json:"htmlUrl"`
	IconUrl    string     `json:"iconUrl"`
}

type frFeedResponse struct {
	Subscriptions []frSubscription `json:"subscriptions"`
}

func getNCFeeds() ([]ncFeed, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", BaseUrl+"/feeds", nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Auth", "Basic", Credentials)
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Basic " + Credentials},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var ncResponse ncFeedResponse
	err = json.Unmarshal(resBody, &ncResponse)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ncResponse.Feeds, nil
}

type folder struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Opened bool   `json:"opened"`
	Feeds  []feed `json:"feeds"`
}

type feed struct {
	// Define the fields for the feeds if there are any
	// Leaving it empty since the provided JSON shows an empty array
}

type ncFolderReponse struct {
	Folders []folder `json:"folders"`
}

func getNCFolders() ([]folder, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", BaseUrl+"/folders", nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header = http.Header{
		"Authorization": {"Basic " + Credentials},
		"Content-Type":  {"application/json"},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var ncResponse ncFolderReponse
	err = json.Unmarshal(resBody, &ncResponse)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ncResponse.Folders, nil
}

func GetSubscriptionsList(w http.ResponseWriter, r *http.Request) {

	// if output=json set
	if r.URL.Query().Get("output") != "json" {
		notImplemented(w, r)
		return
	}

	fmt.Println("Getting NC Feeds")
	ncFeeds, err := getNCFeeds()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Got NC Feeds", len(ncFeeds))

	fmt.Println("Getting NC Folders")
	ncFolders, err := getNCFolders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Got NC Folders", len(ncFolders))

	folderMap := make(map[int]string)
	for _, folder := range ncFolders {
		folderMap[folder.ID] = folder.Name
	}

	response := frFeedResponse{}

	for _, subscription := range ncFeeds {
		response.Subscriptions = append(response.Subscriptions, frSubscription{
			ID:    strconv.Itoa(subscription.ID),
			Title: subscription.Title,
			Categories: []category{
				{ID: "user/-/label/" + folderMap[subscription.FolderID], Label: folderMap[subscription.FolderID]},
			},
			URL:     subscription.URL,
			HTMLUrl: subscription.URL,
			IconUrl: subscription.FaviconLink,
		})
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	println("Returning subscriptions " + string(jsonResponse))
	io.WriteString(w, string(jsonResponse))
}
