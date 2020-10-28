package main

import (
	"fmt"
)

func hello() {
	fmt.Println("hello")
}

func sum(i, j int) int{
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
}