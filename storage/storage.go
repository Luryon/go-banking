package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func InitMySqlDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", os.Getenv("DB_HOST"))
		if err != nil {
			log.Fatalf("can't open the DB, error: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't ping the DB, error: %v", err)
		}

		fmt.Println("Connected to MySQL")

	})

}

// Pool return a once instanse of db
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}
