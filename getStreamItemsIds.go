package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

/*
# All items
s=user/-/state/com.google/reading-list

# Read items
s=user/-/state/com.google/read

# Starred items
s=user/-/state/com.google/starred

# Liked items
s=user/-/state/com.google/like

# Folder (CHECK IF USED BY NNN)
s=user/-/label/...

# Subscription (CHECK IF USED BY NNN)
s=feed/00157a17b192950b65be3791
*/

/*
 {
      "id": 6703,
      "guid": "tag:daringfireball.net,2024:/feeds/sponsors//11.41045",
      "guidHash": "2bb97b62fed0d414becbb49f83a01855",
      "url": "https://1password.com/daringfireball",
      "title": "[Sponsor] 1Password: You Can’t Ignore GDPR Anymore",
      "author": "Daring Fireball Department of Commerce",
      "pubDate": 1721785723,
      "updatedDate": null,
      "body": "\n<p>When the EU enacted GDPR in 2018, executives and security professionals waited anxiously to see how the law would be enforced. And then they kept waiting ... and waiting ... but the Great European Privacy Crackdown never came. For a while it seemed like the only way you’d get slapped with a GDPR fine was to do something truly egregious or be named Mark Zuckerberg. (Or preferably both.)</p>\n\n<p>But the days of betting that you’re too big or too small to be noticed by GDPR are over. Recently, EU member nations (plus the UK) have started taking action against data controllers of all sizes–from the big (Amazon), to the medium (a trucking company), to the truly minuscule (a Spanish citizen whose home security cameras bothered their neighbors).</p>\n\n<p>So what changed between 2018 and 2024? Perhaps the biggest factor was the EU cracking down on companies putting bogus “headquarters” in countries with friendly regulators, particularly <a target=\"_blank\" rel=\"noreferrer\" target=\"_blank\" rel=\"noreferrer\" href=\"https://www.irishtimes.com/opinion/2023/01/23/irelands-data-commissioner-out-of-step-with-european-peers/\">Ireland</a>. But it certainly didn’t help that the last few years have seen an unending tide of data breach stories, and the public’s relationship with tech has increasingly soured. There’s an <em>appetite</em> for enforcement these days, and it’ll probably get worse before it gets better.</p>\n\n<p>If you’re an IT or security professional, you may be wondering what to do with this information. Unfortunately, GDPR compliance isn’t the kind of thing you can solve by buying a tool or scheduling a training session. The best place to start is to adopt a policy of data minimization: collect only the data you truly need to function, on both customers and employees.</p>\n\n<p>After that, your second priority must be securing the data you have. Of course, that’s easier said than done, but you can start with doing more to protect against common breach culprits like compromised passwords. (Call us biased, but getting a password manager for every employee really is table stakes for good security.) You also need to monitor where all your data is going, so PII doesn’t disappear onto Shadow IT apps and unmanaged devices.</p>\n\n<p>We’ll close with a <a target=\"_blank\" rel=\"noreferrer\" target=\"_blank\" rel=\"noreferrer\" href=\"https://ico.org.uk/about-the-ico/media-centre/news-and-blogs/2022/10/biggest-cyber-risk-is-complacency-not-hackers/\">2022 quote from John Edwards</a>, the UK Information Commissioner:</p>\n\n<blockquote>\n  <p>“The biggest cyber risk businesses face is not from hackers outside of their company, but from complacency within their company. If your business doesn’t regularly monitor for suspicious activity in its systems and fails to act on warnings, or doesn’t update software and fails to provide training to staff, you can expect a similar fine from my office.”</p>\n</blockquote>\n\n<p>In other words: it’s time to get serious about GDPR.</p>\n\n<p><a target=\"_blank\" rel=\"noreferrer\" target=\"_blank\" rel=\"noreferrer\" href=\"https://blog.1password.com/get-serious-gdpr-compliance/\">To learn more about GDPR compliance, read the full blog</a>.</p>\n\n<div>\n<a target=\"_blank\" rel=\"noreferrer\" target=\"_blank\" rel=\"noreferrer\" title=\"Permanent link to ‘1Password: You Can’t Ignore GDPR Anymore’\" href=\"https://daringfireball.net/feeds/sponsors/2024/07/1password_you_cant_ignore_gdpr\"> ★ </a>\n</div>\n\n\t",
      "enclosureMime": null,
      "enclosureLink": null,
      "mediaThumbnail": null,
      "mediaDescription": null,
      "feedId": 57,
      "unread": true,
      "starred": false,
      "lastModified": 1721791092,
      "rtl": false,
      "fingerprint": "447a47fbc65945b845cb2737bb65be4c",
      "contentHash": "c26902a4a62e298acd3f08f8077931c0"
    }
*/

type ncItem struct {
	ID               int    `json:"id"`
	Guid             string `json:"guid"`
	GuidHash         string `json:"guidHash"`
	Url              string `json:"url"`
	Title            string `json:"title"`
	Author           string `json:"author"`
	PubDate          int    `json:"pubDate"`
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

type ncItemResponse struct {
	Items []ncItem `json:"items"`
}

type frItem struct {
	ID string `json:"id"`
}

type frItemsResponse struct {
	Items []frItem `json:"itemRefs"`
}

func getNCItems(count int, since int, ignoreRead bool) ([]ncItem, error) {
	client := http.Client{}

	urlStr := BaseUrl
	if since == 0 {
		urlStr += "/items?batchSize=" + strconv.Itoa(count)
		if ignoreRead {
			urlStr += "&getRead=false"
		}
	} else {
		urlStr += "/items/updated?lastModified=" + strconv.Itoa(since)
	}

	req, err := http.NewRequest("GET", urlStr, nil)
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

	var ncResponse ncItemResponse
	err = json.Unmarshal(resBody, &ncResponse)
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

// e.g. /api/greader.php/reader/api/0/stream/items/ids?n=1000&output=json&s=user/-/state/com.google/reading-list&xt=user/-/state/com.google/read
func GetStreamItemsIds(w http.ResponseWriter, r *http.Request) {
	outputFormat := r.URL.Query().Get("output")
	if outputFormat != "json" {
		notImplemented(w, r)
		return
	}

	// https://github.com/theoldreader/api?tab=readme-ov-file#items

	excludeTarget := r.URL.Query().Get("xt") // exclude target
	ignoreRead := excludeTarget == "user/-/state/com.google/read"

	// filterTarget := r.URL.Query().Get("it")  // include target (not used by NNN)

	count := 20
	if r.URL.Query().Get("n") != "" {
		n, err := strconv.Atoi(r.URL.Query().Get("n"))
		if err != nil {
			count = 20
		} else {
			count = n
		}
	}

	// order := r.URL.Query().Get("r") //  d|n|o, o ascending, n descending, d descending

	startTime := 0 // unix timestamp
	if r.URL.Query().Get("ot") != "" {
		ot, err := strconv.Atoi(r.URL.Query().Get("ot"))
		if err != nil {
			startTime = ot
		}
	}
	// This is not used by NNN
	// endTime := 0 // unix timestamp
	// if r.URL.Query().Get("et") != "" {
	// 	et, err := strconv.Atoi(r.URL.Query().Get("et"))
	// 	if err != nil {
	// 		endTime = et
	// 	}
	// }

	// continuationToken := r.URL.Query().Get("c") // used to get next page if exists

	streamIdInfos := r.URL.Query().Get("s") // this appears to only be reading-list or starred
	ncItems := []ncItem{}
	if streamIdInfos == "user/-/state/com.google/reading-list" {
		items, err := getNCItems(count, startTime, ignoreRead)
		ncItems = items
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Got NC Items", len(ncItems))

	} else if streamIdInfos == "user/-/state/com.google/starred" {

	}
	fmt.Println("Returning NC Items", len(ncItems))
	response := frItemsResponse{}

	for _, ncItem := range ncItems {
		fmt.Println("NC Item", ncItem.ID)
		response.Items = append(response.Items, frItem{
			ID: strconv.Itoa(ncItem.ID),
		})
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	println("Returning", string(jsonResponse))
	io.WriteString(w, string(jsonResponse))
}
