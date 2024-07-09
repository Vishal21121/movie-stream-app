package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lpernett/godotenv"
	"github.com/vishal21121/movie-stream-app/internals/handlers"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}",
	}))

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]any{
			"statusCode": 200,
			"message":    "Health Ok!",
		})
	})
	e.GET("/", handlers.GetMovies)

	e.Logger.Fatal(e.Start(":8080"))
}
