package resources

import (
	"encoding/json"

	"github.com/erbesharat/gogenius/handlers"
)

func JsonToSong(data []byte) interface{} {
	var datastruct map[string]interface{}
	_ = json.Unmarshal([]byte(data), &datastruct)
	return datastruct["response"]
}

func GetJsonSong(songID string) string {
	url := handlers.GetResourceURI("songs/", songID)
	return string(handlers.SendRequest(url))
}
