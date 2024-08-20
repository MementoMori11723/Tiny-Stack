package middleware

import "tiny/stack/database"

func getTodos() ([]database.Todo, error) {
  return database.Fetch() 
}

func addTodo(title string) error { 
  return database.Insert(database.Todo{
    Title: title, 
    Completed: false,
  })
}

func removeTodo(title string) error {
  return database.Delete(database.Todo{
    Title: title,
  })
}

func todoCompleted(title string) error {
  return database.Update(database.Todo{
    Title: title, 
    Completed: true,
  })
}
