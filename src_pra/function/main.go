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

func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("divided by zero")
	}
	return i / j, nil
}

func sum(i, j int) int {
	return i + j
}

func swap(i, j int) (int, int) {
	return j, i
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

	n, err := div(10, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
