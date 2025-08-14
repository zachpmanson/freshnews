package main

import (
	"io"
	"net/http"
)
func GetStreamItemContents(w http.ResponseWriter, r *http.Request) {
	// only post
	if r.Method != "POST" {
		notImplemented(w, r)
		return
	}

	// check for form data
	r.ParseForm()

	// get form data
	// itemIds := r.Form.Get("i")

	io.WriteString(w, `
		{
			"id": "user/-/state/com.google/reading-list",
			"updated": 1716875047,
			"items": [
			{
				"id": "tag:google.com,2005:reader/item/0006197cb0d66926",
				"crawlTimeMsec": "1716873193744",
				"timestampUsec": "1716873193744678",
				"published": 1716240780,
				"title": "iOS 17.5.1 Includes Fix for Bug That Resurfaced Deleted Photos",
				"canonical": [
				{
					"href": "https://www.macrumors.com/2024/05/20/apple-releases-ios-17-5-1-photos-bug/"
				}
				],
				"alternate": [
				{
					"href": "https://www.macrumors.com/2024/05/20/apple-releases-ios-17-5-1-photos-bug/"
				}
				],
				"categories": ["user/-/state/com.google/reading-list", "user/-/label/m"],
				"origin": {
				"streamId": "feed/4",
				"htmlUrl": "https://daringfireball.net/",
				"title": "Daring Fireball"
				},
				"summary": {
				"content": "<p>MacRumors, quoting Apple’s own <a href=\"https://support.apple.com/en-us/118723#175\">release notes</a>:</p>\n\n<blockquote>\n  <p>This update provides important bug fixes and addresses a rare\nissue where photos that experienced database corruption could\nreappear in the Photos library even if they were deleted.</p>\n</blockquote>\n\n<p>That’s a nasty bug, so it’s no surprise that 17.5.1 is here just one week after 17.5.0.</p>\n\n<p>Last week MacRumors <a href=\"https://www.macrumors.com/2024/05/17/ios-17-5-bug-wiped-devices-photos-resurfacing/\">also reported on a claim</a> that iOS 17.5 was resurfacing photos on devices that had been wiped and resold (or given away), but that was an extraordinary claim that didn’t jibe with our understanding of how “wiping” an iOS device works. All storage on iOS devices is encrypted, and when you wipe the device (Settings → General → Transfer or Reset iPhone/iPad → Erase All Content and Settings), the encryption key is destroyed. The system doesn’t, and doesn’t need to, overwrite the storage with 0’s or random bits. It just destroys the encryption key from the Secure Enclave, rendering the data already written to storage unrecoverable. That report was based on a single post on Reddit, which has since been deleted. (MacRumors has an update appended to <a href=\"https://www.macrumors.com/2024/05/17/ios-17-5-bug-wiped-devices-photos-resurfacing/\">that report</a>, but I think they should move that update to the top of the post, not the bottom. All evidence suggests that it was a false alarm.)</p>\n\n<div>\n<a title=\"Permanent link to ‘iOS 17.5.1 Includes Fix for Bug That Resurfaced Deleted Photos’\" href=\"https://daringfireball.net/linked/2024/05/20/ios-17-deleted-photos-bug\"> ★ </a>\n</div>"
				},
				"author": "John Gruber"
			},
			{
				"id": "tag:google.com,2005:reader/item/0006197cb0d6691d",
				"crawlTimeMsec": "1716873193744",
				"timestampUsec": "1716873193744669",
				"published": 1715892600,
				"title": "Samsung Pepsis Its Pants Again",
				"canonical": [
				{
					"href": "https://twitter.com/SamsungMobileUS/status/1790824457365594487"
				}
				],
				"alternate": [
				{
					"href": "https://twitter.com/SamsungMobileUS/status/1790824457365594487"
				}
				],
				"categories": ["user/-/state/com.google/reading-list", "user/-/label/m"],
				"origin": {
				"streamId": "feed/4",
				"htmlUrl": "https://daringfireball.net/",
				"title": "Daring Fireball"
				},
				"summary": {
				"content": "<p><a href=\"https://daringfireball.net/linked/2024/05/16/new-ipad-pros-bend-resistant\">Speaking of</a> Apple’s <a href=\"https://daringfireball.net/linked/2024/05/09/dhh-crush\">“Crush” ad</a>, Samsung has posted a “response”, depicting a woman guitarist sitting atop a paint-splash-strewn platform standing in for a hydraulic press, with the slogan “We would never crush creativity. #UnCrush”</p>\n\n<p>Rather than sit back and enjoy Apple own-goaling itself last week, they couldn’t resist gracelessly piling on, accomplishing nothing but to remind everyone that they’re Pepsi to Apple’s Coke — content to sit in second place forever, copying not just Apple’s hardware and software designs, but even parodying Apple’s ads. This one is the equivalent of picking ideas out of Apple’s trash. Sad.</p>\n\n<p><strong>Update:</strong> This marketing strategy didn’t turn out well <a href=\"https://mastodon.social/@limi/112451347727656529\">for Commodore</a>.</p>\n\n<div>\n<a title=\"Permanent link to ‘Samsung Pepsis Its Pants Again’\" href=\"https://daringfireball.net/linked/2024/05/16/samsung-uncrush\"> ★ </a>\n</div>"
				},
				"author": "John Gruber"
			}
			]
		}
	`)
}