package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger     log.Logger
	HTTPServer http.Server
}

// NewServer - func create server
func NewServer(logger log.Logger) *Server {
	// create http-router (ServeMux)
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	// create struct http.Server
	httpServer := http.Server{
		Addr:         ":8080",          // port 8080
		Handler:      router,           // http-router
		ErrorLog:     &logger,          // logger
		ReadTimeout:  5 * time.Second,  // time-out for reading (5 seconds)
		WriteTimeout: 10 * time.Second, // time-out for recording (10 seconds)
		IdleTimeout:  15 * time.Second, // time-out waiting (15 seconds)
	}

	// return server
	return &Server{
		Logger:     logger,
		HTTPServer: httpServer,
	}
}
