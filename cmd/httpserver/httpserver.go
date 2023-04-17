package main

import (
	"productmanagement/config"
	"productmanagement/repository"
	"productmanagement/repository/gormrepo"
	"productmanagement/service"
	"productmanagement/transport"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// FIXME: Move logging from stdout/err to a file
func main() {
	// Config
	cfg := config.NewConfigImpl()

	// Repository
	repo := repository.NewRepository()
	db, err := gorm.Open(
		mysql.Open(cfg.GetString("DATABASE_URL")),
		&gorm.Config{},
	)
	if err != nil {
		panic("Failed to open gorm MySQL: " + err.Error())
	}
	repo.SetUser(gormrepo.NewUserImpl(db))

	// Service
	svc := service.NewService()
	svc.SetUser(service.NewUserImpl(cfg, repo))

	// Transport
	tp := transport.NewHTTPTransport(cfg, svc)
	e := echo.New()
	tp.Start(e)
}
