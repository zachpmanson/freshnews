package utils

import (
	"net/url"
	"strconv"
)

func CapString(b []byte, length int) string {

	if len(b) <= length {
		return string(b)
	}
	return string(b[:length]) + "... (len=" + strconv.Itoa((len(b))) + ")"
}

func ToQueryParams(pairs [][2]string) string {
	values := url.Values{}
	for _, p := range pairs {
		values.Add(p[0], p[1])
	}
	return values.Encode()
}
