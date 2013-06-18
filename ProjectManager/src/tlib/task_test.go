package tlib

import (
	"testing"
)

func TestOps(t *testing.T) {
	tm := NewTaskManager()

	if tm == nil {
		t.Error("new taskmanager failed.")
	}

	if tm.Len() != 0 {
		t.Error("Not empty")
	}

	t0 := &TodoEntry{"1001", "2005-1-1", "9:00", "10:00", "test tomato", 1, 1001}
	tm.Add(t0)

	if tm.Len() != 1 {
		t.Error("add failed")
	}

	todo := tm.FindByTodoId("1001")
	if todo == nil {
		t.Error("todo find failed.")
	}

	m, err := tm.Get(0)
	if m == nil {
		t.Error("get failed", err)
	}

	m = tm.Remove(0)
	if m == nil || tm.Len() != 0 {
		t.Error("remove failed", err)
	}
}
