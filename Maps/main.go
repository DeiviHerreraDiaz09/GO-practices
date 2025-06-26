package main

import "fmt"

func main(){

	var mapp map[string]int = make(map[string]int, 5)
	mapp["Uno"] = 1
	fmt.Println(mapp)

	// Methods of the map
	delete(mapp, "Uno")
	clear(mapp)
	
}