package main

import (
	"database_anonymizer/app/src/cache"
	"database_anonymizer/app/src/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cache := cache.NewCacheManager()
	cache.StartCacheManger("redis-anonymizer:6379", "", 0)

	handlers := handlers.NewHandler(cache)
	handlers.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
