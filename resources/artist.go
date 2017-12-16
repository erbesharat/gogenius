package resources

import "github.com/erbesharat/gogenius/handlers"

type ArtistResponse struct {
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
	Response struct {
		Artist struct {
			AlternateNames        []string `json:"alternate_names"`
			APIPath               string   `json:"api_path"`
			Description           string   `json:"description"`
			FacebookName          string   `json:"facebook_name"`
			FollowersCount        int      `json:"followers_count"`
			HeaderImageURL        string   `json:"header_image_url"`
			ID                    int      `json:"id"`
			ImageURL              string   `json:"image_url"`
			InstagramName         string   `json:"instagram_name"`
			IsMemeVerified        bool     `json:"is_meme_verified"`
			IsVerified            bool     `json:"is_verified"`
			Name                  string   `json:"name"`
			TwitterName           string   `json:"twitter_name"`
			URL                   string   `json:"url"`
			CurrentUserMetadata   string   `json:"current_user_metadata"`
			Iq                    int      `json:"iq"`
			DescriptionAnnotation string   `json:"description_annotation"`
			User                  string   `json:"user"`
		} `json:"artist"`
	} `json:"response"`
}

// func JsonToArtist(data []byte) *ArtistResponse {
// 	var r = new(ArtistResponse)
// 	err := json.NewDecoder(bytes.NewReader(data)).Decode(&r)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return r
// }

func GetJsonArtist(artistID string) string {
	url := handlers.GetResourceURI("artists/", artistID)
	return string(handlers.SendRequest(url))
}

func GetJsonArtistSongs(artistID string) string {
	url := handlers.GetArtistSongsURI(artistID)
	return string(handlers.SendRequest(url))
}
