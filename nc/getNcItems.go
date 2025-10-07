package nc

import (
	"encoding/json"
	"fmt"
	"freshnews/utils"
	"strconv"
)

type NcItem struct {
	ID               int    `json:"id"`
	Guid             string `json:"guid"`
	GuidHash         string `json:"guidHash"`
	Url              string `json:"url"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	PubDate          int64  `json:"pubDate"`
	UpdatedDate      int    `json:"updatedDate"`
	Body             string `json:"body"`
	EnclosureMime    string `json:"enclosureMime"`
	EnclosureLink    string `json:"enclosureLink"`
	MediaThumbnail   string `json:"mediaThumbnail"`
	MediaDescription string `json:"mediaDescription"`
	FeedId           int    `json:"feedId"`
	Unread           bool   `json:"unread"`
	Starred          bool   `json:"starred"`
	LastModified     int    `json:"lastModified"`
	Rtl              bool   `json:"rtl"`
	Fingerprint      string `json:"fingerprint"`
	ContentHash      string `json:"contentHash"`
}

type NcItemResponse struct {
	Items []NcItem `json:"items"`
}

func GetNCItems(count int, since int, ignoreRead bool) ([]NcItem, error) {
	path := ""
	if since == 0 {
		path += "/items?"
		path += "batchSize=" + strconv.Itoa(count) + "&"
		if ignoreRead {
			path += "getRead=false&"
		}
	} else {
		path += "/items/updated?lastModified=" + strconv.Itoa(since)
	}

	res, err := utils.NcGetReq(path)
	if err != nil {
		return nil, err
	}

	var ncResponse NcItemResponse
	err = json.Unmarshal(res, &ncResponse)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	items := ncResponse.Items
	if len(items) > count {
		items = items[:count]
	}

	return items, nil
}
