package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lahnasti/test/internal/server"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loadin .env file")
	}

	srv := server.NewServer()
	srv.Run(":8080")
}
