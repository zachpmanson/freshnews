package handlers

import (
	"encoding/json"
	"freshnews/nc"
	"freshnews/utils"
	"net/http"
)

type FrTag struct {
	ID   string `json:"id"`
	Type string `json:"type,omitempty"`
}

type FrTagsResponse struct {
	Tags []FrTag `json:"tags"`
}

func GetTagList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("output") != "json" {
		NotImplemented(w, r)
		return
	}

	folders, err := nc.GetFolders()
	if err != nil {
		utils.WriteError(w, err, 500)
		return
	}
	response := FrTagsResponse{
		Tags: []FrTag{{ID: "user/-/state/com.google/starred"}},
	}
	for _, folder := range folders {
		response.Tags = append(response.Tags, FrTag{
			ID: "user/-/label/" + folder.Name, Type: "folder",
		})

	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteBytes(w, jsonResponse, true)
}
