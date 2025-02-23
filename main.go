package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/arunvm123/demtech/config"
	"github.com/arunvm123/demtech/email"
	"github.com/arunvm123/demtech/email/mockses"
	"github.com/arunvm123/demtech/model"
	"github.com/arunvm123/demtech/model/postgres"
)

type server struct {
	routes http.Handler
	email  email.Email
	db     model.DB
}

func newServer() *server {
	s := server{}
	return &s
}

// @title           Demtech Mock SES
// @version         1.0
// @description     APIs available ofr fairsplits.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:9090
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	s := newServer()

	// Flags to read configuration
	filePath := flag.String("config-path", "config.yaml", "Filepath to configuration file")
	env := flag.Bool("config-env", false, "If set to true, Configuration is read from environment variable")
	flag.Parse()

	// Reading config variables
	config, err := config.Initialise(*filePath, *env)
	if err != nil {
		log.Fatalf("Failed to read config\n%v", err)
	}
	log.Printf("%+v", config)

	s.routes = initialiseRoutes(s)

	s.email = mockses.New()

	// "host=localhost user=gorm password=gorm dbname=gorm port=9920"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Database.Host, config.Database.User, config.Database.Password,
		config.Database.DatabaseName, config.Database.Port)
	pg, err := postgres.New(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database\n%v", err)
	}

	pg.MigrateDB()

	s.db = pg

	log.Fatal(http.ListenAndServe(":"+config.Port, s.routes))

}
