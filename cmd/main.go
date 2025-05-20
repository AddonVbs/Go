package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Task1 string `json:"task"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func GetHandler(c echo.Context) error {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{"error", "Could not get tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func PostHandler(c echo.Context) error {
	var t Task
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, Response{"error", "Invalid payload"})
	}

	if err := db.Create(&t).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{"error", "Could not add task"})
	}

	return c.JSON(http.StatusOK, t)
}

func PatchHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{"error", "Invalid ID"})
	}

	var payload Task
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, Response{"error", "Invalid payload"})
	}

	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, Response{"error", "Task not found"})
	}

	task.Task1 = payload.Task1
	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{"error", "Could not update task"})
	}

	return c.JSON(http.StatusOK, task)
}

func DeleteHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{"error", "Invalid ID"})
	}

	if err := db.Delete(&Task{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{"error", "Could not delete task"})
	}

	return c.JSON(http.StatusNoContent, Response{"Seccess", "Was Del."})
}

func main() {
	initDB()

	e := echo.New()
	e.GET("/tasks", GetHandler)
	e.POST("/tasks", PostHandler)
	e.PATCH("/tasks/:id", PatchHandler)
	e.DELETE("/tasks/:id", DeleteHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
