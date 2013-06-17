package tlib

import (
	"errors"
)

type TaskManager struct {
	TodoList []TodoEntry //任务列表
}

func NewTaskManager() *TaskManager {
	return &TaskManager{make([]TodoEntry, 0)}
}

func (t *TaskManager) Len() int {
	return len(t.TodoList)
}

func (t *TaskManager) Add(todo *TodoEntry) {
	t.TodoList = append(t.TodoList, *todo)
}

func (t *TaskManager) Get(index int) (todo *TodoEntry, err error) {
	if index < 0 || index >= len(t.TodoList) {
		return nil, errors.New(("Index out of range."))
	}
	return &t.TodoList[index], nil
}

func (t *TaskManager) FindByTodoId(todoid string) *TodoEntry {
	if len(t.TodoList) == 0 {
		return nil
	}
	for _, v := range t.TodoList {
		if v.todo_id == todoid {
			return &v
		}
	}
	return nil
}

func (t *TaskManager) Remove(index int) *TodoEntry {
	if index < 0 || index >= len(t.TodoList) {
		return nil
	}

	removeTodo := &t.TodoList[index]
	t.TodoList = append(t.TodoList[:index], t.TodoList[index+1:]...)
	return removeTodo
}

func (t *TaskManager) RemoveByTodoID(todoid string) *TodoEntry {
	if len(t.TodoList) == 0 {
		return nil
	}

	for i, v := range t.TodoList {
		if v.todo_id == todoid {
			return t.Remove(i)
		}
	}
	return nil
}
