# Freshnews

A bridge between Nextcloud News API calls and the FreshRSS API, so that I can Nextcloud News as the backend for NetNewsWire.

To use:

1. Start up with `freshnews --credentials $BASE64nextclouduser:passwd` 
2. Set the FreshRSS instance to `http://localhost:8080/api/greader.php/`, 


## Docs

- **nextcloud news api -> [freshrss greader api](https://freshrss.github.io/FreshRSS/en/developers/06_GoogleReader_API.html) bridge**, allow netnewswire support for nextcloud news backend
	- [freshrss greader api doc](https://freshrss.github.io/FreshRSS/en/developers/06_GoogleReader_API.html)
	- [freshrss greader api implementation](https://github.com/FreshRSS/FreshRSS/blob/edge/p/api/greader.php#L184) (use this i think)