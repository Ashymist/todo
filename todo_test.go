package todo

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
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
