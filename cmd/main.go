package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"BackEnd/internal/db"
	"BackEnd/internal/handler"
	"BackEnd/internal/taskservice"
	tasks "BackEnd/internal/web/tasks"
)

func main() {
	// 1) Инициализируем БД и миграции
	db.InitDB()
	// … допустим, вы уже мигрировали (AutoMigrate) где-то в InitDB

	// 2) Создаём repository → service → handler
	repo := taskservice.NewTaskRepository(db.Db)
	service := taskservice.NewTaskService(repo)
	handler := handler.NewStrictTaskHandler(service)

	// 3) Заводим Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 4) Регистрируем «строгий» хендлер, сгенеренный oapi-codegen
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	// 5) Запускаем сервер
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("server fail: %v", err)
	}
}
