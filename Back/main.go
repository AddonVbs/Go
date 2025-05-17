package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID    int    `json:"id"`
	Task1 string `json:"task"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// all global var


func POSTHendler(h echo.Context) error {
	var t task = task{Tasks: "Помыть посуду"}
	/*
		if err := h.Bind(&t); err != nil {
			return h.JSON(http.StatusBadRequest, "Неверный формат запроса")

		} */

var IDtask = 1
var my_task []Task

func GetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, my_task)
}

func PostHandler(c echo.Context) error {
	var t Task

	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add Task",
		})
	}


	t.ID = IDtask
	IDtask++

	my_task = append(my_task, t)

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Task was added successfully",
	})
}

func PatchHandler(c echo.Context) error {
	IDparam := c.Param("id")

	id, err := strconv.Atoi(IDparam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: "Bad", Message: "Invalid Patch !"})
	}
	var updataTask Task

	if err := c.Bind(&updataTask); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not updata Task",
		})
	}

	updata := false
	for i, task := range my_task {
		if task.ID == id {
			updataTask.ID = id
			my_task[i] = updataTask
			updata = true
			break

		}

	}
	if !updata {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error-updata (77 строка) ",
			Message: "Could not updata Task",
		})

	}
	return c.JSON(http.StatusOK, Response{Status: "Success", Message: "Was update "})

}

func main() {
	e := echo.New()


	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", GetHandler)
	e.POST("/task", PostHandler)
	e.PATCH("/task/:id", PatchHandler)

	e.Start("localhost:8080")



	e.Start(":8080")
}
