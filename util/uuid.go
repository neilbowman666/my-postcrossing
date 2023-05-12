package util

import (
	"github.com/gofrs/uuid"
	"log"
)

func GetUUID() (string, error) {
	u2, err := uuid.NewV4()
	if err != nil {
		log.Printf("Something went wrong: %s", err)
		return "", err
	}
	return u2.String(), nil
}
