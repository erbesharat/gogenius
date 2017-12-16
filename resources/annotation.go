package resources

import "github.com/erbesharat/gogenius/handlers"

func GetJsonAnnotation(id string) string {
	url := handlers.GetResourceURI("annotations/", id)
	return string(handlers.SendRequest(url))
}
