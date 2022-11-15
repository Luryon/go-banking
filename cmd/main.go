package main

import (
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

	opeStorage := storage.NewMySQLOperation(db)
	opeService := handler.NewOperation(opeStorage)
	accounts := e.Group("/accounts")
	operations := e.Group("/operations")

	accounts.GET("/", accService.Migrate)
	accounts.POST("/new", accService.Create)
	accounts.PUT("/update/:id", accService.Update)
	accounts.GET("", accService.GetAll)
	accounts.GET("/:id", accService.GetById)
	accounts.DELETE("/:id", accService.Delete)
	operations.POST("/send", opeService.Send)

	// if err := e.StartTLS(":443", os.Getenv("TEST_CRT_PATH_WIN"), os.Getenv("TEST_KEY_PATH_WIN")); err != http.ErrServerClosed {
	// 	log.Fatal(err)
	// }

	e.Start(":8080")
}
