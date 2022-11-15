package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/luryon/go-banking/model"
)

type operation struct {
	storage Operator
}

func NewOperation(storage Operator) operation {
	return operation{storage}
}

func (o *operation) Send(c echo.Context) error {

	tx := model.Operation{}
	c.Bind(&tx)
	tx.Time = time.Now()
	tx.State = "Started"

	err := o.storage.Send(&tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusAccepted, "Transaccion mandada")
}
