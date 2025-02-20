package main

import (
	"log"
	"net/http"
)

type server struct {
	routes http.Handler
}

func newServer() *server {
	s := server{}
	return &s
}

func main() {
	s := newServer()

	s.routes = initialiseRoutes(s)

	log.Fatal(http.ListenAndServe(":9090", s.routes))

}
