package server

import (
	"awesomeProject1/handler"
	"awesomeProject1/usecase"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func Run() error {

	addr := ":" + strconv.Itoa(8080)

	log.Printf("Server started on http://%s\n", addr)

	webSocketUsecawse := usecase.NewWebSocketUseCase()

	r := handler.NewRouter(webSocketUsecawse)

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		if err := srv.Serve(l); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	<-ctx.Done()
	ctx, _ = context.WithTimeout(context.Background(), time.Second*5)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	return nil
}
