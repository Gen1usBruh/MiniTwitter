package scope

import (
	"log/slog"

	"github.com/Gen1usBruh/MiniTwitter/internal/storage/postgres/sqlc"
)

type Dependencies struct {
	Sl *slog.Logger
	Db *postgresdb.Queries
	Secret string
}