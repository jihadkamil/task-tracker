package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Status string

const (
	Done       Status = "done"
	Inprogress Status = "in-progress"
	Todo       Status = "todo"
)

type Task struct {
	Id        int
	Name      string
	Status    Status
	CreatedAt time.Time
	UpdateAt  time.Time
}

type TaskList struct {
	Tasks []Task
}

func NewTaskList() *TaskList {
	return &TaskList{Tasks: []Task{}}
}

// get all task list
func (s *TaskList) GetAllTasks() {
	if len(s.Tasks) == 0 {
		fmt.Println("No tasks available")
		return
	}

	for _, v := range s.Tasks {
		fmt.Printf("%d. %s [%s] created at %s (last updated at %s) \n", v.Id, v.Name, v.Status, v.CreatedAt.Format("2006-01-02 15:04"), v.UpdateAt.Format("2006-01-02 15:04"))
	}
}

// get all task list filtered by status
func (s *TaskList) GetAllTasksByStatus(status Status) {
	if len(s.Tasks) == 0 {
		fmt.Println("No tasks available")
		return
	}

	for _, v := range s.Tasks {
		if v.Status == status {
			fmt.Printf("%d. %s [%s]\n", v.Id, v.Name, v.Status)
		}
	}
}

// create new task will automatically generate id, status ("todo") createdAt, updatedAt -->
func (s *TaskList) CreateTask(name string) {
	s.Tasks = append(s.Tasks, Task{
		Name:      name,
		Id:        getLastId(s.Tasks) + 1,
		Status:    Todo,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	})
	fmt.Printf("Task %s created successfully \n", name)
}

// updating task name (by id)
func (s *TaskList) UpdateTask(id, name string) {
	idInt, _ := strconv.Atoi(id)
	for index, v := range s.Tasks {
		if v.Id == idInt {
			s.Tasks[index].Name = name
			s.Tasks[index].UpdateAt = time.Now()
			fmt.Printf("Task %d updated successfully \n", idInt)
			return
		}
	}
	fmt.Printf("Task id=%d is not found \n", idInt)
}

// removing specific task (by id) from list
func (s *TaskList) DeleteTask(id string) {
	idInt, _ := strconv.Atoi(id)
	for index, v := range s.Tasks {
		if v.Id == idInt {
			s.Tasks = append(s.Tasks[:index], s.Tasks[index+1:]...)
			fmt.Printf("Task %d deleted successfully \n", idInt)
			return
		}
	}
	fmt.Printf("Task id=%d is not found \n", idInt)
}

// updated task status by id
func (s *TaskList) UpdateTaskStatus(id string, status Status) {
	idInt, _ := strconv.Atoi(id)
	for index, v := range s.Tasks {
		if v.Id == idInt {
			s.Tasks[index].Status = status
			s.Tasks[index].UpdateAt = time.Now()
			fmt.Printf("Task %d state successfully updated into %s \n", idInt, status)

			return
		}
	}
	fmt.Printf("Task id=%d is not found \n", idInt)
}

// get latest number used as id
func getLastId(listTask []Task) int {
	if len(listTask) == 0 {
		return 0 // or handle it differently if no tasks exist
	}
	return listTask[len(listTask)-1].Id
}

// validating input from user has 2 args & the last argument is suit with the task name rule
func isValidAddInput(inputs []string) bool {
	if len(inputs) == 0 {
		return false
	}
	// Check if input starts with 'add "'
	if strings.HasPrefix(inputs[1], `"`) && strings.HasSuffix(inputs[1], `"`) {
		// Remove the "add " part and check if the value inside quotes is non-empty
		content := strings.TrimSpace(inputs[1][0:])
		if len(content) >= 2 && strings.HasPrefix(content, `"`) && strings.HasSuffix(content, `"`) {
			return true
		}
	}
	return false
}

// validating input from user has 3 args & the last argument is suit with the task name rule
func isValidUpdateInput(inputs []string) bool {
	if len(inputs) > 2 {
		// check id
		if _, err := strconv.Atoi(inputs[1]); err != nil {
			return false
		}

		// check task name
		if valid := isValidAddInput(inputs[1:]); !valid {
			return false
		}
		return true
	}
	return false

}

// validating input from user has 2 args
// this func can be used by other command with same input validation
func isValidDeleteInput(inputs []string) bool {
	if len(inputs) > 1 {
		// check id
		if _, err := strconv.Atoi(inputs[1]); err == nil {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("--------------------------- START ---------------------------")

	reader := bufio.NewReader(os.Stdin)
	tasks := NewTaskList()

	for {
		fmt.Print("task cli> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")
		if len(args) == 0 {
			continue
		}

		// read the command
		command := strings.ToLower(args[0])

		switch {
		case command == "add" && isValidAddInput(args):
			tasks.CreateTask(args[1])

		case command == "update" && isValidUpdateInput(args):
			tasks.UpdateTask(args[1], args[2])

		case command == "delete" && isValidDeleteInput(args):
			tasks.DeleteTask(args[1])
		case command == "mark-in-progress" && isValidDeleteInput(args):
			tasks.UpdateTaskStatus(args[1], Inprogress)

		case command == "mark-done" && isValidDeleteInput(args):
			tasks.UpdateTaskStatus(args[1], Done)

		case command == "list":
			if len(args) > 1 {
				tasks.GetAllTasksByStatus(Status(args[1]))
				continue
			}
			tasks.GetAllTasks()

		case command == "exit":
			fmt.Println("Exiting program")
			return
		case command == "clear":
			fmt.Print("\033[H\033[2J")
		default:
			fmt.Println("Unknown command", command)
		}

	}
}
