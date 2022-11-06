package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/luryon/go-bank-system/model"
)

type operation struct {
	CurrentID int
	operation Operator
}

func NewOperation(o Operator) operation {
	return operation{0, o}
}

func (o *operation) Send(c echo.Context) error {
	o.CurrentID += 1

	tx := model.Operation{}
	c.Bind(&tx)
	fmt.Printf("%v\n", tx)
	tx.Id = o.CurrentID
	tx.Time = time.Now()
	tx.State = "Started"

	err := o.operation.Send(tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusAccepted, "Transaccion mandada")
}
