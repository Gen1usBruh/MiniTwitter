package main

import (
	"log"

	"github.com/Gen1usBruh/MiniTwitter/internal/app"
	"github.com/Gen1usBruh/MiniTwitter/internal/config"
	"github.com/Gen1usBruh/MiniTwitter/internal/logger/sl"
	"github.com/Gen1usBruh/MiniTwitter/internal/rest"
	"github.com/Gen1usBruh/MiniTwitter/internal/scope"
	"github.com/Gen1usBruh/MiniTwitter/internal/storage/postgres"
	postgresdb "github.com/Gen1usBruh/MiniTwitter/internal/storage/postgres/sqlc"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("Could not create config: %v\n", err)
	}
	conn, err := postgres.ConnectDB(&conf.Database)
	if err != nil {
		log.Fatalf("Could not connect to postgres: %v\n", err)
	}
	restServer := rest.NewHandler(rest.HandlerConfig{
		Dep: &scope.Dependencies{
			Sl:     sl.SetupLogger(&conf.Logger),
			Db:     postgresdb.New(conn),
			Secret: "test secret",
		},
	})

	server, err := app.NewApp(conf.Server, restServer)
	if err != nil {
		log.Fatalf("Unable to start server: %v\n", err)
	}

	log.Printf("Starting server...\nAddress: %v", server.Server.Addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
