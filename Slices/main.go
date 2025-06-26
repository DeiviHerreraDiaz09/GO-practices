package main

import "fmt"

func main(){

	// Array default
	var cadenas[3]string = [3]string{"Hola", "Mundo", "!"} 
	convertArray(cadenas)
	
	// Slice default
	var cadenasSlice []string = []string{"Hola", "Mundo", "!"}
	convertSlice(cadenasSlice)

	fmt.Println(cadenas)      
	fmt.Println(cadenasSlice)

}

func convertArray(arr [3]string) {
	arr[2] = "Cambio"
}

func convertSlice(slice []string) {
	slice[2] = "Cambio"
}
