package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func run(addr string) error {
	mux := http.NewServeMux()

	db := NewSqliteDatabase()
	registerHandlers(mux, &db)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			// TODO: return error from routine
			log.Fatal("server failted")
		}
	}()

	notify, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-notify.Done()
	return startGracefullShutdown(srv)
}

func startGracefullShutdown(server *http.Server) error {
	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cleanupCancel()

	if err := server.Shutdown(cleanupCtx); err != nil {
		return fmt.Errorf("error shutting down server: %w", err)
	}

	return nil
}

func main() {
	addr := ":5000"
	if err := run(addr); err != nil {
		log.Fatalf("server stopped")
	}
}
