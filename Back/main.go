package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type task struct {
	Tasks string `json:"task"`
}

var main_task = task{}

// 1

func POSTHendler(h echo.Context) error {
	var t task

	if err := h.Bind(&t); err != nil {
		return h.JSON(http.StatusBadRequest, "Неверный формат запроса")

	}

	main_task = t

	return h.JSON(http.StatusOK, "Все ок, ваша задача сохранена!")
}

func GETtask(h echo.Context) error {

	return h.JSON(http.StatusOK, main_task)
}

func main() {

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/task", POSTHendler)
	e.GET("/task", GETtask)

	e.Start("localhost:8080")

}
