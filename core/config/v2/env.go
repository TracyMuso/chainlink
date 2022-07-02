package v2

import (
	"os"
	"strings"
)

var (
	Dev = strings.ToLower(os.Getenv("CL_DEV")) == "true"
)
