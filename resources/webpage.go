package resources

import (
	"bytes"
	"encoding/json"

	"github.com/erbesharat/gogenius/handlers"
)

type WebPageResponse struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Response struct {
		WebPage struct {
			APIPath         string `json:"api_path"`
			Domain          string `json:"domain"`
			ID              int    `json:"id"`
			NormalizedURL   string `json:"normalized_url"`
			ShareURL        string `json:"share_url"`
			Title           string `json:"title"`
			URL             string `json:"url"`
			AnnotationCount int    `json:"annotation_count"`
		} `json:"web_page"`
	} `json:"response"`
}

// This method works correctly, other resources methods should work like this.
func JsonToWebPage(data []byte) *WebPageResponse {
	var r = new(WebPageResponse)
	err := json.NewDecoder(bytes.NewReader(data)).Decode(&r)
	if err != nil {
		panic(err)
	}
	return r
}

func GetJsonWebPage(pageURL string) string {
	url := handlers.GetWebpageURI(pageURL)
	return string(handlers.SendRequest(url))
}
