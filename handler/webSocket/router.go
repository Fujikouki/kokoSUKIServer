package webSocket

import (
	"awesomeProject1/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"net/http"
)

type handler struct {
	u  usecase.WebSocketUseCase
	c  usecase.ChatMessageU
	se *sessions.CookieStore
}

func NewRouter(u usecase.WebSocketUseCase, chu usecase.ChatMessageU, se *sessions.CookieStore) http.Handler {

	r := chi.NewRouter()

	h := &handler{
		u:  u,
		c:  chu,
		se: se,
	}
	r.Get("/chat", h.Open)
	return r

}
