package webSocket

import (
	"awesomeProject1/usecase"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type handler struct {
	u usecase.WebSocketUseCase
	c usecase.ChatMessageU
}

func NewRouter(u usecase.WebSocketUseCase, chu usecase.ChatMessageU) http.Handler {

	r := chi.NewRouter()

	h := &handler{
		u: u,
		c: chu,
	}
	r.Get("/chat", h.Open)
	return r

}
