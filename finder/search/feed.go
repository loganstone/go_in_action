package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed .
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds .
func RetrieveFeeds() ([]*Feed, error) {
	f, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var feeds []*Feed
	err = json.NewDecoder(f).Decode(&feeds)
	return feeds, err
}
