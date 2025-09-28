package main

import "strconv"

func CapString(b []byte, length int) string {

	if len(b) <= length {
		return string(b)
	}
	return string(b[:length]) + "... (len=" + strconv.Itoa((len(b))) + ")"
}
