package main

import (
	"fmt"
)

func main() {

	fmt.Println("Insert distance in Km")
	var distance, time float64
	_, err := fmt.Scanf("%f", &distance)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Insert time minutes")
	_, err2 := fmt.Scanf("%f", &time)

	if err2 != nil {
		fmt.Println(err2)
	}

	velocity := distance / (time / 60)

	fmt.Printf("Velocity = %.2f Km/h \n ", velocity)

}
