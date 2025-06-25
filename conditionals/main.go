package main

import "fmt"

func main() {
	
	var name string = "John" 
	var age int = 40
	
	// Conditional about the name 
	if name == "John" {
		fmt.Println("Hello John")
	} else {
		fmt.Println("Hello Stranger")
	}

	// Conditional about the age
	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You cannot vote")
	}

	// Conditional about the number
	if 8%2 == 0 {
		fmt.Println("8 is an even number")
	} else {
		fmt.Println("8 is an odd number")
	}

	// Conditional with multiple conditions
	if otherName := "Alice"; otherName == "Rodrigo" {
		fmt.Println("Hello Alice")
	} else {
		fmt.Println("Hello Unknown")
	}


} 