package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	/*	var name string

		fmt.Scanln(&name)

		fmt.Println(name)
	*/
	fmt.Println("Enter your name")

	name, _ := reader.ReadString('\n')

	fmt.Println("Entered name is ", name)

	fmt.Println("Enter any int")

	var input int64

	fmt.Scanln(&input)

	if input >= 1 && input <= 10 {
		fmt.Println("The value is between 1 and 10")
	} else {
		fmt.Println("The value of a is outside the range of 1 and 10")
	}

}
