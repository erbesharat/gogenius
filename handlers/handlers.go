package handlers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const baseURI = "https://api.genius.com/"

func GetReferentsURI(key string) string {
	return fmt.Sprintf("%s%s%s", baseURI, "referents?web_page_id=", key)
}

func GetAccessToken() string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s %s", "Bearer", os.Getenv("ACCESS_TOKEN"))
}

// referents?web_page_id=10347
