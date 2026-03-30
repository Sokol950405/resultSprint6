package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// create logger
	logger := log.New(os.Stdout, "[MAIN] ", log.Ldate|log.Ltime|log.Lshortfile)

	// create server
	srv := server.NewServer(*logger)

	// start server
	logger.Println("Starting server on :8080")
	if err := srv.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatal("Server error:", err)
	}
}
