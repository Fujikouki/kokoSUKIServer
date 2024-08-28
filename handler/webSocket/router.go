package webSocket

import (
	"awesomeProject1/usecase"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type handler struct {
	u usecase.WebSocketUseCase
}

func NewRouter(u usecase.WebSocketUseCase) http.Handler {

	r := chi.NewRouter()

	h := &handler{
		u: u,
	}
	r.Get("/chat", h.Open)
	return r

}
