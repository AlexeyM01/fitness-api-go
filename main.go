// main.go
package main

import (
	"fitness-api/cmd/handlers" // присоединяем хэндлеры
	"fitness-api/cmd/repositories"
	"fitness-api/cmd/storage"
	"log"

	"github.com/labstack/echo/v4" // сам фреймворк echo
)

func main() {
	e := echo.New()

	// инициализация БД
	db, err := storage.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// инициализация репозиториев
	userRepo := repositories.NewUserRepository(db)
	measureRepo := repositories.NewMeasurementRepository(db)

	// инициализация хэндлеров
	userHandler := handlers.NewUserHandler(userRepo)
	measureHandler := handlers.NewMeasurementHandler(measureRepo)

	e.Use(handlers.LogRequest)

	// роуты
	e.GET("/", handlers.Home)

	e.GET("/users", userHandler.HandleGetUsers)
	e.POST("/users/", userHandler.HandleCreateUser)
	e.PUT("/users/:id", userHandler.HandleUpdateUser)

	e.POST("/measurements", measureHandler.HandleCreateMeasurement)
	e.PUT("/measurements/:id", measureHandler.HandleUpdateMeasurement)

	e.Logger.Fatal(e.Start(":8080"))
}
