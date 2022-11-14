package main

import "fmt"

func main() {

	//map is unordered

	sampleMap := map[string]int{}

	sampleMap["Delhi"] = 1
	sampleMap["kolkata"] = 2
	sampleMap["Mumbai"] = 3
	sampleMap["chennai"] = 4

	fmt.Println(sampleMap)

	//deletion of key
	delete(sampleMap, "kolkata")

	fmt.Println(sampleMap)

	//Adding key
	sampleMap["Chandigarh"] = 5

	fmt.Println(sampleMap)

	//iteration over map
	for key, val := range sampleMap {
		fmt.Println(key, val)
	}

}
