package model

import (
	"time"
)

// account es el modelo de cada cuenta
type Account struct {
	Id             int       `json:"id"`
	Token          string    `json:"token"`
	Name           string    `json:"name"`
	Last_name      string    `json:"last_name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Amount_on_acc  float64   `json:"amount_on_acc"`
	Created_at     time.Time `json:"created_at"`
	Last_operation time.Time `json:"last_operation"`
}

// accounts es una lista de cuentas
type Accounts []*Account

// NewAccount genera una nueva cuenta
func (list_acc *Accounts) NewAccount(token, name, last_name string, amount float64) {
	acc := Account{
		Id:             len(*list_acc) + 1,
		Token:          token,
		Name:           name,
		Last_name:      last_name,
		Amount_on_acc:  amount,
		Created_at:     time.Now(),
		Last_operation: time.Time{},
	}
	*list_acc = append(*list_acc, &acc)
}

// NewAccountsList crea una lista de cuentas
func NewAccountsList() *Accounts {
	return &Accounts{}
}
