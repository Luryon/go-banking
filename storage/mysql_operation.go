package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/luryon/go-banking/model"
)

const (
	getAcc           = `SELECT * FROM accounts WHERE id = ?`
	UpdateAccBalance = `UPDATE accounts SET amount_on_acc = ?, last_operation = ?
	WHERE id = ?`
)

type mySQLOperation struct {
	db *sql.DB
}

// NewmySQLProduct return new pointer of mySQLProduct
func NewMySQLOperation(db *sql.DB) *mySQLOperation {
	return &mySQLOperation{db}
}

func (o *mySQLOperation) Send(tx *model.Operation) error {
	row := o.db.QueryRow(getAcc, tx.Origin_acc)
	origin_acc, err := scanRowAccount(row)
	if err != nil {
		return err
	}

	if !(origin_acc.Amount_on_acc > tx.Amount) {
		return ErrAmountExceedAccountAmount
	}

	row = o.db.QueryRow(getAcc, tx.Dest_acc)
	dest_acc, err := scanRowAccount(row)
	if err != nil {
		return err
	}

	fmt.Println(origin_acc, dest_acc)

	origin_acc.Amount_on_acc = origin_acc.Amount_on_acc - tx.Amount
	dest_acc.Amount_on_acc = dest_acc.Amount_on_acc + tx.Amount

	o.UpdateAccBalance(origin_acc.Id, origin_acc)
	o.UpdateAccBalance(dest_acc.Id, dest_acc)

	return nil
}

func (o *mySQLOperation) UpdateAccBalance(id int, acc *model.Account) error {
	stmt, err := o.db.Prepare(UpdateAccBalance)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(
		acc.Amount_on_acc,
		time.Now(),
		id)
	if err != nil {
		return err
	}

	if rows, err := result.RowsAffected(); rows != 1 {
		return err
	}

	return nil
}
