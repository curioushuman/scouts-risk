package main

import (
	"fmt"
	"scouts-risk/internal/server"
)

func main() {

	httpServer := server.NewHttp()

	err := httpServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
