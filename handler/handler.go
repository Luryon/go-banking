package handler

import "github.com/luryon/go-bank-system/model"

type Storage interface {
	Create(acc *model.Account) error
	Update(id int, acc *model.Account) (*model.Account, error)
	GetAll() (*model.Accounts, error)
	GetById(id int) (*model.Account, error)
	Delete(id int) error
}

type Operator interface {
	Send(tx model.Operation) error
}
