package operations

import "github.com/luryon/go-banking/model"

type new_operation struct {
	send_acc int
	rec_acc  int
	amount   float64
}

func SendAmount(mem []model.Accounts, send_acc, rec_acc int, amount float64) error {
	o := new_operation{send_acc, rec_acc, amount}
	_ = o
	return nil
}
