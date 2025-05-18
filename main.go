package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate: %v", err)

	}

}

type Task struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Task1 string `json:"task"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// all global var

var IDtask = 1

func GetHandler(c echo.Context) error {
	var task []Task

	if err := db.Find(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: "error", Message: "Could not get task"})
	}

	return c.JSON(http.StatusOK, &task)
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

	if err := db.Create(&t).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Status: "Error", Message: "Could not add task"})
	}

	return c.JSON(http.StatusOK, t)
}

func PatchHandler(c echo.Context) error {
	IDparam := c.Param("id")
	id, err := strconv.Atoi(IDparam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: "Bad", Message: "Invalid Patch !"})
	}

	var updateTask Task
	if err := c.Bind(&updateTask); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not update Task",
		})
	}

	var existingTask Task
	if err := db.First(&existingTask, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, Response{
			Status:  "Error",
			Message: "Task not found",
		})
	}

	existingTask.Task1 = updateTask.Task1
	if err := db.Save(&existingTask).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Could not update Task",
		})
	}

	return c.JSON(http.StatusOK, existingTask)
}

func Deletahendler(c echo.Context) error {
	IDparam := c.Param("id")
	id, err := strconv.Atoi(IDparam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Status: "404"})
	}

	if err := db.Delete(&Task{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Could not delete Task",
		})
	}

	return c.JSON(http.StatusNoContent, Response{Status: "Success", Message: "Was deleted"})
}

func main() {
	e := echo.New()

	e.GET("/tasks", GetHandler)
	e.POST("/tasks", PostHandler)
	e.PATCH("/tasks/:id", PatchHandler)
	e.DELETE("tasks/:id", Deletahendler)

	e.Start(":8080")
}
