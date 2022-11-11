package storage

import (
	"database/sql"
	"fmt"
)

const (
	MigrateMySQL = `CREATE TABLE IF NOT EXISTS accounts(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) not null,
		token VARCHAR(250) not null,
		last_name VARCHAR(25),
		amount_on_acc INT not null DEFAULT 0,
		created_at TIMESTAMP not null DEFAULT Now(),
		last_operation TIMESTAMP
	)`
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
