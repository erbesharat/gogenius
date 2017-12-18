package resources

import (
	"bytes"
	"encoding/json"

	"github.com/erbesharat/gogenius/handlers"
	"github.com/erbesharat/gogenius/helpers"
)

type Referent struct {
	Type                 string `json:"_type"`
	AnnonatorID          int    `json:"annotator_id"`
	AnnotatorLogin       string `json:"annotator_login"`
	APIPath              string `json:"api_path"`
	Classification       bool   `json:"classification"`
	Featured             string `json:"featured"`
	Fragment             string `json:"fragment"`
	ID                   int    `json:"id"`
	IsDescription        bool   `json:"is_description"`
	Path                 string `json:"path"`
	Range                string `json:"range"`
	SongID               int    `json:"song_id"`
	URL                  string `json:"url"`
	VerifiedAnnotatorIDs string `json:"verified_annotator_ids"`
	Annotatable          string `json:"annotatable"`
	Annotations          string `json:"annotations"`
}

func JsonToRefrent(data []byte) *Referent {
	var r = new(Referent)
	err := json.NewDecoder(bytes.NewReader(data)).Decode(&r)
	helpers.CheckErr(err)
	return r
}

func GetJsonReferents(webPageID string) string {
	url := handlers.GetResourceURI("referents?web_page_id=", webPageID)
	return string(handlers.SendRequest(url))
}
