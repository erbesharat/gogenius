package resources

import "github.com/erbesharat/gogenius/handlers"

func GetJsonSearch(query string) string {
	url := handlers.GetResourceURI("search/?q=", query)
	return string(handlers.SendRequest(url))
}
