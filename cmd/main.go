package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	"github.com/luryon/go-banking/handler"
	"github.com/luryon/go-banking/storage"
)

func main() {
	e := echo.New()
	godotenv.Load()

	storage.InitMySqlDB()
	db := storage.Pool()

	accStorage := storage.NewMySQLAccount(db)
	accService := handler.NewAccount(accStorage)

	// opeService := handler.NewOperation(&mem)
	accounts := e.Group("/accounts")
	// operations := e.Group("/operations")

	accounts.GET("/", accService.Migrate)
	// accounts.POST("/new", accService.Create)
	// accounts.PUT("/update/:id", accService.Update)
	// accounts.GET("", accService.GetAll)
	// accounts.GET("/:id", accService.GetById)
	// accounts.DELETE("/:id", accService.Delete)
	// operations.POST("/send", opeService.Send)

	// if err := e.StartTLS(":443", os.Getenv("CRT_PATH_WIN"), os.Getenv("KEY_PATH_WIN")); err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
