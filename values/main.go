package main

import "fmt"

func main(){

	// Using different types of variables
	var Integer int = 30
	var Float float64 = 3.14
	result := Integer + int(Float)
	fmt.Println("The result is:", result)
}