package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {
	firstName, lastName string
	age                 uint8
	role                string
}

var singleInstance *Single

func getInstance() *Single {

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("Creating Single instance")
			singleInstance = &Single{firstName: "Shashwat", lastName: "Mishra", age: 22, role: "Intern"}
		} else {
			fmt.Println("Single instance already exists")
		}
	} else {
		fmt.Println("Single instance already exists")
	}

	return singleInstance
}

func main() {
	for i := 0; i < 100; i++ {
		var singleObj = getInstance()

		fmt.Println(*singleObj)
		fmt.Println("Iteration No.", i)
	}
}
