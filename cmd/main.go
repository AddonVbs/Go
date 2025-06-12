package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"BackEnd/internal/db"
	h "BackEnd/internal/handler"
	"BackEnd/internal/taskservice"
	userservice "BackEnd/internal/userservice"

	tasks "BackEnd/internal/web/tasks"
	users "BackEnd/internal/web/users"
	// путь до user-хендлера
)

func main() {
	db.InitDB()

	// ===== TASKS =====
	taskRepo := taskservice.NewTaskRepository(db.Db)
	taskService := taskservice.NewTaskService(taskRepo)
	taskHandler := h.NewStrictTaskHandler(taskService)

	// ===== USERS =====
	userRepo := userservice.NewUserRepository(db.Db)
	userService := userservice.NewUserService(userRepo)
	userHandler := h.NewStrictUserHandler(userService)

	// ===== Echo =====
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ===== Register routes =====
	tasks.RegisterHandlers(e, tasks.NewStrictHandler(taskHandler, nil))
	users.RegisterHandlers(e, users.NewStrictHandler(userHandler, nil))

	// ===== Start server =====
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server fail: %v", err)
	}
}
