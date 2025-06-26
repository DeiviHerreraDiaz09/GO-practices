package main

import "fmt"

func main(){
	
	var i int = 1
	var rango int = 10

	// For 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// For 2

	for ite := 1; ite <= 10; ite++ {
		fmt.Println(ite)
	}

	// For 3

	for limite := range rango {
		fmt.Println(limite + 1)
	}

}