package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/platinumscatter/port-service/internal/config"
	"github.com/platinumscatter/port-service/internal/repository/inmem"
	"github.com/platinumscatter/port-service/internal/services"
	"github.com/platinumscatter/port-service/internal/transport"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	cfg := config.Read()

	portStoreRepo := inmem.NewPortStore()
	portService := services.NewPortService(portStoreRepo)
	httpServer := transport.NewHttpService(portService)

	router := mux.NewRouter()
	router.HandleFunc("/port", httpServer.GetPort).Methods("GET")
	router.HandleFunc("/count", httpServer.CountPorts).Methods("GET")
	router.HandleFunc("/ports", httpServer.UploadPorts).Methods("POST")

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting server HTTP on %s ", cfg.HTTPAddr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP Server ListenAndServe Error: %v", err)
	}
	<-stopped
	log.Printf("Have a nice day!")

	return nil
}
