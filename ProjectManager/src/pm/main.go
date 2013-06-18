package main

import (
	"fmt"
	"tlib"
)

var tm *tlib.TaskManager

func main() {
	fmt.Println("lib init")
	tm = tlib.NewTaskManager()

	t0 := &tlib.TodoEntry{"1001", "2005-1-1", "9:00", "10:00", "test tomato", 1, 1001}
	tm.Add(t0)

	if tm.Len() != 1 {
		fmt.Println("add failed")
	}

	todo := tm.FindByTodoId("1001")
	if todo != nil {
		fmt.Println(todo.Todo_Id)
	}

	fmt.Println("taskid is", tm.TodoList[0].Todo_Id)

	tm.RemoveByTodoID("1001")

	fmt.Println(tm.Len())
}
