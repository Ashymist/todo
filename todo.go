package todo

import (
	"encoding/json"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

func (l *List) Complete(i int) error {

	if i <= 0 || i > len(*l) {
		return &ItemDoesntExist{ItemId: i}
	}

	list := *l
	list[i-1].Done = true
	list[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *List) Delete(i int) error {
	if i <= 0 || i > len(*l) {
		return &ItemDoesntExist{ItemId: i}
	}

	list := *l
	*l = append(list[:i-1], list[i:]...)
	return nil
}
