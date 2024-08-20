package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func connect() (*sql.DB, error) {
	return sql.Open("sqlite3", "./todos.db")
}

func Insert(t Todo) error {
	db, err := connect()
	defer db.Close()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", t.Title, t.Completed)
	return err
}

func Delete(t Todo) error {
	db, err := connect()
	defer db.Close()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM todos WHERE title = ?", t.Title)
	return err
}

func Update(t Todo) error {
	db, err := connect()
	defer db.Close()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE todos SET completed = ? WHERE title = ?", t.Completed, t.Title)
	return err
}

func Fetch() ([]Todo, error) {
	db, err := connect()
	defer db.Close()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		err = rows.Scan(&t.Title, &t.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}
