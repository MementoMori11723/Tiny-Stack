package database

import (
  "testing"
)

func TestInsert(t *testing.T) {
  testTodo := Todo{"Test 1",false}
  err := Insert(testTodo)
  t.Log(err)
  if err != nil {
    t.Fail()
  }
  t.Log("Insertion successful")
}

func TestDelete(t *testing.T) {
  testTodo := Todo{"Test 1",false}
  err := Delete(testTodo)
  if err != nil {
    t.Fail()
  }
  t.Log("Deletion successful")
}

func TestUpdate(t *testing.T) {
  testTodo := Todo{"Learn Go",true}
  err := Update(testTodo)
  if err != nil {
    t.Fail()
  }
  t.Log("Update successful")
}

func TestFetch(t *testing.T) {
  todos, err := Fetch()
  if err != nil {
    t.Fail()
  }
  t.Log("Fetch successful")
  for _, todo := range todos {
    t.Log(todo.Title, todo.Completed)
  }
}
