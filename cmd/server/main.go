package main

import (
	"log"

	"github.com/Shishir-rmv/logmatix/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
