package main

import "fmt"

func main() {
	var date, month, year, age int

	fmt.Println("Enter the day month year and age in the given order")
	fmt.Scan(&date, &month, &year, &age)

	fmt.Println("Current year", (year + age))

	var a = []int32{1, 2, 4, 6, 8}

	avg := (float32)(a[0]+a[1]+a[2]+a[3]+a[4]) / 5.0
	fmt.Println("Average of the hardcoded vals ", avg)

}
