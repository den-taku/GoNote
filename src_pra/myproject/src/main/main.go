package main
import (
	f"fmt"
	."strings"
)

// var message string = "hello world"

func hello() {
	f.Println("hello")
}

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

	n = 10
	switch n {
	case 15:
		f.Println("FizzBuzz")
	case 5, 10:
		f.Println("Buzz")
	case 3, 6, 9:
		f.Println("Fizz")
	default:
		f.Println(n)
	}
	n = 3
	switch n {
	case 3:
		n = n - 1
		fallthrough
	case 2:
		n = n - 1
		fallthrough
	case 1:
		n = n - 1
		f.Println(n)
		
	}
	n = 10
	switch {
	case n%15==0:
		f.Println("FizzBuzz")
	case n%5==0:
		f.Println("Buzz")
	case n%3==0:
		f.Println("Fizz")
	default:
		f.Println(n)
		
	}
	hello()
}