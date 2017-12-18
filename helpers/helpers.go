package helpers

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
