package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]any{
			"statusCode": 200,
			"message":    "Health Ok!",
		})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
