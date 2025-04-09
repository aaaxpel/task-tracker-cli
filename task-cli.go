package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

type Task struct {
	// Capitalization so it's visible to json.MarshalIndent
	Id          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

// Listing all tasks
func list(tasks []Task) {
	for _, task := range tasks {
		if task.UpdatedAt == task.CreatedAt {
			fmt.Printf("- [%v] %v - %v (%v)\n", task.Id, task.Description, task.Status, task.CreatedAt)
		} else {
			fmt.Printf("- [%v] %v - %v (%v | Updated: %v)\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}

// Listing tasks by status
func listStatus(tasks []Task, status string) {
	var tasksByStatus []Task
	for i := range tasks {
		if tasks[i].Status == status {
			tasksByStatus = append(tasksByStatus, tasks[i])
		}
	}
	if len(tasksByStatus) == 0 {
		fmt.Printf("Tasks with status %v not found\n", status)
	} else {
		fmt.Printf("Tasks with status %v:\n", status)
		list(tasksByStatus)
	}
}

// Adding a new task
func add(tasks []Task, description string) {
	id := 1
	if tasks != nil {
		id += tasks[len(tasks)-1].Id
	}
	task := Task{id, description, "todo", time.Now().Format("Jan 2, 2006 15:04"), time.Now().Format("Jan 2, 2006 15:04")}

	tasks = append(tasks, task)

	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Task added successfully (ID: %v)", task.Id)
	}
}

// Updating a new task
func update(tasks []Task, id string, description string) {
	for i := range tasks {
		if strconv.Itoa(tasks[i].Id) == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now().Format("Jan 2, 2006 15:04")
			save(tasks)
			fmt.Printf("Task updated successfully (ID: %v)", id)
			return
		}
	}
	fmt.Printf("Task was not found (ID: %v)", id)
}

// Deleting a new task
func delete(tasks []Task, id string) {
	for i := range tasks {
		if strconv.Itoa(tasks[i].Id) == id {
			tasks = slices.Delete(tasks, i, i+1)
			save(tasks)
			fmt.Printf("Task deleted successfully (ID: %v)", id)
			return
		}
	}
	fmt.Printf("Task was not found (ID: %v)", id)
}

// Mark task as done or in-progress
func mark(tasks []Task, id string, status string) {
	for i := range tasks {
		if strconv.Itoa(tasks[i].Id) == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now().Format("Jan 2, 2006 15:04")
			save(tasks)
			fmt.Printf("Task updated successfully (ID: %v)", id)
			return
		}
	}
	fmt.Printf("Task was not found (ID: %v)", id)
}

// Save to tasks.json
func save(tasks []Task) bool {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		panic(err)
	} else {
		return true
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
	switch os.Args[1] {
	case "list":
		if len(os.Args) >= 3 {
			// check if it's valid
			listStatus(tasks, os.Args[2])
		} else {
			list(tasks)
		}
	case "add":
		// check for os.Args[2]
		add(tasks, os.Args[2])
	case "delete":
		// check for os.Args[2] = int
		delete(tasks, os.Args[2])
	case "update":
		// check for os.Args[2] = int
		update(tasks, os.Args[2], os.Args[3])
	case "mark-in-progress":
		mark(tasks, os.Args[2], "in-progress")
	case "mark-done":
		mark(tasks, os.Args[2], "done")
	}
}
