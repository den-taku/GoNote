package main

import (
	"fmt"
)

type Task struct {
	ID     int
	Detail string
	done   bool
	*User
}

func (task Task) String() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

type Stringer interface {
	String() string
}

func Print(stringer Stringer) {
	fmt.Println(stringer.String())
}

type User struct {
	FirstName string
	LastName  string
}

func (u *User) FullName() string {
	fullname := fmt.Sprintf("%s %s",
		u.FirstName, u.LastName)
	return fullname
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}

func NewTask(id int, detail,
	firstName, lastName string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
		User:   NewUser(firstName, lastName),
	}
	return task
}

func main() {
	task := NewTask(1, "buy the milk", "Jxck", "Daniel")
	fmt.Println(task.FirstName)
	fmt.Println(task.LastName)
	fmt.Println(task.FullName())
	fmt.Println(task.User)
	fmt.Println(task)

	var i uint8 = 3
	var j uint32 = uint32(i)
	fmt.Println(j)

	var s string = "abc"
	var b []byte = []byte(s)
	fmt.Println(b)
}
