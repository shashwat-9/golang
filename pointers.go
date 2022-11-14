package main

import "fmt"

func main() {
	var ptr *int
	var num int

	num = 15

	ptr = &num

	fmt.Println(ptr, "\n", num)

}
