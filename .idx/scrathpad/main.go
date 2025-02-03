package main

import (
	"fmt"
)

func sum(i, j int) int {
	return i + j
}

type myenum string

const (
	AK myenum = "Aadya"
	RK myenum = "Rhea"
	SK myenum = "Sanjay"
	GG myenum = "Gunjan"
)

var myenumValues = []myenum{AK, RK, SK, GG}

func main() {
	fmt.Println("go ")
	fmt.Println(sum(1, 2))

	for _, val := range myenumValues {
		fmt.Println(val)
	}

	fmt.Println("go end")
}
