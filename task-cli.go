package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	// "errors"
)

type task struct {
	id          int
	description string
	status      string
	createdAt   string
	updatedAt   string
}

// Listing all tasks
func list() {
	// error handling
	test, _ := json.Marshal(`hello`)
	fmt.Print(string(test))
}

func add(arg string) {
	// read json file to know the id
	// get a different time format
	fmt.Println(task{1, arg, "todo", time.Now().Format(time.RFC822), time.Now().Format(time.RFC822)})
}

func main() {
	// check for args
	arg := os.Args[1]
	switch arg {
	case "list":
		list()
	case "add":
		// check for os.Args[2]
		add(os.Args[2])
	}
	fmt.Println("hi, you just said: ", arg)
}
