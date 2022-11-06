package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/luryon/go-banking/handler"
	"github.com/luryon/go-banking/storage"
)

func main() {
	e := echo.New()

	mem := storage.NewMemory()
	accService := handler.NewAccount(&mem)
	opeService := handler.NewOperation(&mem)

	accounts := e.Group("/accounts")
	operations := e.Group("/operations")

	accounts.POST("/new", accService.Create)
	accounts.PUT("/update/:id", accService.Update)
	accounts.GET("", accService.GetAll)
	accounts.GET("/:id", accService.GetById)
	accounts.DELETE("/:id", accService.Delete)
	operations.POST("/send", opeService.Send)
	e.Start(":8080")
}
