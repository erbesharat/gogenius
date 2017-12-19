package gogenius

import (
	"github.com/erbesharat/gogenius/resources"
)

// GetReferents returns referents by content item
// or user responsible for an included annotation
// @format: JSON
// https://docs.genius.com/#referents-h2
func GetReferents(webPageID string) string {
	return resources.GetJsonReferents(webPageID)
}

// GetSong returns data for a specific song.
// @format: JSON
// https://docs.genius.com/#songs-h2
func GetSong(songID string) string {
	return resources.GetJsonSong(songID)
}

// GetArtist returns data for a specific artist.
// @format: JSON
// https://docs.genius.com/#artists-h2
func GetArtist(artistID string) string {
	return resources.GetJsonArtist(artistID)
}

// GetArtistSongs returns documents (songs) for the artist specified.
// @format: JSON
// https://docs.genius.com/#artists-h2
func GetArtistSongs(artistID string) string {
	return resources.GetJsonArtistSongs(artistID)
}

// GetWebpage returns information about a web page retrieved by the page's full URL.
// @format: JSON
// https://docs.genius.com/#web_pages-h2
func GetWebpage(pageURL string) string {
	return resources.GetJsonWebPage(pageURL)
}

// Search documents hosted on Genius.
// @format: JSON
// https://docs.genius.com/#search-h2
func Search(query string) string {
	return resources.GetJsonSearch(query)
}

// GetAnnotation returns data for a specific annotation.
// @format: JSON
// https://docs.genius.com/#annotations-h2
func GetAnnotation(id string) string {
	return resources.GetJsonAnnotation(id)
}
