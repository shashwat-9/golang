package main

import "fmt"

func f1() {
	fmt.Println("This function is being called by another function")
}

func add(x int, y int) int {

	f1()
	return x + y
}

func main() {

	fmt.Println("Enter two no.s")

	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println("Sum of the input no.s ", add(x, y))

}
