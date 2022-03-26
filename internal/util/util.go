package util

import (
	"log"
)

func AssertError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
