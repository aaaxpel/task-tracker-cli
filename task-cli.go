package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	// "errors"
)

type Task struct {
	id          int
	description string
	status      string
	createdAt   string
	updatedAt   string
}

// Listing all tasks
func list(tasks []Task) {
	fmt.Println(tasks)
}

func add(tasks []Task, description string) {
	id := 1
	if tasks != nil {
		// assuming last added task has the highest id
		id += tasks[len(tasks)-1].id
	}
	task := Task{id, description, "todo", time.Now().Format(time.RFC822), time.Now().Format(time.RFC822)}

	tasks = append(tasks, task)

	// !! problem somewhere here, either gets lost in json conversion or doesn't write to file
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		os.Create("tasks.json")
	}
	var tasks []Task
	json.Unmarshal(file, &tasks)

	// check for args
	arg := os.Args[1]
	switch arg {
	case "list":
		list(tasks)
	case "add":
		// check for os.Args[2]
		add(tasks, os.Args[2])
	}
}
