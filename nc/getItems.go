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

type ItemType string

const (
	TypeFeed    ItemType = "feed"
	TypeFolder  ItemType = "folder"
	TypeStarred ItemType = "starred"
	TypeAll     ItemType = "all"
)

var stateName = map[ItemType]int{
	TypeFeed:    0,
	TypeFolder:  1,
	TypeStarred: 2,
	TypeAll:     3,
}

func GetItems(count int, since int, ignoreRead bool, itemType ItemType) ([]NcItem, error) {
	path := ""
	args := [][2]string{
		{"type", strconv.Itoa(stateName[itemType])},
	}
	if since == 0 {
		path = "/items"
		args = append(args, [2]string{"limit", strconv.Itoa(count)})

		if ignoreRead {
			args = append(args, [2]string{"getRead", "false"})
		}
	} else {
		path = "/items/updated"
		args = append(args, [2]string{"lastModified", strconv.Itoa(since)})
	}
	path += "?" + utils.ToQueryParams(
		args,
	)

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
	if count >= 0 && len(items) > count {
		items = items[:count]
	}

	return items, nil
}
