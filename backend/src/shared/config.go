package shared

import (
	"os"
	"strconv"
)

func GetStringFromEnvDef(param string, def string) string {
	if retVal := os.Getenv(param); retVal != "" {
		return retVal
	}
	return def
}

func GetIntFromEnvDef(key string, def int) int {
	if val := os.Getenv(key); val != "" {
		r, err := strconv.Atoi(val)
		if err != nil {
			return def
		}
		return r
	}
	return def
}
