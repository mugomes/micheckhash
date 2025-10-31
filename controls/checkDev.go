package controls

import (
	"os"
	"strconv"
)

func IsDevMode() bool {
	b, _ := strconv.ParseBool(os.Getenv("MIDEV"))
	return b
}