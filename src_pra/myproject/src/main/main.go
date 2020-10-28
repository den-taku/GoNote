package main
import (
	f"fmt"
	."strings"
)

// var message string = "hello world"

var foo, bar, buz string = "foo", "bar", "buz"

var (
	a string = "aaa"
	b = "bbb"
	c = "ccc"
	d = "ddd"
	e = "eee"
)

func main() {
	message := "hello world"
	const Hello string = "hello"
	// Hello = "bye" error
	var i int
	f.Println(i)
	f.Println(ToUpper(message))

	a, b := 10, 100
	if a > b {
		f.Println("a is larger than b")
	} else if a < b {
		f.Println("a is smaller than b")
	} else {
		f.Println("a equals b")
	} 

	for i := 0; i < 10; i++ {
		f.Println(i)
	}
	n := 0
	for n < 10 {
		f.Printf("n = %d\n", n)
		n++
	}
	m := 0
	for {
		m++
		if m > 10 {
			break
		}
		if m % 2 == 0 {
			continue
		}
		f.Println(m)
	}
}