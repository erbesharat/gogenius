package gogenius

import (
	"io/ioutil"
	"net/http"

	"github.com/erbesharat/gogenius/handlers"
)

// Referents Returns a list of all refrents from given page
func Referents(webPageID string) string {
	client := &http.Client{}
	url := handlers.GetResourceURI("referents?web_page_id=", webPageID)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", handlers.GetAuthorizationToken())
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
	return string(data)
}

func GetSong(songID string) string {
	client := &http.Client{}
	url := handlers.GetResourceURI("songs/", songID)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", handlers.GetAuthorizationToken())
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
	return string(data)
}

func GetArtist(artistID string) string {
	client := &http.Client{}
	url := handlers.GetResourceURI("songs/", artistID)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", handlers.GetAuthorizationToken())
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
	return string(data)
}

func GetArtistSongs(artistID string) string {
	client := &http.Client{}
	url := handlers.GetArtistSongsURI(artistID)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", handlers.GetAuthorizationToken())
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
	return string(data)
}

func GetWebpage(pageURL string) string {
	client := &http.Client{}
	url := handlers.GetWebpageURI(pageURL)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", handlers.GetAuthorizationToken())
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
	return string(data)
}

func Search(query string) string {
	client := &http.Client{}
	url := handlers.GetResourceURI("search/?q=", query)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", handlers.GetAuthorizationToken())
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
	return string(data)
}

func GetCurrentAccount() string {
	client := &http.Client{}
	url := handlers.GetResourceURI("account?access_token=", handlers.GetAccessToken())
	req, err := http.NewRequest("GET", url, nil)
	panic(url)
	// req.Header.Add("Authorization", handlers.GetAuthorizationToken())
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
	return string(data)
}

// ?raw_annotatable_url=https%3A%2F%2Fdocs.genius.com
