package postgresdb

import (
	"log"
	"os"
	"testing"

	"github.com/Gen1usBruh/MiniTwitter/internal/config"
	"github.com/Gen1usBruh/MiniTwitter/internal/storage/postgres"
	"github.com/Gen1usBruh/MiniTwitter/internal/logger/sl"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("Could not create config: %v\n", err)
	}

	log := sl.SetupLogger(&conf.Logger)

	conn, err := postgres.ConnectDB(&conf.Database)
	if err != nil {
		log.Error("Could not connect to postgres", sl.Err(err))
		os.Exit(1)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}