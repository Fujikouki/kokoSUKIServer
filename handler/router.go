package handler

import (
	"awesomeProject1/handler/account"
	"awesomeProject1/handler/aunth"
	"awesomeProject1/handler/webSocket"
	"awesomeProject1/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewRouter(wu usecase.WebSocketUseCase, chu usecase.ChatMessageU, acc usecase.AccountUsecase) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/", webSocket.NewRouter(wu, chu))
	r.Mount("/login", aunth.NewRouter())
	r.Mount("/account", account.NewRouter(acc))

	return r

}
