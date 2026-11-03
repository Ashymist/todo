package todo_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	. "github.com/Ashymist/todo"
)

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

func TestSave(t *testing.T) {
	list := List{}
	task := "take out garbage"
	list.Add(task)

	js, err := json.Marshal(list)
	if err != nil {
		t.Fatalf("%q", err.Error())
	}

	if err := list.Save("tasks.json"); err != nil {
		t.Fatalf("Tried saving file, got %q instead", err)
	}

	file, err := os.ReadFile("tasks.json")
	if err != nil {
		t.Fatalf("Tried to read from a saved file, got %q instead", err)
	}

	if !bytes.Equal(js, file) {
		t.Errorf("Contents before and after save do not match")
	}

}

func TestDeleteInMiddle(t *testing.T) {
	list := List{}
	list.Add("task1")
	list.Add("task2")
	list.Add("task3")

	if err := list.Delete(2); err != nil {
		t.Fatalf("Tried to delete task, got %q", err.Error())
	}

	if len(list) != 2 {
		t.Fatalf("Expected length 2 after deletion, got %v instead", len(list))
	}

	if list[0].Task != "task1" && list[1].Task != "task3" {
		t.Fatalf("Expected [task1,task3] after deletion, got [%q,%q] instead", list[0].Task, list[1].Task)
	}
}

func TestGet(t *testing.T) {
	list := List{}
	if err := list.Get("tasks.json"); err != nil {
		t.Fatalf("Tried to get tasks from file, got %q instead", err.Error())
	}
	if len(list) == 0 {
		t.Errorf("List should not be empty")
	}
	if list[0].Task != "take out garbage" {
		t.Errorf("Contents of file and list do not match after importing")
	}

}
