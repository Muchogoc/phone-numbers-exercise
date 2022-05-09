package main

import (
	"log"

	"github.com/Muchogoc/phone-numbers-exercise/service/presentation"
)

func main() {
	port := 8080
	srv := presentation.PrepareServer(port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server")
	}

}
