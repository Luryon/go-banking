package storage

import "errors"

var (
	//ErrAccountCantBeNil es usado para marcar cuando un cuenta es mandada como nil
	ErrAccountCantBeNil = errors.New("la cuenta no puede ser nil")
	//ErrAmountExceedAccountAmount .
	ErrAmountExceedAccountAmount = errors.New("la cuenta no tiene fondos suficientes")
)
