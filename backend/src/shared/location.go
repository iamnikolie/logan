package shared

import (
	"time"
)

var (
	location *time.Location
)

func Location() *time.Location {
	if location != nil {
		return location
	}
	var err error
	location, err = time.LoadLocation(GetStringFromEnvDef("LOCATION", "Europe/Kiev"))
	if err != nil {
		def, _ := time.LoadLocation("UTC")
		return def
	}
	return location
}
