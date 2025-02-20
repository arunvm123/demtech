package main

import (
	"log"
	"net/http"

	"github.com/arunvm123/demtech/email"
	"github.com/arunvm123/demtech/email/mockses"
)

type server struct {
	routes http.Handler
	email  email.Email
}

func newServer() *server {
	s := server{}
	return &s
}

func main() {
	s := newServer()

	s.routes = initialiseRoutes(s)

	s.email = mockses.New()

	log.Fatal(http.ListenAndServe(":9090", s.routes))

}
