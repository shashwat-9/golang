package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	PhoneNo string `json:"phone_no"`
	Role    string `json:"role"`
}

func main() {

	emp := Employee{"Shashwat", 22, "9631674355", "Intern"}

	fmt.Println(emp)

	jsonVal, err := json.Marshal(emp)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonVal))

	jsonIndentVal, err1 := json.MarshalIndent(emp, "", "  ")

	if err1 != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonIndentVal))
}
