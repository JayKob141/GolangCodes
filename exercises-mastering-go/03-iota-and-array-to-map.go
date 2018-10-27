package main

import (
	"fmt"
	"strconv"
)

type Power4 int

func main() {
	fmt.Println("---------")

	// some array
	days := [7]string{"monday", "tuesday", "wednesday", "thrusday", "friday", "saturday", "sunday"}

	const (
		lunes = iota
		martes
		miercoles
		jueves
		viernes
	)

	fmt.Println(days[lunes])
	fmt.Println(days[martes])
	fmt.Println(days[miercoles])

	const (
		p4_0 Power4 = 1 << (2 * iota)
		p4_1
		p4_2
		p4_3
		p4_4
		p4_5
	)

	fmt.Println("4^0:", p4_0)
	fmt.Println("4^1:", p4_1)
	fmt.Println("4^2:", p4_2)
	fmt.Println("4^3:", p4_3)
	fmt.Println("4^4:", p4_4)

	// exercise: convert any array into map
	fmt.Println(days)
	mapped := map[string]string{}
	for i, value := range days {
		newKey := strconv.Itoa(i)
		mapped[newKey] = value
	}
	fmt.Println(mapped)

}
