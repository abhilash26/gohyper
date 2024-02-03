package server

import (
	"github.com/abhilash26/gohyper/internal/option"

	"fmt"
	"log"
	"net/http"
	"time"
)

// ServerOptions holds the configuration for the HTTP server.
type ServerOptions struct {
	Port              int
	IdleTimeout       time.Duration
	WriteTimeout      time.Duration
	ReadTimeout       time.Duration
	MaxHeaderBytes    int
	ReadHeaderTimeout time.Duration
}

func loadServerOptions() ServerOptions {
	return ServerOptions{
		Port:           option.GetIntEnv("SERVER_PORT", 8080),
		IdleTimeout:    option.GetDurationEnv("SERVER_IDLE_TIMEOUT", "60s"),
		WriteTimeout:   option.GetDurationEnv("SERVER_WRITE_TIMEOUT", "15s"),
		ReadTimeout:    option.GetDurationEnv("SERVER_READ_TIMEOUT", "15s"),
		MaxHeaderBytes: option.GetIntEnv("SERVER_MAX_HEADER_BYTES", 1*1024*1024),
	}
}

func Start(h http.Handler) {

	opts := loadServerOptions()

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", opts.Port),
		Handler:           h,
		IdleTimeout:       opts.IdleTimeout,
		WriteTimeout:      opts.WriteTimeout,
		ReadTimeout:       opts.ReadTimeout,
		MaxHeaderBytes:    opts.MaxHeaderBytes,
		ReadHeaderTimeout: opts.ReadHeaderTimeout,
	}

	log.Printf("Server is listening on :%d", opts.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
