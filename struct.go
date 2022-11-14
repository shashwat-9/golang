package main

import "fmt"

func main() {
	type house struct {
		noRooms int32
		price   float32
		city    string
	}

	h1 := house{5, 10000.0, "kolkata"}

	fmt.Println(h1)

	h1.price = 12000.0
	h1.noRooms = 3

	fmt.Println(h1)
}
