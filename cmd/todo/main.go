package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Ashymist/todo"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	if _, err := os.Stat(todoFileName); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
