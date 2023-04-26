package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	handlers "jawncorp.com/bank/user/handler"
	"os"
)

func Router() *echo.Echo {
	e := echo.New()

	logger := middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"id":"${id}","time":"${time_rfc3339}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}"}` + "\n",
		Output: os.Stdout,
	})

	e.Use(logger)
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.INFO)
	e.HideBanner = true
	e.HidePort = false

	healthCheckHandler := handlers.NewHealthCheckHandler()
	userHandler := handlers.NewUserHandler(NewSQLHandler())

	e.GET("/healthcheck", healthCheckHandler.HealthCheck())

	e.GET("/users", userHandler.GetUsers())
	e.GET("/user/:id", userHandler.GetUser())
	e.POST("/user", userHandler.CreateUser())
	e.POST("/user/leave", userHandler.UpdateLeaveAttr())

	return e
}
