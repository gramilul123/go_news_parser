package db

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//The DB struct contains a connect with postgres db or a error connection.
type DB struct {
	Connect *sqlx.DB
	Error   error
}

var db *DB
var once sync.Once

// GetDBConnect returns a db connection struct. If there was any error then a struct will have error.
func GetDBConnect(host, user, password, dbname string) *DB {

	dbFunc := func() {
		strConnect := fmt.Sprintf("host=%s sslmode=disable user=%s password=%s dbname=%s", host, user, password, dbname)
		db = &DB{}
		db.Connect, db.Error = sqlx.Connect("postgres", strConnect)
	}

	once.Do(dbFunc)

	return db
}
