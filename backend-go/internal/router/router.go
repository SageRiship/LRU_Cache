package router

import (
	"apica-backend/internal/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())
	e.GET("/cache/:key", controller.GetCacheData)
	e.POST("/cache", controller.SetCacheData)

	return e
}
