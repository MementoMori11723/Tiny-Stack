package database

import (
	"database/sql"
	"github.com/MementoMori11723/Tiny-Stack/config"
	_ "github.com/mattn/go-sqlite3"
)

var (
  DB *sql.DB
  DATABASE_PATH string
)

func init() {
  var err error
  config.LoadDatabasePath(&DATABASE_PATH)
  DB, err = sql.Open("sqlite3", DATABASE_PATH)
  if err != nil {
    panic(err)
  }
}

