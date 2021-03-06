package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sguessou/todo"
)

var todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for sguessou inc\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2022\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage information:\n")
		flag.PrintDefaults()
	}

	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	filter := flag.Bool("filter", false, "List uncompleted tasks")
	del := flag.Int("del", 0, "Item to be deleted")
	date := flag.Bool("date", false, "Show current date")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *date:
		fmt.Fprintf(os.Stdout, "The current date is: %s", todo.ShowCurrentDate())
	case *list:
		fmt.Print(l)
	case *filter:
		l.FilterUncomplete()
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		if err := l.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		tasks, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, t := range tasks {
			l.Add(t)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) ([]string, error) {
	buf := make([]string, 0)
	if len(args) > 0 {
		return append(buf, strings.Join(args, " ")), nil
	}

	s := bufio.NewScanner(r)

	for s.Scan() {
		if err := s.Err(); err != nil {
			return []string{}, err
		}
		if len(s.Text()) == 0 {
			return []string{}, fmt.Errorf("Task cannot be blank")
		}
		buf = append(buf, s.Text())
	}

	return buf, nil
}
