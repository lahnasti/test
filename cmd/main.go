package main

import (
	"github.com/lahnasti/test/internal/server"
)

func main() {
	srv := server.NewServer()
	srv.Run(":8080")
}
