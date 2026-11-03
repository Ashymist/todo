package todo

import "testing"

func TestAdd(t *testing.T) {
	list := List{}
	task := "take out garbage"
	list.Add(task)
	if len(list) != 1 {
		t.Errorf("Error: The list is empty")
	}
	if list[0].Task != task {
		t.Errorf("Error: Expected %q, got %q instead", task, list[0].Task)
	}

}

func TestComplete(t *testing.T) {
	list := List{}
	task := "take out garbage"
	list.Add(task)
	if list[0].Done {
		t.Errorf("Task should not be done")
	}
	err := list.Complete(1)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if !list[0].Done {
		t.Errorf("Task should be done")
	}
}
