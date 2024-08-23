package database

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB 

func Connect() {
  var err error
  db, err = sql.Open("sqlite3", "./database.db")
  if err != nil {
    panic(err)
  }
}

// test function

type ReturnType struct {
	Data interface{}
}

type contacts struct {
	Name  string
	Email string
}

func Fetch() ReturnType {
	return ReturnType{
		Data: []contacts{
			{"John Doe", "johndoe@none.com"},
			{"Monkey D. Luffy", "monkeydluffy@chad.com"},
		},
	}
}
