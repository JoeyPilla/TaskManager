package dbclient

import "time"

type IBoltClient interface {
	OpenBoltDB()
	ListTodos()
	AddTodo(string, time.Time)
	RemoveTodo(string)
}
