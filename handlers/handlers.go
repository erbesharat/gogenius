package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/erbesharat/gogenius/helpers"
	"github.com/joho/godotenv"
)

const baseURI = "https://api.genius.com/"

func GetResourceURI(url string, key string) string {
	return fmt.Sprintf("%s%s%s", baseURI, url, key)
}

func GetArtistSongsURI(artistID string) string {
	return fmt.Sprintf("%s%s", GetResourceURI("artists/", artistID), "/songs")
}

func GetWebpageURI(url string) string {
	return fmt.Sprintf("%s%s%s", baseURI, "web_pages/lookup?raw_annotatable_url=", url)
}

func GetAccessToken() string {
	err := godotenv.Load()
	helpers.CheckErr(err)
	return os.Getenv("ACCESS_TOKEN")
}

func SendRequest(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", GetAuthorizationToken())
	helpers.CheckErr(err)
	resp, err := client.Do(req)
	helpers.CheckErr(err)
	data, err := ioutil.ReadAll(resp.Body)
	helpers.CheckErr(err)

	return data
}

func GetAuthorizationToken() string {
	err := godotenv.Load()
	helpers.CheckErr(err)
	return fmt.Sprintf("%s %s", "Bearer", os.Getenv("ACCESS_TOKEN"))
}

// referents?web_page_id=10347
