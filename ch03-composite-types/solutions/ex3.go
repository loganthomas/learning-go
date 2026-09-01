package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func ex3() {
	// Exercise 3
	// This exercise covers defining a struct and initializing it.
	// The struct literals with and without keys are interchangeable,
	// but using keys is more maintainable and readable.
	header("exercise 3")

	e1 := Employee{"Hingle", "McCringleberry", 1}

	e2 := Employee{
		firstName: "Jackmerius",
		lastName:  "Tacktheratrix",
		id:        2,
	}

	var e3 Employee
	e3.firstName = "Bismo"
	e3.lastName = "Funyuns"
	e3.id = 3

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}
