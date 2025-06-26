package main

import "fmt"

func main() {

	var contenedor[5]int
	var i int = 0

	
	for i < len(contenedor) {
		contenedor[i] = i + 1
		i++
	}
	
	fmt.Println("Arreglo completo:", contenedor)
}