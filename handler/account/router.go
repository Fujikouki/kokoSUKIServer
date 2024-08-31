package account

import (
	"awesomeProject1/usecase"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type handler struct {
	accountUsecase usecase.AccountUsecase
}

func NewRouter(u usecase.AccountUsecase) http.Handler {
	r := chi.NewRouter()

	h := &handler{
		accountUsecase: u,
	}

	r.Post("/", h.CreateAccount)

	return r
}
