package main

import (
	"log"

	"github.com/longlnOff/prolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer("127.0.0.1:8080")
	log.Fatal(srv.ListenAndServe())
}
