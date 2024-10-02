package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"scouts-risk/internal/store"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	store store.Service
}

/**
 * Creates a new HTTP server
 */
func NewHttp() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		store: store.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
