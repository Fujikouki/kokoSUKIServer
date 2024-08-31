package account

import (
	"awesomeProject1/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"net/http"
)

type handler struct {
	accountUseCase usecase.AccountUsecase
	cs             *sessions.CookieStore
}

func NewRouter(u usecase.AccountUsecase, se *sessions.CookieStore) http.Handler {
	r := chi.NewRouter()

	h := &handler{
		accountUseCase: u,
		cs:             se,
	}

	r.Post("/create", h.CreateAccount)

	r.Post("/login", h.Login)

	return r
}
