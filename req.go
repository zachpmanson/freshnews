package main

import (
	"fmt"
	"io"
	"net/http"
)

func NcGetReq(path string) ([]byte, error) {

	fullUrl := BaseUrl + path
	client := http.Client{}
	fmt.Println("--> GET", path)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// fmt.Println("Auth", "Basic", Credentials)
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

	fmt.Println("<--", CapString(resBody, 500))

	return resBody, nil
}
