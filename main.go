package main

import (
	"fmt"

	correctimplementation "github.com/odanaraujo/golang/solid/single-responsability-principle/correct_implementation"
)

func main() {
	// people := godobject.People{
	// 	Name: "John",
	// 	Age:  20,
	// }

	// fmt.Println(people.PrintPeople())

	people := correctimplementation.People{
		Name: "John",
		Age:  20,
	}

	fmt.Println(people.PrintPeople())
}
