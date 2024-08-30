package aunth

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter() http.Handler {

	r := chi.NewRouter()

	r.Get("/set", setSession)
	r.Get("/get", getSession)

	return r

}
