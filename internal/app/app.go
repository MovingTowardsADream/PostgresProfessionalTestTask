package app

import (
	"PostgresProfessionalTestTask/configs"
	v1 "PostgresProfessionalTestTask/internal/controller/http/v1"
	"PostgresProfessionalTestTask/internal/repository"
	"PostgresProfessionalTestTask/internal/service"
	"PostgresProfessionalTestTask/pkg/httpserver"
	"PostgresProfessionalTestTask/pkg/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log/slog"
)

type App struct {
	HTTPServer *httpserver.Server
	DB         *postgres.Postgres
}

func New(log *slog.Logger, cfg *configs.Config) *App {
	// Connect postgres db
	pg, err := postgres.NewPostgresDB(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		panic("app - Run - postgres.NewPostgresDB: " + err.Error())
	}
	defer pg.Close()

	repos := repository.NewRepositories(pg)
	services := service.NewService(repos)
	handler := gin.New()

	v1.NewRouter(handler, log, services)
	httpServer := httpserver.New(log, handler, httpserver.Port(cfg.HTTP.Port), httpserver.WriteTimeout(cfg.HTTP.Timeout))

	return &App{
		HTTPServer: httpServer,
		DB:         pg,
	}
}
