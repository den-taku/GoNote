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

func Print3(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("value is string: %s\n", v)
	case int:
		fmt.Printf("value is int: %d\n", v)
	case Stringer:
		fmt.Printf("value is String: %s\n", v)
	}
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

func Print2(value interface{}) {
	s, ok := value.(string)
	if ok {
		fmt.Printf("value is string: %s\n", s)
	} else {
		fmt.Printf("value is not string\n")
	}
}

// type Stringer interface {
// 	String() string
// }

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

	Print2("abc")
	Print2(10)

	Print3("abc")
	Print3(10)
	Print3(new(Task))
}
