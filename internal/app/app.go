package app

import (
	"net/http"

	"github.com/Gen1usBruh/MiniTwitter/internal/config"
	"github.com/Gen1usBruh/MiniTwitter/internal/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	*http.Server
}

// NewApp takes config params from environments and returns server instance ready to start listen.
// @title           MiniTwitter Swagger API
// @version         1.0
// @description     This is a MiniTwitter service API.
// @BasePath  /api/v1
func NewApp(server config.Server, restServer http.Handler) (*App, error) {
	sm := http.NewServeMux()

	sm.HandleFunc("/health-check/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	sm.Handle(
		"/api/v1/",
		http.StripPrefix("/api/v1", middleware.CorsMiddleware(restServer)),
	)
	sm.Handle("/swagger/", httpSwagger.Handler(httpSwagger.URL("/api/v1/swagger/doc.json")))

	httpServer := http.Server{
		Addr:         server.Address,
		Handler:      sm,
		ReadTimeout:  server.Timeout,
		WriteTimeout: server.Timeout,
		IdleTimeout:  server.IdleTimeout,
	}

	return &App{
		Server: &httpServer,
	}, nil
}
