package main

import (
	"fmt"

	"scouts-risk/simple"
)

// ! UP TO HERE
// Working locally
// NOT WORKING in Railway
// * Build up from basics

func main() {
  httpServer := simple.NewHttp()

	err := httpServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

	// ! When is the store being closed?

}
