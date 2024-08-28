package handler

import (
	"awesomeProject1/handler/webSocket"
	"awesomeProject1/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewRouter(wu usecase.WebSocketUseCase) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/", webSocket.NewRouter(wu))

	return r

}
