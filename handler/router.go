package handler

import (
	"awesomeProject1/handler/account"
	"awesomeProject1/handler/webSocket"
	"awesomeProject1/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"net/http"
)

func NewRouter(wu usecase.WebSocketUseCase, chu usecase.ChatMessageU, acc usecase.AccountUsecase, se *sessions.CookieStore) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/we", webSocket.NewRouter(wu, chu, se))
	r.Mount("/account", account.NewRouter(acc, se))

	return r

}
