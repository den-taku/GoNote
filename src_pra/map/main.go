package main

import(
	"fmt"
)

var month map[int]string = map[int]string {}

func main() {
	month[1] = "January"
	month[2] = "February"
	fmt.Println(month)

	month2 := map[int]string {
		1: "Junuary",
		2: "February",
	}
	fmt.Println(month2)

	jan := month[1]
	fmt.Println(jan)

	_, ok := month[1]
	if ok {
		fmt.Println("yes")
	}
	delete(month, 1)
	fmt.Println(month)

	for key, value := range month2 {
		fmt.Printf("%d %s\n", key, value)
	}
}