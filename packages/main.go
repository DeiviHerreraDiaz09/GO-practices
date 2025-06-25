package main

import ("fmt"
		"os")

func main(){

	envVar := os.Getenv("HOME")
	if envVar == "" {
		fmt.Println("Environment variable 'HOME' is not set.")
	} else {
		fmt.Println("The value of the 'HOME' environment variable is:", envVar)
	}

}
