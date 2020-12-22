package main

// @host localhost:3000
// @BasePath /

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/starptech/fay/docs"
	"github.com/starptech/fay/internals/server"
)

func main() {
	s := server.New()

	// Start server
	go func() {
		if err := s.Server.Start(":3000"); err != nil {
			s.Server.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("could not shutdown %s", err)
	}
}
