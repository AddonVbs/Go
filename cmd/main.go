package main

import (
	taskhendler "BackEnd/internal/TaskHendler"
	ts "BackEnd/internal/TaskService"
	"BackEnd/internal/db"
	"log"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Task1 string `json:"task"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	dataBase, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	taskRepo := ts.NewTaskRepository(dataBase)
	taskServices1 := ts.NewTaskService(taskRepo)
	taskhendlers := taskhendler.NewTaskHendler(taskServices1)

	e := echo.New()
	e.GET("/tasks", taskhendlers.GetHandler)
	e.POST("/tasks", taskhendlers.PostHandler)
	e.PATCH("/tasks/:id", taskhendlers.PatchHandler)
	e.DELETE("/tasks/:id", taskhendlers.DeleteHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
