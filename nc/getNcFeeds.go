package nc

import (
	"encoding/json"
	"fmt"
	"freshnews/utils"
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

func GetNCFeeds() ([]ncFeed, error) {
	resBody, err := utils.NcGetReq("/feeds")
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
	fmt.Println("NC Feeds:", len(ncResponse.Feeds))

	return ncResponse.Feeds, nil
}
