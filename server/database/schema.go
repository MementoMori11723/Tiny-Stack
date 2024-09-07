package database

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "log"
)

var (
  DB *sql.DB
)

func init() {
  var err error
  DB, err = sql.Open("sqlite3", "./tiny-stack.db")
  if err != nil {
    log.Fatal(err)
  }
  if err = DB.Ping(); err != nil {
    log.Fatal(err)
  }
}
