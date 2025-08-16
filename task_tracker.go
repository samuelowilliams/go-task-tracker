package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

type Data struct {
	Tasks []Task
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
		tasks := make([]Task, 0)
		dataBytes, err := json.Marshal(tasks)
		if err != nil {
			return err
		}

		err = os.WriteFile(filename, dataBytes, 0664)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadData() {
	filename := "tasks.json"
	var tasks []Task
	err := checkFile(filename)
	check(err)

	file, err := os.ReadFile("data.json")
	check(err)

	err = json.Unmarshal(file, &tasks)
	check(err)
}

func checkCommand(arguments []string, argumentsLength int) {
	var command string = arguments[0]
	switch command {
	case "add":
		if argumentsLength == 2 {
			fmt.Println("Run add command")
		} else {
			fmt.Println("No argument was found: Kindly Input argument for add command")
		}
	case "delete":
		if argumentsLength == 2 {
			fmt.Println("Run delete command")
		} else {
			fmt.Println("No argument was found: Kindly Input argument for delete command")
		}
	case "update":
		if argumentsLength == 3 {
			fmt.Println("Run update command")
		} else {
			fmt.Println("No argument was found: Kindly Input argument for update command")
		}
	case "mark-in-progress":
		if argumentsLength == 2 {
			fmt.Println("Run mark-in-progress command")
		} else {
			fmt.Println("No argument was found: Kindly Input argument for mark-in-progress command")
		}
	case "mark-done":
		if argumentsLength == 2 {
			fmt.Println("Run mark-done command")
		} else {
			fmt.Println("No argument was found: Kindly Input argument for mark-done command")
		}
	case "list":

		var subCommand string
		if argumentsLength == 2 {
			subCommand = arguments[1]
		} else {
			subCommand = ""
		}
		switch subCommand {
		case "todo":

			fmt.Println("Run list todo command")

		case "in-progress":

			fmt.Println("Run list in-progress command")
		case "done":

			fmt.Println("Run list done command")
		default:
			fmt.Println("Run list command")
		}

	}
}

func main() {
	// var tasks []Task
	loadData()
	arguments := os.Args[1:]
	argumentsLength := len(arguments)
	if argumentsLength == 0 {
		fmt.Println("No command was entered: Kindly input a  command")
		return
	}
	checkCommand(arguments, argumentsLength)
	fmt.Println("Length of arguments", argumentsLength)
}
