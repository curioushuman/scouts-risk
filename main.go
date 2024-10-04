package main

import (
	"fmt"
	"scouts-risk/simple"
)

func main() {

	httpServer := simple.NewHttp()

	err := httpServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	// ! When is the store being closed?
}
