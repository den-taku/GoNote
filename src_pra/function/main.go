package main

import (
	"fmt"
	// "os"
	"errors"
	"log"
)

func hello() {
	fmt.Println("hello")
}

func div(i, j int) (result int, err error) {
	if j == 0 {
		err = errors.New("divided by zero")
		return
	}
	result = i / j
	return
}

func sum(i, j int) int {
	return i + j
}

func swap(i, j int) (int, int) {
	return j, i
}

var sum2 func(i, j int) = func(i, j int) {
	fmt.Println(i + j)
}

func main() {
	hello()
	n := sum(1, 2)
	fmt.Println(n)
	x, y := 3, 4
	x, y = swap(x, y)
	fmt.Println(x, y)
	x, _ = swap(x, y)
	fmt.Println(x)
	swap(x, y)

	// file, err := os.Open("hoge.go")
	// if err != nil {
	// Err handling
	// }

	n, err := div(10, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	func(i, j int) {
		fmt.Println(i + j)
	}(2, 4)

	sum2(2, 4)

}
