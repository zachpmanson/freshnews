package handlers

import (
	"encoding/json"
	"fmt"
	"freshnews/nc"
	"freshnews/utils"
	"net/http"
	"strconv"
)

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

func GetSubscriptionsList(w http.ResponseWriter, r *http.Request) {

	// if output=json set
	if r.URL.Query().Get("output") != "json" {
		NotImplemented(w, r)
		return
	}

	fmt.Println("Getting NC Feeds")
	ncFeeds, err := nc.GetNCFeeds()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Got NC Feeds", len(ncFeeds))

	fmt.Println("Getting NC Folders")
	ncFolders, err := nc.GetFolders()
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

	utils.WriteBytes(w, jsonResponse, false)
}
