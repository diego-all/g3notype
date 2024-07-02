package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQLite(dsn string) (*DB, error) {

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	dbConn.SQL = db

	return dbConn, nil
}

func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Error!", err)
		return err
	}
	fmt.Println("Database ping successful!")
	return nil
}
