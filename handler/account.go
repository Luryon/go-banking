package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/luryon/go-banking/model"
)

type account struct {
	storage Storage
}

func NewAccount(storage Storage) account {
	return account{storage}
}

func (a *account) Migrate(c echo.Context) error {
	return a.storage.Migrate()
}

func (a *account) Create(c echo.Context) error {
	data := model.Account{}
	err := c.Bind(&data)
	if err != nil {
		return err
	}
	data.Created_at = time.Now()
	if err != nil {
		return err
	}
	a.storage.Create(&data)
	return c.JSON(http.StatusCreated, data)
}

func (a *account) Update(c echo.Context) error {
	data := model.Account{}
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	data.Last_operation = time.Now()
	err = a.storage.Update(id, &data)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Cuenta Acualizada Correctamente")
}

func (a *account) GetAll(c echo.Context) error {
	resp, err := a.storage.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (a *account) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	resp, err := a.storage.GetById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (a *account) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	err = a.storage.Delete(id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Cuenta eliminada Correctamente")
}
