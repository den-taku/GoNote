package main

import (
	"fmt"
)

type ID int
type Priotiry int

type Task struct {
	ID     int
	Detail string
	done   bool
}

func ProcessTask(id ID, priority Priotiry) {
	fmt.Println(id, priority)
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

func (task *Task) Finish() {
	task.done = true
}

func main() {
	var id ID = 3
	var priority Priotiry = 5
	ProcessTask(id, priority)

	var task Task = Task{
		ID:     1,
		Detail: "buy the milk",
		done:   true,
	}
	fmt.Println(task.ID)
	fmt.Println(task.Detail)
	fmt.Println(task.done)
	fmt.Println(task)

	var task2 Task = Task{
		2, "buy the mikan", false,
	}
	fmt.Println(task2.ID)
	fmt.Println(task2.Detail)
	fmt.Println(task2.done)
	fmt.Println(task2)

	task3 := Task{}
	fmt.Println(task3.ID)
	fmt.Println(task3.Detail)
	fmt.Println(task3.done)
	fmt.Println(task3)

	var task4 Task = Task{}
	var task_ref *Task = &Task{}
	fmt.Println(task4)
	fmt.Println(task_ref)

	var task5 *Task = new(Task)
	task.ID = 1
	task.Detail = "buy the orange"
	fmt.Println(task5.Detail)
	fmt.Println(NewTask(1, "buy the milk"))

	fmt.Printf("%s", task)

	task6 := NewTask(7, "hoge")
	task6.Finish()
	fmt.Printf("%+v", task6)
}
