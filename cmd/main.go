package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	echo "github.com/labstack/echo/v4"
	"github.com/luryon/go-banking/handler"
	"github.com/luryon/go-banking/storage"
)

func main() {
	uri := os.Getenv("DATABASE_URI")
	fmt.Println(uri)
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

	if err := e.StartTLS(":443", "/usr/local/go-api/go-banking/cmd/fullchain.pem", "/usr/local/go-api/go-banking/cmd/privkey.pem"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
