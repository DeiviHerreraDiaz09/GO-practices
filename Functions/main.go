package main

import "fmt"

func main() {

	// Functions in GO 1
	// var number1, number2 int
	// fmt.Println("Enter the first number:")
	// fmt.Scanln(&number1)
	// fmt.Println("Enter the second number:")
	// fmt.Scanln(&number2)
	
	// result := sum(number1, number2)
	// fmt.Println(result)

	// Functions in GO 2

	var name string = "John"
	var lastName string = "Doe"

	map1, map2 := multValues(name, lastName)
	fmt.Println(map1)
	fmt.Println(map2)


}

func sum(number1, number2 int) int {
	result := number1 + number2
	return result
}

func multValues(name, lastname string) (map[string]string, map[string]string) {
	map1 := make(map[string]string)
	map2 := make(map[string]string)

	map1["name"] = name
	map1["lastname"] = lastname
	map2["name"] = name
	map2["lastname"] = lastname

	return map1, map2
}