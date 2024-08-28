package main

import (
	"myApp/internal/http"
	"myApp/internal/logic"
	"myApp/internal/myDb"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	//инициализация экземпляра echo и middleware
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//инициализация логгера
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	//инициализация репо для раюоты c базой
	db, err := myDb.NewDb()
	if err != nil {
		logger.Fatalf("Ошбка: %v", err)
	}
	dbRepo := myDb.NewRepository(db)

	//инициализация логики приложения/создвние экземпляра Logic
	appLogic := logic.NewLogic(dbRepo, logger)

	//инициализация обработчиков
	httpHandlers := http.NewPersonHandlers(appLogic)

	//настройка маршрутов HTTP
	e.GET("/get/allperson", httpHandlers.GetAllPerson)
	e.GET("/get/person/:id", httpHandlers.GetPerson)
	e.POST("/create/person", httpHandlers.CreatePerson)
	e.PUT("/update/person/:id", httpHandlers.UpdatePerson)
	e.DELETE("/delete/person/:id", httpHandlers.DeletePerson)

	//запуск сервера
	go func() {
		if err := e.Start(":8080"); err != nil {
			logger.Fatalf("Ошибка запуска сервера %v", err)
		}
	}()
	logger.Info("Cервер успешно запущен на порту :8080")
	//блокирум горутину
	select {}
}
