package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Item struct {
	// Define the fields for the items if there are any
	// Leaving it empty since the provided JSON shows an empty array
}

type NCFeed struct {
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
	Items            []Item      `json:"items"`
}

type NCFeedResponse struct {
	StarredCount int      `json:"starredCount"`
	Feeds        []NCFeed `json:"feeds"`
	NewestItemID int      `json:"newestItemId"`
}

type Category struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// Define the structs
type FRSubscription struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	Categories []Category `json:"categories"`
	URL        string     `json:"url"`
	HTMLUrl    string     `json:"htmlUrl"`
	IconUrl    string     `json:"iconUrl"`
}

type FRResponse struct {
	Subscriptions []FRSubscription `json:"subscriptions"`
}

func getNCFeeds() ([]NCFeed, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", BaseUrl+"/feeds", nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

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

	var ncResponse NCFeedResponse
	err = json.Unmarshal(resBody, &ncResponse)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ncResponse.Feeds, nil
}

type Folder struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Opened bool   `json:"opened"`
	Feeds  []Feed `json:"feeds"`
}

type Feed struct {
	// Define the fields for the feeds if there are any
	// Leaving it empty since the provided JSON shows an empty array
}

type NCFolderReponse struct {
	Folders []Folder `json:"folders"`
}

func getNCFolders() ([]Folder, error) {
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

	var ncResponse NCFolderReponse
	err = json.Unmarshal(resBody, &ncResponse)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ncResponse.Folders, nil
}

func GetSubscriptionsList(w http.ResponseWriter, r *http.Request) {

	// if output=json set
	if r.URL.Query().Get("output") == "json" {
		// if output=json, return json
		fmt.Println("Getting NC Feeds")
		ncFeeds, err := getNCFeeds()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Getting NC Folders")
		ncFolders, err := getNCFolders()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		folderMap := make(map[int]string)
		for _, folder := range ncFolders {
			folderMap[folder.ID] = folder.Name
		}

		response := FRResponse{}

		for _, subscription := range ncFeeds {
			response.Subscriptions = append(response.Subscriptions, FRSubscription{
				ID:    strconv.Itoa(subscription.ID),
				Title: subscription.Title,
				Categories: []Category{
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

		io.WriteString(w, string(jsonResponse))
		return
	} else {
		notImplemented(w, r)
	}

}
