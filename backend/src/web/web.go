package web

import (
	log "logan"

	"../shared"
)

func Start() {
	server.Handler = mux()
	server.Addr = shared.GetStringFromEnvDef("LOGAN_DEFAULT_HOST", ":9999")
	log.Printf("app run at port: %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("cannot start web server: %s", err)
	}
}

func Stop() {
}
