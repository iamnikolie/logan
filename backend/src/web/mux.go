package web

import "github.com/go-chi/chi"

func mux() *chi.Mux {
	r := chi.NewMux()

	r.Get("/{level}", getLogs)

	return r
}
