package web

import (
	"net/http"
	"time"
)

var (
	server *http.Server

	location *time.Location
)

func init() {
	server = &http.Server{
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1024,
	}
}
