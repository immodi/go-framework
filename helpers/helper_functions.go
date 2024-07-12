package helpers

import (
	"os"
)

func CheckForStaticFiles() bool {
	if _, err := os.Stat("static"); os.IsNotExist(err) {
		return false
	}

	return true
}
