package main

import "fmt"

func main() {

	//num := []int{1, 4, 56, 7}

	//append(num, {3,5,67})

	val := "Hello World"

	first := val[0:5]
	second := val[5:]

	fmt.Println(first, second)

	//Slice of a Slice
	doubleSlice := val[0:5][1:4]
	fmt.Println(doubleSlice)
}
