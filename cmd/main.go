package main

import (
	"BackEnd/internal/db"
	"BackEnd/internal/taskhandler"
	"BackEnd/internal/taskservice"
	"log"

	"github.com/labstack/echo/v4"
)

/*
	type Task struct {
		ID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
		Task1 string `json:"task"`
	}

	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
*/
func main() {
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	repo := taskservice.NewTaskRepository(dbConn)
	svc := taskservice.NewTaskService(repo)
	handler := taskhandler.NewTaskHandler(svc)

	e := echo.New()
	e.POST("/tasks", handler.PostHandler)
	e.GET("/tasks", handler.GetHandler)
	e.PATCH("/tasks/:id", handler.PatchHandler)
	e.DELETE("/tasks/:id", handler.DeleteHandler)

	log.Fatal(e.Start(":8080"))
}
