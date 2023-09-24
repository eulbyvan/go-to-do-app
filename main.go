package main

import (
	"bufio"
	"fmt"
	"os"
)

var tasks []string

func main() {
	for {
		fmt.Println("To-Do List Application")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Quit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask()
		case 2:
			listTasks()
		case 3:
			removeTask()
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addTask() {
	fmt.Print("Enter the task: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	task := scanner.Text()

	tasks = append(tasks, task)
	fmt.Println("Task added successfully!")
}

func listTasks() {
	fmt.Println("Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func removeTask() {
	fmt.Print("Enter the task number to remove: ")
	var taskNum int
	fmt.Scanln(&taskNum)

	if taskNum < 1 || taskNum > len(tasks) {
		fmt.Println("Invalid task number. Please try again.")
		return
	}

	// Remove the task by creating a new slice without the selected task.
	tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)

	fmt.Println("Task removed successfully!")
}
