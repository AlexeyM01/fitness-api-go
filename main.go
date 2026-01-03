// main.go
package main

import (
	"fitness-api/cmd/handlers" // присоединяем хэндлеры
	"fitness-api/cmd/storage"

	"github.com/labstack/echo/v4" // сам фреймворк echo
)

func main() {
	e := echo.New()

	storage.InitDB()

	e.Use(handlers.LogRequest)

	e.GET("/", handlers.Home)
	e.GET("/users", handlers.HandleGetUsers)
	e.POST("/users/", handlers.HandleCreateUser)
	e.PUT("/users/:id", handlers.HandleUpdateUser)

	e.POST("/measurements", handlers.HandleCreateMeasurement)
	e.PUT("/measurements/:id", handlers.HandleUpdateMeasurement)

	e.Logger.Fatal(e.Start(":8080"))
}
