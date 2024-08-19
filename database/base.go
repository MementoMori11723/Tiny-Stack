package database

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
