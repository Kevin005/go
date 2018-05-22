package dbServe

import (
	"fmt"
	"awesomeProject/web/restFul/bean"
)

var currentId int
//初始化了一个Todos类型变量
var todos bean.Todos

// Give us some seed data
func init() {
	RepoCreateTodo(bean.Todo{Name: "Write presentation"})
	RepoCreateTodo(bean.Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) bean.Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return bean.Todo{}
}

func RepoCreateTodo(t bean.Todo) bean.Todos {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return todos
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
