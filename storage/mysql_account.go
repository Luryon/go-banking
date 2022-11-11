package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/luryon/go-banking/model"
)

const (
	GetEpassword = `SELECT password FROM accounts WHERE id = ?`
	MigrateMySQL = `CREATE TABLE IF NOT EXISTS accounts(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		token VARCHAR(100) not null,
		name VARCHAR(25) not null,
		last_name VARCHAR(25),
		email VARCHAR(50) not null,
		password VARCHAR(100) not null,
		amount_on_acc INT not null DEFAULT 0,
		created_at TIMESTAMP not null DEFAULT Now(),
		last_operation TIMESTAMP
	)`
	CreateMySQL = `INSERT INTO accounts(token, name, last_name, email,
	password, amount_on_acc, created_at)
	VALUES(?, ?, ?, ?, ?, ?, ?)`
	UpdateMySQL = `UPDATE accounts SET name = ?, last_name = ?, token = ?, last_operation = ?
	WHERE id = ?`
)

// mySQLProduct used fir work with postgres - product
type mySQLAccount struct {
	db *sql.DB
}

// NewmySQLProduct return new pointer of mySQLProduct
func NewMySQLAccount(db *sql.DB) *mySQLAccount {
	return &mySQLAccount{db}
}

func (m *mySQLAccount) Migrate() error {
	stmt, err := m.db.Prepare(MigrateMySQL)
	if err != nil {
		return err
	}

	result, err := stmt.Exec()
	if err != nil {
		return nil
	}

	fmt.Println(result)
	return nil
}

func (m *mySQLAccount) Create(acc *model.Account) error {
	if acc == nil {
		return ErrAccountCantBeNil
	}

	stmt, err := m.db.Prepare(CreateMySQL)
	if err != nil {
		return err
	}

	acc.Password, err = Hash_Password(acc.Password)
	if err != nil {
		return err
	}

	fmt.Println(acc.Password)

	stmt.Exec(
		acc.Token,
		acc.Name,
		acc.Last_name,
		acc.Email,
		acc.Password,
		1500,
		time.Now(),
	)

	fmt.Println("Cuenta Creada")
	return nil
}

func (m *mySQLAccount) Update(id int, acc *model.Account) error {
	res := m.db.QueryRow(GetEpassword, id)
	fmt.Println(res)

	var e_pass string
	res.Scan(&e_pass)
	fmt.Println(e_pass)

	check, err := check_password(acc.Password, e_pass)
	if err != nil {
		return err
	}

	fmt.Println(check)

	stmt, err := m.db.Prepare(UpdateMySQL)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(acc.Name, acc.Last_name, acc.Token, time.Now(), id)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
