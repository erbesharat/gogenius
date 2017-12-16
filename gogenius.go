package gogenius

import (
	"github.com/erbesharat/gogenius/resources"
)

// Referents Returns a list of all refrents from given page
func GetReferents(webPageID string) string {
	return resources.GetJsonReferents(webPageID)
}

func GetSong(songID string) string {
	return resources.GetJsonSong(songID)
}

func GetArtist(artistID string) string {
	return resources.GetJsonArtist(artistID)
}

func GetArtistSongs(artistID string) string {
	return resources.GetJsonArtistSongs(artistID)
}

func GetWebpage(pageURL string) string {
	return resources.GetJsonWebPage(pageURL)
}

func Search(query string) string {
	return resources.GetJsonSearch(query)
}

func GetAnnotation(id string) string {
	return resources.GetJsonAnnotation(id)
}
