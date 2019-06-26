package services

import (
	"fmt"

	"github.com/google/uuid"
)

func PrefixedUUID(prefix string) string {
	var cleanedPrefix string

	if len(prefix) > 2 {
		fmt.Println("Prefix is Too Long.  It's been Trimmed to 2 Characters")
		cleanedPrefix = prefix[0:2] + "_"
	} else if len(prefix) == 0 {
		fmt.Println("Prefix not provided.  Just so you know")
	} else {
		cleanedPrefix = prefix + "_"
	}

	id, _ := uuid.NewRandom()
	return fmt.Sprintf("%s%s", cleanedPrefix, id)
}
