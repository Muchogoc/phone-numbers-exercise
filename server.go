package main

import (
	"log"

	"github.com/Muchogoc/phone-numbers-exercise/service/presentation"
)

var (
	port = 8080
)

func main() {
	srv := presentation.PrepareServer(port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server")
	}

}
