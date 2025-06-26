package main

import (
	"fmt"
	"time"
)

func main() {

	var i int = 2

	fmt.Println("Escribe", i, "como")

	switch i {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	}

	switch time.Now().Weekday() {
	case time.Sunday:
		fmt.Println("Hoy es domingo")
	case time.Monday:
		fmt.Println("Hoy es lunes")
	case time.Tuesday:
		fmt.Println("Hoy es martes")
	case time.Wednesday:
		fmt.Println("Hoy es miercoles")
	case time.Thursday:
		fmt.Println("Hoy es jueves")
	case time.Friday:
		fmt.Println("Hoy es viernes")
	case time.Saturday:
		fmt.Println("Hoy es sabado")
	default:
		fmt.Println("No es un dia valido")
	}

}
