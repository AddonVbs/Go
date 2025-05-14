package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type task struct {
	Tasks string `json:"task"`
}

var Task string = "Содить за молоком"

// 1
func PostHendler(h echo.Context) error {

	return h.JSON(http.StatusOK, map[string]string{"task": Task})

}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/myTask", PostHendler)

	e.Start("localhost:8080")

}
