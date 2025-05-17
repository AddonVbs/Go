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

var IDtask = 1
var my_task = make(map[int]Task)

func GetHandler(c echo.Context) error {
	var t []Task

	for _, key := range my_task {
		t = append(t, key)

	}
	return c.JSON(http.StatusOK, &t)
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

	my_task[t.ID] = t

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

	if _, exist := my_task[id]; !exist {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not updata Task",
		})
	}

	updataTask.ID = id
	my_task[id] = updataTask

	return c.JSON(http.StatusOK, Response{Status: "Success", Message: "Was update "})

}

func Deletahendler(c echo.Context) error {
	IDparam := c.Param("id")

	id, err := strconv.Atoi(IDparam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: "Bad", Message: "Invalid Patch !"})
	}

	if _, exist := my_task[id]; !exist {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not updata Task",
		})
	}

	delete(my_task, id)
	return c.JSON(http.StatusOK, Response{Status: "Success", Message: "Was delete "})

}

func main() {
	e := echo.New()

	e.GET("/task", GetHandler)
	e.POST("/task", PostHandler)
	e.PATCH("/task/:id", PatchHandler)
	e.DELETE("task/:id", Deletahendler)

	e.Start(":8080")
}
