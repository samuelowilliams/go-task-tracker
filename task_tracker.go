package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

type Data struct {
	ID    int
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
		data := Data{ID: 0, Tasks: make([]Task, 0)}
		dataBytes, err := json.Marshal(data)
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

func loadData() Data {
	filename := "tasks.json"
	var data Data
	err := checkFile(filename)
	check(err)

	file, err := os.ReadFile(filename)
	check(err)

	err = json.Unmarshal(file, &data)
	check(err)
	return data
}

func commands(arguments []string, argumentsLength int, dataID int, tasks []Task) {
	var command string = arguments[0]
	switch command {
	case "add":
		if argumentsLength == 2 {
			description := arguments[1]
			add(dataID, description, tasks)
		} else {
			fmt.Println("No argument was found: Kindly Input argument for add command")
		}
	case "delete":
		if argumentsLength == 2 {
			id, err := strconv.Atoi(arguments[1])
			check(err)
			delete(dataID, id, tasks)
		} else {
			fmt.Println("No argument was found: Kindly Input argument for delete command")
		}
	case "update":
		if argumentsLength == 3 {
			id, err := strconv.Atoi(arguments[1])
			check(err)
			description := arguments[2]
			update(dataID, id, tasks, description)
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

func updateData(dataID int, tasks []Task) {
	data := Data{ID: dataID, Tasks: tasks}
	dataBytes, err := json.Marshal(data)
	check(err)

	err = os.WriteFile("tasks.json", dataBytes, 0664)

	check(err)
}

type TaskAndPosition struct {
	position int
	task     Task
}

func findTask(id int, tasks []Task) TaskAndPosition {
	start := 0
	end := len(tasks) - 1
	mid := (start + end) / 2
	var taskAndPosition TaskAndPosition
	for start <= end {
		if id == tasks[mid].ID {
			taskAndPosition = TaskAndPosition{position: mid, task: tasks[mid]}
			break
		}

		if id > tasks[mid].ID {
			start = mid + 1
			mid = (start + end) / 2
		}

		if id < tasks[mid].ID {
			end = mid - 1
			mid = (start + end) / 2
		}
	}

	return taskAndPosition
}

// Commmands

func add(dataID int, description string, tasks []Task) {
	id := dataID
	id++
	createdAt := time.Now().String()
	task := Task{ID: id, Description: description, Status: "todo", CreatedAt: createdAt, UpdatedAt: createdAt}
	tasks = append(tasks, task)
	updateData(id, tasks)
	fmt.Println("The tasks variable  from add fucntion", tasks)
}

func update(dataID int, id int, tasks []Task, description string) {
	start := 0
	end := len(tasks) - 1
	mid := (start + end) / 2
	var currentTask Task
	var currentTaskPosition int
	for start <= end {
		if tasks[mid].ID == id {
			currentTask = tasks[mid]
			currentTaskPosition = mid
			break
		}
		if id > tasks[mid].ID {
			start = mid + 1
			mid = (start + end) / 2
		}

		if id < tasks[mid].ID {
			end = mid - 1
			mid = (start + end) / 2
		}
	}

	if currentTask.ID == id {
		updatedAt := time.Now().String()
		updatedTask := Task{ID: id, Description: description, Status: currentTask.Status, CreatedAt: currentTask.CreatedAt, UpdatedAt: updatedAt}
		tasks[currentTaskPosition] = updatedTask
		updateData(dataID, tasks)
	} else {
		fmt.Println("TASK(ID)", id)
		fmt.Println("Does not exist")
	}
}

func delete(dataID int, id int, tasks []Task) {
	tasksCopy := []Task{}
	taskAndPosition := findTask(id, tasks)
	if taskAndPosition == (TaskAndPosition{}) {
		fmt.Println("TASK(ID)", id)
		fmt.Println("Does not exist")
		return
	}
	for i := 1; i < len(tasks); i++ {
		if tasks[i].ID == id {
			continue
		} else {
			tasksCopy = append(tasksCopy, tasks[i])
		}
	}
	// tasks = append(tasks[:taskAndPosition.position], tasks[taskAndPosition.position:]...) // What does ... even mean? is this like javascript own ... copy operator
	fmt.Println("The tasks after delete operations", tasksCopy)
	updateData(dataID, tasksCopy)
}

func main() {
	data := loadData()
	dataID := data.ID
	tasks := data.Tasks
	arguments := os.Args[1:]
	argumentsLength := len(arguments)
	if argumentsLength == 0 {
		fmt.Println("No command was entered: Kindly input a  command")
		return
	}
	commands(arguments, argumentsLength, dataID, tasks)
	fmt.Println("Length of arguments", argumentsLength)
}
