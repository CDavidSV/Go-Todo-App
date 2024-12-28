package framework

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/CDavidSV/go-todo-app/config"
)

type Task struct {
	ID           int
	Title        string
	Description  string
	CreationDate time.Time
	Category     string
	Completed    bool
}

type TodoList struct {
	tasks map[int]*Task
}

var defaultFilePath = "taskdata.csv"

func parseCSVTaskRecord(record []string, destination *Task) error {
	if len(record) != 6 {
		return fmt.Errorf("record length does not match struct field count")
	}

	ID, err := strconv.Atoi(record[0])
	if err != nil {
		return fmt.Errorf("error parsing ID: %v", err)
	}

	creationDate, err := time.Parse(time.RFC822, record[3])
	if err != nil {
		return fmt.Errorf("error parsing creation date: %v", err)
	}

	completed, err := strconv.ParseBool(record[5])
	if err != nil {
		return fmt.Errorf("error parsing completed field: %v", err)
	}

	destination.ID = ID
	destination.Title = record[1]
	destination.Description = record[2]
	destination.CreationDate = creationDate
	destination.Category = record[4]
	destination.Completed = completed

	return nil
}

func FormatTaskForTable(task Task) []string {
	var completedSymbol string
	if task.Completed {
		completedSymbol = "✔"
	} else {
		completedSymbol = "✘"
	}

	return []string{strconv.Itoa(task.ID), task.Title, task.Description, task.CreationDate.Format("02 Jan 06 15:04"), task.Category, completedSymbol}
}

func FormatTaskForCSV(task Task) []string {
	return []string{strconv.Itoa(task.ID), task.Title, task.Description, task.CreationDate.Format(time.RFC822), task.Category, strconv.FormatBool(task.Completed)}
}

func sortTasksByDate(tasks [][]string) {
	sort.Slice(tasks, func(i, j int) bool {
		timeI, _ := time.Parse("02 Jan 06 15:04", tasks[i][3])
		timeJ, _ := time.Parse("02 Jan 06 15:04", tasks[j][3])

		taskI, _ := strconv.Atoi(tasks[i][0])
		taskJ, _ := strconv.Atoi(tasks[j][0])

		return timeI.Before(timeJ) || taskI < taskJ
	})
}

func NewTodoList() *TodoList {
	todoList := &TodoList{
		tasks: make(map[int]*Task),
	}

	file, err := os.Open(defaultFilePath)
	if err != nil {
		// File does not exists or cannot be read
		return todoList
	}
	defer file.Close()

	// Load tasks from file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	// If the length of the records is less than 1, then the file is empty
	if len(records) < 1 {
		return todoList
	}

	for _, record := range records {
		var task Task
		if err := parseCSVTaskRecord(record, &task); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing task record, file may be corrupted: %v\n", err)
			continue
		}

		todoList.tasks[task.ID] = &task
	}

	return todoList
}

func (t *TodoList) AddTask(title, description, categoryName string) {
	defer t.save()

	task := &Task{
		ID:           len(t.tasks) + 1,
		Title:        title,
		Description:  description,
		CreationDate: time.Now(),
		Category:     categoryName,
		Completed:    false,
	}

	t.tasks[task.ID] = task
}

func (t *TodoList) AddGenericTask(title, description string) {
	t.AddTask(title, description, "general")
}

func (t *TodoList) RemoveTask(ID int) (Task, bool) {
	defer t.save()

	defer delete(t.tasks, ID)

	if task, ok := t.tasks[ID]; ok {
		return *task, true
	}
	return Task{}, false
}

func (t *TodoList) CompleteTask(ID int) bool {
	defer t.save()

	if task, ok := t.tasks[ID]; ok {
		task.Completed = true
		return true
	}

	// Task does not exist
	return false
}

func (t *TodoList) ListAllTasks(completed bool) [][]string {
	return t.ListTasks("", completed)
}

func (t *TodoList) ListTasks(category string, showCompleted bool) [][]string {
	tasks := [][]string{}

	noCategory := category == ""
	for _, task := range t.tasks {
		// If no category is specified, we list all tasks regardless of category
		if !noCategory && category != task.Category {
			continue
		}

		if !showCompleted && task.Completed {
			continue
		}
		tasks = append(tasks, FormatTaskForTable(*task))

	}

	sortTasksByDate(tasks)
	return tasks
}

func (t *TodoList) convertTasksToCSVFormat() [][]string {
	tasks := make([][]string, len(t.tasks))

	i := 0
	for _, task := range t.tasks {
		formatedTask := FormatTaskForCSV(*task)

		tasks[i] = formatedTask
		i++
	}

	return tasks
}

func (t *TodoList) save() {
	tasks := t.convertTasksToCSVFormat()

	tempFilePath := defaultFilePath + ".tmp"
	file, err := os.Create(tempFilePath)

	writer := csv.NewWriter(file)
	err = writer.WriteAll(tasks)
	if err != nil {
		os.Remove(tempFilePath)
		log.Fatalf(config.ErrorSytle.Render("error saving to file: %v"), err)
	}

	file.Close()

	err = os.Rename(tempFilePath, defaultFilePath)
	if err != nil {
		os.Remove(tempFilePath)
		log.Fatalf(config.ErrorSytle.Render("error saving to file: %v"), err)
	}
}
