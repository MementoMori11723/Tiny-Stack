package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type response struct {
	Rows *sql.Rows
	err  error
}

func executer(query string, res bool, args ...interface{}) response {
	db, err := connect()
	if err != nil {
		return response{nil, err}
	}
	defer db.Close()

	if res {
		rows, err := db.Query(query, args...)
		if err != nil {
			return response{nil, err}
		}
		return response{rows, nil}
	} else {
		_, err := db.Exec(query, args...)
		return response{nil, err}
	}
}

func connect() (*sql.DB, error) {
	return sql.Open("sqlite3", "db.sqlite3")
}

func Insert(t Todo) error {
   r := executer("INSERT INTO todos (title, completed) VALUES (?, ?)", false, t.Title, t.Completed)
	return r.err
}

func Delete(t Todo) error {
  r := executer("DELETE FROM todos WHERE title = ?", false, t.Title)
  return r.err
}

func Update(t Todo) error {
  r := executer("UPDATE todos SET completed = ? WHERE title = ?", false, t.Completed, t.Title)
  return r.err
}

func Fetch() ([]Todo, error) {
  r := executer("SELECT * FROM todos", true)
  if r.err != nil {
    return nil, r.err
  }
  defer r.Rows.Close()

  var todos []Todo
  for r.Rows.Next() {
    var t Todo
    r.Rows.Scan(&t.Title, &t.Completed)
    todos = append(todos, t)
  }
  return todos, nil
}
