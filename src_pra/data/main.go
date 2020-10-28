package main

import (
	"fmt"
)

var arr [4]string

func fn(arr [4]string) {
	arr[0] = "x"
	fmt.Println(arr)
}

func main() {

	arr[0] = "a"
	arr[1] = "b"
	arr[2] = "c"
	arr[3] = "d"

	fmt.Println(arr[0])
	fmt.Println(arr)

	arr1 := [...]string{"a", "b", "c", "d"}
	fmt.Println(arr1)

	// var arr2 [5]string

	fn(arr1)
	// fn(arr2)

	fmt.Println(arr1)

	s := []string{"a", "b", "c", "d"}
	fmt.Println(s)

	s = append(s, "a")
	s = append(s, "b")
	s = append(s, "c", "d")
	fmt.Println(s)

	for i, s := range s {
		fmt.Println(i, s)
	}
}
