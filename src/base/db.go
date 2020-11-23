package base

import (
	"github.com/cspinetta/otsql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GetDB() (*sqlx.DB, error) {
	driverName, err := getDriver()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect(driverName, ":memory:")
	if err != nil {
		return nil, err
	}

	dbSetup(db)

	return db, nil
}

func getDriver() (string, error) {
	return otsql.Register(
		"sqlite3",
		otsql.WithAllowRoot(true),
		otsql.WithQuery(true),
		//otsql.WithQueryParams(true),
		otsql.WithInstanceName("sqlite3InMemory"),
	)
}

func dbSetup(db *sqlx.DB) {
	var err error

	_, err = db.Exec("DROP TABLE IF EXISTS t1")
	if err != nil {
		log.Panic(err)
	}

	tableDDL := `CREATE TABLE user (
 	id INTEGER PRIMARY KEY AUTOINCREMENT,
 	name VARCHAR NOT NULL,
 	birthday DATETIME NOT NULL,
 	created_at DATETIME NOT NULL,
 	updated_at DATETIME DEFAULT NULL
)`

	_, err = db.Exec(tableDDL)
	if err != nil {
		log.Panic(err)
	}
}
