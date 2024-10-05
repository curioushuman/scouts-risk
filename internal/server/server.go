package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"scouts-risk/internal/store"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port string
	addr string

	store store.Service
}

/**
 * Creates a new HTTP server
 */
func NewHttp() *http.Server {
	port := httpPort()
	NewServer := &Server{
		port: port,
		addr: fmt.Sprintf(":%s", port),
		// store: store.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         NewServer.addr,
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func httpPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return
}