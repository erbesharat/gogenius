package gogenius

import (
	"io/ioutil"
	"net/http"

	"github.com/erbesharat/gogenius/handlers"
)

// Referents Returns a list of all refrents from given page
func Referents(webPageID string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", handlers.GetReferentsURI(webPageID), nil)
	req.Header.Add("Authorization", handlers.GetAccessToken())
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	panic(string(data))
	return string(data)
}
