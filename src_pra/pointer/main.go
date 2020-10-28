package main

import(
	"fmt"
	"os"
	"log"
	"errors"
)

func callByValue(i int) {
	i = 20
}

func callByRef(i *int) {
	*i = 20
}

func main() {
	var i int = 10
	callByValue(i)
	fmt.Println(i)
	callByRef(&i)
	fmt.Println(i)

	file, err := os.Open("./error.go")
	if err != nil {
		//
	}
	//
	defer file.Close()

	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()

	a := []int{1, 2, 3}
	// fmt.Println(a[10])
	fmt.Println(a)

	a2 := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		if i >= len(a2) {
			panic(errors.New("index out of range"))
		}
		fmt.Println(a2[i])
	}
}