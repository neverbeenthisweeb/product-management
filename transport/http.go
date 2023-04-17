package transport

import (
	"net/http"
	"productmanagement/config"
	"productmanagement/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type httpTransport struct {
	cfg     config.Config
	svcUser service.User
}

func NewHTTPTransport(cfg config.Config, svc *service.Service) *httpTransport {
	return &httpTransport{
		cfg:     cfg,
		svcUser: svc.User,
	}
}

func (tr *httpTransport) GetUser(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return service.NewErrBadRequest("missing user ID")
	}

	user, err := tr.svcUser.Get(c.Request().Context(), userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (tr *httpTransport) Start(e *echo.Echo) {
	// Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	// Route
	e.GET("/users/:id", tr.GetUser)

	err := e.Start(tr.cfg.GetString("HTTP_SERVER_ADDRESS"))
	if err != nil {
		panic("Failed to start HTTP server: " + err.Error())
	}
}
