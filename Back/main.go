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
var main_task = task{}

// 1
func POSTHendler(h echo.Context) error {

	main_task = task{Tasks: Task}

	return h.JSON(http.StatusOK, "Все ок")
}

func GETHendler(h echo.Context) error {

	return h.JSON(http.StatusOK, main_task)
}

func main() {

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/myTask", POSTHendler)
	e.GET("/myTask", GETHendler)

	e.Start("localhost:8080")

}
