package controller

import (
	"apica-backend/internal/models"
	"apica-backend/internal/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var Cache *service.LRUCache

func init() {
	Cache = service.NewLRUCache(1024)
}

func GetCacheData(c echo.Context) error {
	key := c.Param("key")
	value, ok := Cache.Get(key)
	if !ok {
		return c.String(http.StatusNotFound, "Key not found in cache")
	}
	return c.JSON(http.StatusOK, value)
}

func SetCacheData(c echo.Context) error {
	reqBody := new(models.RequestBody)
	if err := c.Bind(reqBody); err != nil {
		return err
	}
	Cache.Set(reqBody.Key, reqBody.Value, time.Duration(reqBody.Expiration)*time.Second)
	return c.NoContent(http.StatusOK)
}
