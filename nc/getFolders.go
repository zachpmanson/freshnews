package nc

import (
	"encoding/json"
	"fmt"
	"freshnews/utils"
)

type ncFolderReponse struct {
	Folders []folder `json:"folders"`
}
type folder struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Opened bool   `json:"opened"`
	Feeds  []feed `json:"feeds"`
}

type feed struct {
	// Define the fields for the feeds if there are any
	// Leaving it empty since the provided JSON shows an empty array
}

func GetFolders() ([]folder, error) {

	resBody, err := utils.NcGetReq("/folders")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var ncResponse ncFolderReponse
	err = json.Unmarshal(resBody, &ncResponse)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return ncResponse.Folders, nil
}
