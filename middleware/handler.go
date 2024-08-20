package middleware

import "tiny/stack/database"

func getTodos() ([]database.Todo, error) {
  return database.Fetch() 
}

func addTodo(title string) error {
  t := database.Todo{Title: title, Completed: false}
  return database.Insert(t)
}

func removeTodo(title string) error {
  t := database.Todo{Title: title}
  return database.Delete(t)
}

func todoCompleted(title string) error {
  t := database.Todo{Title: title, Completed: true}
  return database.Update(t)
}
